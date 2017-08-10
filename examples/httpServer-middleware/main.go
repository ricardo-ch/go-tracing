package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/ricardo-ch/go-tracing"
	"github.com/ricardo-ch/go-tracing/examples/httpServer-middleware/middleware"
)

// this name is use to identify traces inside zipkin
const (
	appName = "my-http-server-middleware"
)

func main() {
	tracing.SetGlobalTracer(appName, "http://localhost:9411/")
	defer tracing.FlushCollector()

	http.Handle("/", middleware.TracingMiddleware("hello-handler", hello))
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
