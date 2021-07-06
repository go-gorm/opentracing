package gormopentracing

import (
	"reflect"
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

func Test_WithErrorTagHook(t *testing.T) {
	opt := defaultOption()
	assert.NotNil(t, opt.errorTagHook)

	// apply
	var hook errorTagHook = func(sp opentracing.Span, err error) {
		println("i do nothing with error")
	}
	WithErrorTagHook(hook)(opt)

	v1 := reflect.ValueOf(hook)
	v2 := reflect.ValueOf(opt.errorTagHook)
	assert.Equal(t, v1.Pointer(), v2.Pointer())
}
