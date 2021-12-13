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

func Test_Option_WithCreateOpName(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ApplyOption
		wantRes operationName
	}{
		{
			name:    "default",
			wantRes: _createOp,
		},
		{
			name:    "empty",
			opts:    []ApplyOption{WithCreateOpName("")},
			wantRes: _createOp,
		},
		{
			name:    "custom",
			opts:    []ApplyOption{WithCreateOpName("gorm_create")},
			wantRes: operationName("gorm_create"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := New(tc.opts...)
			assert.Equal(t, tc.wantRes, p.(opentracingPlugin).opt.createOpName)
		})
	}
}

func Test_Option_WithUpdateOpName(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ApplyOption
		wantRes operationName
	}{
		{
			name:    "default",
			wantRes: _updateOp,
		},
		{
			name:    "empty",
			opts:    []ApplyOption{WithUpdateOpName("")},
			wantRes: _updateOp,
		},
		{
			name:    "custom",
			opts:    []ApplyOption{WithUpdateOpName("gorm_update")},
			wantRes: operationName("gorm_update"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := New(tc.opts...)
			assert.Equal(t, tc.wantRes, p.(opentracingPlugin).opt.updateOpName)
		})
	}
}

func Test_Option_WithQueryOpName(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ApplyOption
		wantRes operationName
	}{
		{
			name:    "default",
			wantRes: _queryOp,
		},
		{
			name:    "empty",
			opts:    []ApplyOption{WithQueryOpName("")},
			wantRes: _queryOp,
		},
		{
			name:    "custom",
			opts:    []ApplyOption{WithQueryOpName("gorm_query")},
			wantRes: operationName("gorm_query"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := New(tc.opts...)
			assert.Equal(t, tc.wantRes, p.(opentracingPlugin).opt.queryOpName)
		})
	}
}

func Test_Option_WithDeleteOpName(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ApplyOption
		wantRes operationName
	}{
		{
			name:    "default",
			wantRes: _deleteOp,
		},
		{
			name:    "empty",
			opts:    []ApplyOption{WithDeleteOpName("")},
			wantRes: _deleteOp,
		},
		{
			name:    "custom",
			opts:    []ApplyOption{WithDeleteOpName("gorm_delete")},
			wantRes: operationName("gorm_delete"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := New(tc.opts...)
			assert.Equal(t, tc.wantRes, p.(opentracingPlugin).opt.deleteOpName)
		})
	}
}

func Test_Option_WithRowOpName(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ApplyOption
		wantRes operationName
	}{
		{
			name:    "default",
			wantRes: _rowOp,
		},
		{
			name:    "empty",
			opts:    []ApplyOption{WithRowOpName("")},
			wantRes: _rowOp,
		},
		{
			name:    "custom",
			opts:    []ApplyOption{WithRowOpName("gorm_row")},
			wantRes: operationName("gorm_row"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := New(tc.opts...)
			assert.Equal(t, tc.wantRes, p.(opentracingPlugin).opt.rowOpName)
		})
	}
}

func Test_Option_WithRawOpName(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ApplyOption
		wantRes operationName
	}{
		{
			name:    "default",
			wantRes: _rawOp,
		},
		{
			name:    "empty",
			opts:    []ApplyOption{WithRawOpName("")},
			wantRes: _rawOp,
		},
		{
			name:    "custom",
			opts:    []ApplyOption{WithRawOpName("gorm_raw")},
			wantRes: operationName("gorm_raw"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := New(tc.opts...)
			assert.Equal(t, tc.wantRes, p.(opentracingPlugin).opt.rawOpName)
		})
	}
}
