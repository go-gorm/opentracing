package gormopentracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Option_WithLogResult(t *testing.T) {
	opt := WithLogResult(true)
	p := New(opt)
	assert.Equal(t, true, p.(opentracingPlugin).logResult)
}
