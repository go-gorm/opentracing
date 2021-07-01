package gormopentracing

import "github.com/opentracing/opentracing-go"

const (
	defaultSpanCtxKey = "gormTraceSpanCtx"
	defaultOpNameKey  = "gormTraceOpName"
)

type options struct {
	// logResult means log SQL operation result into span log which causes span size grows up.
	// This is advised to only open in developing environment.
	logResult bool

	// tracer allows users to use customized and different tracer to makes tracing clearly.
	tracer opentracing.Tracer

	// Whether to log statement parameters or leave placeholders in the queries.
	logSqlParameters bool
	spanCtxKey       string
	opNameKey        string
}

func defaultOption() *options {
	return &options{
		logResult:        false,
		tracer:           opentracing.GlobalTracer(),
		logSqlParameters: true,
		spanCtxKey:       defaultSpanCtxKey,
		opNameKey:        defaultOpNameKey,
	}
}

type applyOption func(o *options)

// WithLogResult enable opentracingPlugin to log the result of each executed sql.
func WithLogResult(logResult bool) applyOption {
	return func(o *options) {
		o.logResult = logResult
	}
}

// WithTracer allows to use customized tracer rather than the global one only.
func WithTracer(tracer opentracing.Tracer) applyOption {
	return func(o *options) {
		if tracer == nil {
			return
		}

		o.tracer = tracer
	}
}
func WithSqlParameters(logSqlParameters bool) applyOption {
	return func(o *options) {
		o.logSqlParameters = logSqlParameters
	}
}
func WithSpanCtxKey(spanCtxKey string) applyOption {
	return func(o *options) {
		o.spanCtxKey = spanCtxKey
	}
}
func WithOpNameKey(opNameKey string) applyOption {
	return func(o *options) {
		o.opNameKey = opNameKey
	}
}
