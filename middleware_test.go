package tracing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(HTTPMiddleware("test-get", http.HandlerFunc(getTestHandler)))
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))

	res, err := http.Get(u.String())
	assert.NoError(err)
	if res != nil {
		defer res.Body.Close()
	}

	_, err = ioutil.ReadAll(res.Body)
	assert.NoError(err)
	assert.Equal(res.StatusCode, 200, "Span should be in the context")
}

func getTestHandler(rw http.ResponseWriter, req *http.Request) {
	if span := opentracing.SpanFromContext(req.Context()); span != nil {
		fmt.Fprint(rw, "span ok")
	}
	http.Error(rw, "SPAN_NOT_FOUND", 404)
}
