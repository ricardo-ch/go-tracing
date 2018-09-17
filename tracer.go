package tracing

import (
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var globalCloser io.Closer

func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

func SetGlobalTracer(appName string) error {
	// Recommended configuration for production.
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		log.Println("could not load jaeger configuration from environment variables")
		return nil
	}

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		appName,
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return err
	}
	globalCloser = closer

	return err
}

func FlushCollector() {
	if globalCloser != nil {
		globalCloser.Close()
	}
}
