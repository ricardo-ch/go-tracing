package main

import (
	"context"
	"time"

	"github.com/ricardo-ch/go-tracing"
	"os"
)

// this name is use to identify traces inside zipkin UI
const (
	appName = "my-application"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", appName)
	err := tracing.SetGlobalTracer()
	if err != nil {
		panic(err)
	}

	defer tracing.FlushCollector()

	doWork(context.Background())
}

func doWork(ctx context.Context) {
	span, ctx := tracing.CreateSpan(ctx, "doWork", nil)
	defer span.Finish()

	time.Sleep(2 * time.Second)
}
