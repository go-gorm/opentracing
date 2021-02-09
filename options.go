package gormopentracing

type options struct {
	logResult bool
}

type applyOption func(o *options)

// WithLogResult enable opentracingPlugin to log the result of each executed sql.
func WithLogResult(logResult bool) applyOption {
	return func(o *options) {
		o.logResult = logResult
	}
}
