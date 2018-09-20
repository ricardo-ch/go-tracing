package tracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setAndGetTracer_OK(t *testing.T) {

	SetGlobalTracer()
	tracer := GetGlobalTracer()
	assert.NotNil(t, tracer)
}
