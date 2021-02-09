package gormopentracing

// operationName defines a type to wrap the name of each operation name.
type operationName string

// String returns the actual string of operationName.
func (op operationName) String() string {
	return string(op)
}

const (
	_createOp operationName = "create"
	_updateOp operationName = "update"
	_queryOp  operationName = "query"
	_deleteOp operationName = "delete"
	_rowOp    operationName = "row"
	_rawOp    operationName = "raw"
)
