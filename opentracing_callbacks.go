package gormopentracing

import "gorm.io/gorm"

func (p opentracingPlugin) beforeCreate(db *gorm.DB) {
	p.injectBefore(db, _createOp)
}

func (p opentracingPlugin) after(db *gorm.DB) {
	p.extractAfter(db)
}

func (p opentracingPlugin) beforeUpdate(db *gorm.DB) {
	p.injectBefore(db, _updateOp)
}

func (p opentracingPlugin) beforeQuery(db *gorm.DB) {
	p.injectBefore(db, _queryOp)
}

func (p opentracingPlugin) beforeDelete(db *gorm.DB) {
	p.injectBefore(db, _deleteOp)
}

func (p opentracingPlugin) beforeRow(db *gorm.DB) {
	p.injectBefore(db, _rowOp)
}

func (p opentracingPlugin) beforeRaw(db *gorm.DB) {
	p.injectBefore(db, _rawOp)
}
