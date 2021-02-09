package gormopentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	_prefix      = "gorm.opentracing"
	_errorTagKey = "error"
)

var (
	_tableTagKey = keyWithPrefix("table")

	_errorLogKey  = keyWithPrefix("error")
	_resultLogKey = keyWithPrefix("result")
	_sqlLogKey    = keyWithPrefix("sql")
	_varsLogKey   = keyWithPrefix("vars")
)

func keyWithPrefix(key string) string {
	return _prefix + "." + key
}

var (
	opentracingSpanKey = "opentracing:span"
)

func injectBefore(db *gorm.DB, op operationName) {
	// make sure context could be used
	if db == nil {
		return
	}

	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "could not inject sp from nil Statement.Context or nil Statement")
		return
	}

	sp, _ := opentracing.StartSpanFromContext(db.Statement.Context, op.String())
	db.InstanceSet(opentracingSpanKey, sp)
}

func extractAfter(db *gorm.DB, verbose bool) {
	// make sure context could be used
	if db == nil {
		return
	}
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "could not extract sp from nil Statement.Context or nil Statement")
		return
	}

	// extract sp from db context
	//sp := opentracing.SpanFromContext(db.Statement.Context)
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
	log(sp, db, verbose)
}

// tag called after operation
func tag(sp opentracing.Span, db *gorm.DB) {
	if err := db.Error; err != nil {
		sp.SetTag(_errorTagKey, true)
	}

	sp.SetTag(_tableTagKey, db.Statement.Table)
}

// log called after operation
func log(sp opentracing.Span, db *gorm.DB, verbose bool) {
	fields := make([]opentracinglog.Field, 0, 4)
	fields = append(fields, opentracinglog.String(_sqlLogKey, db.Statement.SQL.String()))
	fields = append(fields, opentracinglog.Object(_varsLogKey, db.Statement.Vars))

	if err := db.Error; err != nil {
		fields = append(fields, opentracinglog.String(_errorLogKey, err.Error()))
	}

	if verbose {
		// DONE(@yeqown) fill result fields into span log
		// FIXME(@yeqown) db.Statement.Dest still be metatable now ?
		fields = append(fields, opentracinglog.Object(_resultLogKey, db.Statement.Dest))
	}

	sp.LogFields(fields...)
}
