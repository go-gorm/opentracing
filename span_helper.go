package gormopentracing

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

var (
	opentracingSpanKey = "opentracing:span"
)

func injectBefore(db *gorm.DB, op operationName) {
	// make sure context could be used
	if db == nil || db.Statement == nil || db.Statement.Context == nil {
		return
	}

	sp, _ := opentracing.StartSpanFromContext(db.Statement.Context, op.String())
	db.InstanceSet(opentracingSpanKey, sp)
}

func extractAfter(db *gorm.DB) {
	// make sure context could be used
	if db == nil || db.Statement == nil || db.Statement.Context == nil {
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
	log(sp, db)
}

// tag called after operation
func tag(sp opentracing.Span, db *gorm.DB) {

}

// log called after operation
func log(sp opentracing.Span, db *gorm.DB) {

}
