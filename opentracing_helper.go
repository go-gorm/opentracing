package gormopentracing

import (
	"context"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	_prefix      = "gorm.opentracing"
	_errorTagKey = "error"
)

var (
	// span.Tag keys
	_tableTagKey = keyWithPrefix("table")
	// span.Log keys
	// _errorLogKey        = keyWithPrefix("error")
	_resultLogKey       = keyWithPrefix("result")
	_sqlLogKey          = keyWithPrefix("sql")
	_rowsAffectedLogKey = keyWithPrefix("rowsAffected")
)

func keyWithPrefix(key string) string {
	return _prefix + "." + key
}
func getSpanCtx(c context.Context, spanCtxKey string) (opentracing.SpanContext, bool) {
	var spanCtx opentracing.SpanContext
	spanCtxData := c.Value(spanCtxKey)
	spanCtx, ok := spanCtxData.(opentracing.SpanContext)
	return spanCtx, ok
}
func getOpNameCtx(c context.Context, opNameKey string) string {
	opNameData := c.Value(opNameKey)
	opName, ok := opNameData.(string)
	if ok {
		return opName
	}
	return ""
}

var (
	opentracingSpanKey = "opentracing:span"
	json               = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (p opentracingPlugin) injectBefore(db *gorm.DB, op operationName) {
	// make sure context could be used
	if db == nil {
		return
	}

	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "could not inject sp from nil Statement.Context or nil Statement")
		return
	}
	ctx := db.Statement.Context
	spanCtx, ok := getSpanCtx(ctx, p.opt.spanCtxKey)
	if !ok {
		return
	}
	opName := getOpNameCtx(ctx, p.opt.opNameKey)
	if opName == "" {
		opName = op.String()
	}
	sp := p.opt.tracer.StartSpan(opName, opentracing.ChildOf(spanCtx))
	db.InstanceSet(opentracingSpanKey, sp)
}

func (p opentracingPlugin) extractAfter(db *gorm.DB) {
	// make sure context could be used
	if db == nil {
		return
	}
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "could not extract sp from nil Statement.Context or nil Statement")
		return
	}

	// extract sp from db context
	// sp := opentracing.SpanFromContext(db.Statement.Context)
	v, ok := db.InstanceGet(opentracingSpanKey)
	if !ok || v == nil {
		return
	}

	sp, ok := v.(opentracing.Span)
	if !ok || sp == nil {
		return
	}
	defer sp.Finish()

	// tag and log fields we want.
	tag(sp, db)
	log(sp, db, p.opt.logResult, p.opt.logSqlParameters)
}

// tag called after operation
func tag(sp opentracing.Span, db *gorm.DB) {
	if err := db.Error; err != nil {
		sp.SetTag(_errorTagKey, true)
	}

	sp.SetTag(_tableTagKey, db.Statement.Table)
}

// log called after operation
func log(sp opentracing.Span, db *gorm.DB, verbose bool, logSqlVariables bool) {
	fields := make([]opentracinglog.Field, 0, 4)
	fields = appendSql(fields, db, logSqlVariables)
	fields = append(fields, opentracinglog.Object(_rowsAffectedLogKey, db.Statement.RowsAffected))

	// log error
	if err := db.Error; err != nil {
		fields = append(fields, opentracinglog.Error(err))
	}

	if verbose && db.Statement.Dest != nil {
		// DONE(@yeqown) fill result fields into span log
		// FIXED(@yeqown) db.Statement.Dest still be metatable now ?
		v, err := json.Marshal(db.Statement.Dest)
		if err == nil {
			fields = append(fields, opentracinglog.String(_resultLogKey, *(*string)(unsafe.Pointer(&v))))
		} else {
			db.Logger.Error(context.Background(), "could not marshal db.Statement.Dest: %v", err)
		}
	}

	sp.LogFields(fields...)
}

func appendSql(fields []opentracinglog.Field, db *gorm.DB, logSqlVariables bool) []opentracinglog.Field {
	if logSqlVariables {
		fields = append(fields, opentracinglog.String(_sqlLogKey,
			db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	} else {
		fields = append(fields, opentracinglog.String(_sqlLogKey, db.Statement.SQL.String()))
	}
	return fields
}
