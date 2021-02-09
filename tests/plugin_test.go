package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gormopentracing "github.com/yeqown/gorm-opentracing"
)

func Test_usePlugin(t *testing.T) {
	plg := gormopentracing.New()
	err := DB.Use(plg)

	require.Nil(t, err)
	assert.Contains(t, DB.Plugins, plg.Name())
}
