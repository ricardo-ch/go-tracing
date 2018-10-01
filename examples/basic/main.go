package main

import (
	"context"
	"time"

	"github.com/ricardo-ch/go-tracing"
	"log"
	"os"
)

// this name is use to identify traces inside zipkin UI
const (
	appName = "my-application"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", appName)
	os.Setenv("JAEGER_AGENT_HOST", "localhost")
	os.Setenv("JAEGER_AGENT_PORT", "6831")
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")

	err := tracing.SetGlobalTracer()
	if err != nil {
		log.Fatal(err)
	}

	defer tracing.FlushCollector()

	doWork(context.Background())
}

func doWork(ctx context.Context) {
	span, ctx := tracing.CreateSpan(ctx, "doWork", nil)
	defer span.Finish()

	time.Sleep(2 * time.Second)
}
