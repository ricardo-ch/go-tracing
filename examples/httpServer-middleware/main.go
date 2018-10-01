package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/ricardo-ch/go-tracing"
	"os"
)

// this name is use to identify traces inside zipkin
const (
	appName = "my-http-server-middleware"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", appName)
	os.Setenv("JAEGER_AGENT_HOST", "localhost")
	os.Setenv("JAEGER_AGENT_PORT", "6831")
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")

	tracing.SetGlobalTracer()
	defer tracing.FlushCollector()

	http.Handle("/", tracing.HTTPMiddleware("hello-handler", http.HandlerFunc(hello)))
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	nestedFunc(r.Context())

	io.WriteString(w, "Hello world!")
}

func nestedFunc(ctx context.Context) {
	span, ctx := tracing.CreateSpan(ctx, "nestedFunc", nil)
	defer span.Finish()

	time.Sleep(1 * time.Second)
}
