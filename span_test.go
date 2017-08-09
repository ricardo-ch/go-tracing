package tracing

import (
	"context"
	"errors"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func Test_createSpan_OK(t *testing.T) {
	span, ctx := CreateSpan(context.Background(), "testSpan", &map[string]interface{}{})
	assert.NotNil(t, span)
	assert.NotNil(t, ctx)
}

func Test_createSpan_fromClientContext_OK(t *testing.T) {
	span, ctx := CreateSpanFromClientContext(&http.Request{}, "testSpanFromClient", &map[string]interface{}{})
	assert.NotNil(t, span)
	assert.NotNil(t, ctx)
}

func Test_setSpanError_OK(t *testing.T) {
	span, ctx := CreateSpan(context.Background(), "testSpan", &map[string]interface{}{})
	assert.NotNil(t, span)
	assert.NotNil(t, ctx)

	SetSpanError(span, errors.New("errorSpan"))
}
