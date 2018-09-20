package tracing

import (
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var globalCloser io.Closer

func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

func SetGlobalTracer() error {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return err
	} else {
		var tracer opentracing.Tracer
		tracer, globalCloser, err = cfg.New(
			cfg.ServiceName,
		)
		if err != nil {
			return err
		}
		opentracing.SetGlobalTracer(tracer)
	}
	return err
}

func FlushCollector() {
	if globalCloser != nil {
		globalCloser.Close()
	}
}
