package gormopentracing

import "gorm.io/gorm"

type opentracingPlugin struct {
	// logResult means log SQL operation result into span log which causes span size grows up.
	// This is advised to only open in developing environment.
	logResult bool
}

// New constructs a new plugin based opentracing. It supports to trace all operations in gorm,
// so if you have already traced your servers, now this plugin will perfect your tracing job.
func New(opts ...applyOption) gorm.Plugin {
	dst := new(options)
	for _, apply := range opts {
		apply(dst)
	}

	return opentracingPlugin{
		logResult: dst.logResult,
	}
}

func (p opentracingPlugin) Name() string {
	return "opentracing"
}

// Initialize registers all needed callbacks
func (p opentracingPlugin) Initialize(db *gorm.DB) error {
	// create
	_ = db.Callback().Create().Before("gorm:create").
		Register("opentracing:before_create", p.beforeCreate)
	_ = db.Callback().Create().After("gorm:create").
		Register("opentracing:after_create", p.after)

	// update
	_ = db.Callback().Update().Before("gorm:update").
		Register("opentracing:before_update", p.beforeUpdate)
	_ = db.Callback().Update().After("gorm:update").
		Register("opentracing:after_update", p.after)

	// query
	_ = db.Callback().Query().Before("gorm:query").
		Register("opentracing:before_query", p.beforeQuery)
	_ = db.Callback().Query().After("gorm:query").
		Register("opentracing:after_query", p.after)

	// delete
	_ = db.Callback().Delete().Before("gorm:delete").
		Register("opentracing:before_delete", p.beforeDelete)
	_ = db.Callback().Delete().After("gorm:delete").
		Register("opentracing:after_delete", p.after)

	// row
	_ = db.Callback().Row().Before("gorm:row").
		Register("opentracing:before_row", p.beforeRow)
	_ = db.Callback().Row().After("gorm:row").
		Register("opentracing:after_row", p.after)

	// raw
	_ = db.Callback().Raw().Before("gorm:raw").
		Register("opentracing:before_raw", p.beforeRaw)
	_ = db.Callback().Raw().After("gorm:raw").
		Register("opentracing:after_raw", p.after)

	return nil
}
