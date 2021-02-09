package gormopentracing

import "gorm.io/gorm"

type operationName string

func (op operationName) String() string {
	return string(op)
}

const (
	_createOp operationName = "create"
)

func (p opentracingPlugin) beforeCreate(db *gorm.DB) {
	injectBefore(db, _createOp)
}

func (p opentracingPlugin) afterCreate(db *gorm.DB) {
	extractAfter(db)
}
