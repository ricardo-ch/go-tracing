package tracing

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_setAndGetTracer_OK(t *testing.T) {
	SetGlobalTracer("testAPIName","testHost")
	tracer := GetGlobalTracer()
	assert.NotNil(t, tracer)
}
