package main

import (
	"context"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ricardo-ch/go-tracing"
)

// this name is use to identify traces inside zipkin UI
const (
	appName = "my-application"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", appName)
	tracing.SetGlobalTracer()
	defer tracing.FlushCollector()

	doWork(context.Background())
}

func doWork(ctx context.Context) {
	span, ctx := tracing.CreateSpan(ctx, "doWork", nil)
	defer span.Finish()

	r, err := googleRequest(ctx)
	if err != nil {
		os.Exit(-1)
	}

	option := func(trace *httptrace.ClientTrace, span opentracing.Span) {
		trace.GotFirstResponseByte = func() {
			span.LogKV("event", "Got First Response byte (changed)")
		}
	}

	// Add The trace to the Request
	r = tracing.InjectSpan(r, option)
	client := &http.Client{
		Timeout: time.Duration(100) * (time.Millisecond),
	}
	client.Do(r)
}

func googleRequest(ctx context.Context) (r *http.Request, er error) {
	r, er = http.NewRequest("GET", "http://www.google.ca/", nil) // ;)
	if er != nil {
		return nil, er

	}
	r = r.WithContext(ctx)

	return r, nil
}
