package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/ricardo-ch/go-tracing"
)

// this name is use to identify traces inside zipkin
const (
	appName = "my-http-server"
)

func main() {
	tracing.SetGlobalTracer(appName)
	defer tracing.FlushCollector()

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracing.CreateSpanFromClientContext(r, "hello-handler", nil)
	defer span.Finish()

	time.Sleep(1 * time.Second)
	nestedFunc(ctx)

	io.WriteString(w, "Hello world!")
}

func nestedFunc(ctx context.Context) {
	span, ctx := tracing.CreateSpan(ctx, "nestedFunc", nil)
	defer span.Finish()

	time.Sleep(1 * time.Second)
}
