package gormopentracing

import "gorm.io/gorm"

func (p opentracingPlugin) beforeCreate(db *gorm.DB) {
	p.injectBefore(db, p.opt.createOpName)
}

func (p opentracingPlugin) after(db *gorm.DB) {
	p.extractAfter(db)
}

func (p opentracingPlugin) beforeUpdate(db *gorm.DB) {
	p.injectBefore(db, p.opt.updateOpName)
}

func (p opentracingPlugin) beforeQuery(db *gorm.DB) {
	p.injectBefore(db, p.opt.queryOpName)
}

func (p opentracingPlugin) beforeDelete(db *gorm.DB) {
	p.injectBefore(db, p.opt.deleteOpName)
}

func (p opentracingPlugin) beforeRow(db *gorm.DB) {
	p.injectBefore(db, p.opt.rowOpName)
}

func (p opentracingPlugin) beforeRaw(db *gorm.DB) {
	p.injectBefore(db, p.opt.rawOpName)
}
