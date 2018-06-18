package tracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setAndGetTracer_OK(t *testing.T) {
	SetGlobalTracer("testAPIName", "http://testHost:9411")
	tracer := GetGlobalTracer()
	assert.NotNil(t, tracer)
}
