package gormopentracing

import (
	"testing"

	"github.com/opentracing/opentracing-go"

	"github.com/stretchr/testify/assert"
)

func Test_Option_WithLogResult(t *testing.T) {
	opt := WithLogResult(true)
	p := New(opt)
	assert.Equal(t, true, p.(opentracingPlugin).opt.logResult)
}

func Test_Option_DefaultOption(t *testing.T) {
	opt := defaultOption()
	assert.Equal(t, false, opt.logResult)
	assert.Equal(t, opentracing.GlobalTracer(), opt.tracer)
	assert.Equal(t, true, opt.logSqlParameters)
}

func Test_Option_WithTracer_nil(t *testing.T) {
	opt := defaultOption()
	WithTracer(nil)(opt)
	assert.Equal(t, opentracing.GlobalTracer(), opt.tracer)
}

func Test_Option_WithSqlParameters(t *testing.T) {
	opt := WithSqlParameters(false)
	p := New(opt)
	assert.Equal(t, false, p.(opentracingPlugin).opt.logSqlParameters)
}
