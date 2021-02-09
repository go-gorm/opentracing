package gormopentracing

import "gorm.io/gorm"

type opentracingPlugin struct {
}

// New constructs a new plugin based opentracing. It supports to trace all operations in gorm,
// so if you have already traced your servers, now this plugin will perfect your tracing job.
func New() gorm.Plugin {
	return opentracingPlugin{}
}

func (p opentracingPlugin) Name() string {
	return "opentracing"
}

func (p opentracingPlugin) Initialize(db *gorm.DB) error {
	// register needed callbacks
	_ = db.Callback().Create().Before("gorm:create").
		Register("opentracing:before_create", p.beforeCreate)
	_ = db.Callback().Create().After("gorm:create").
		Register("opentracing:after_create", p.afterCreate)

	return nil
}
