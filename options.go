package gormopentracing

import "github.com/opentracing/opentracing-go"

type options struct {
	// logResult means log SQL operation result into span log which causes span size grows up.
	// This is advised to only open in developing environment.
	logResult bool

	// tracer allows users to use customized and different tracer to makes tracing clearly.
	tracer opentracing.Tracer

	// Whether to log statement parameters or leave placeholders in the queries.
	logSqlParameters bool

	// errorTagHook allows users to customized error what kind of error tag should be tagged.
	errorTagHook errorTagHook

	// createOpName defines operation name for "create" span
	createOpName operationName

	// updateOpName defines operation name for "update" span
	updateOpName operationName

	// queryOpName defines operation name for "query" span
	queryOpName operationName

	// deleteOpName defines operation name for "delete" span
	deleteOpName operationName

	// rowOpName defines operation name for "row" span
	rowOpName operationName

	// rawOpName defines operation name for "raw" span
	rawOpName operationName
}

func defaultOption() *options {
	return &options{
		logResult:        false,
		tracer:           opentracing.GlobalTracer(),
		logSqlParameters: true,
		errorTagHook:     defaultErrorTagHook,
		createOpName:     _createOp,
		updateOpName:     _updateOp,
		queryOpName:      _queryOp,
		deleteOpName:     _deleteOp,
		rowOpName:        _rowOp,
		rawOpName:        _rawOp,
	}
}

type ApplyOption func(o *options)

// WithLogResult enable opentracingPlugin to log the result of each executed sql.
func WithLogResult(logResult bool) ApplyOption {
	return func(o *options) {
		o.logResult = logResult
	}
}

// WithTracer allows to use customized tracer rather than the global one only.
func WithTracer(tracer opentracing.Tracer) ApplyOption {
	return func(o *options) {
		if tracer == nil {
			return
		}

		o.tracer = tracer
	}
}

func WithSqlParameters(logSqlParameters bool) ApplyOption {
	return func(o *options) {
		o.logSqlParameters = logSqlParameters
	}
}

func WithErrorTagHook(errorTagHook errorTagHook) ApplyOption {
	return func(o *options) {
		if errorTagHook == nil {
			return
		}

		o.errorTagHook = errorTagHook
	}
}

func WithCreateOpName(name operationName) ApplyOption {
	return func(o *options) {
		if name == "" {
			return
		}

		o.createOpName = name
	}
}

func WithUpdateOpName(name operationName) ApplyOption {
	return func(o *options) {
		if name == "" {
			return
		}

		o.updateOpName = name
	}
}

func WithQueryOpName(name operationName) ApplyOption {
	return func(o *options) {
		if name == "" {
			return
		}

		o.queryOpName = name
	}
}

func WithDeleteOpName(name operationName) ApplyOption {
	return func(o *options) {
		if name == "" {
			return
		}

		o.deleteOpName = name
	}
}

func WithRowOpName(name operationName) ApplyOption {
	return func(o *options) {
		if name == "" {
			return
		}

		o.rowOpName = name
	}
}

func WithRawOpName(name operationName) ApplyOption {
	return func(o *options) {
		if name == "" {
			return
		}

		o.rawOpName = name
	}
}
