package tracing

import (
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var globalFlusher func()

func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

type TracerConfig struct {
	TracingService string
	AppName        string
	TracingHost    string
}

type TracerOption func(config *TracerConfig)

func UsingZipkin(appName string, host string) TracerOption {
	return func(config *TracerConfig) {
		config.TracingService = "zipkin"
		config.TracingHost = host
		config.AppName = appName
	}
}

func UsingJaeger() TracerOption {
	return func(config *TracerConfig) {
		config.TracingService = "jaeger"
	}
}

func SetGlobalTracer(options ...TracerOption) error {
	config := &TracerConfig{}
	for _, option := range options {
		option(config)
	}

	switch config.TracingService {
	case "zipkin":
		setGlobalZipkinTracer(config.AppName, config.TracingHost)
	case "jaeger":
	default:
		setGlobalJaegerTracer()
	}
	return nil
}

func FlushCollector() {
	if globalFlusher != nil {
		globalFlusher()
	}
}

func setGlobalJaegerTracer() error {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return err
	}

	tracer, globalCloser, err := cfg.NewTracer()
	if err != nil {
		return err
	}
	opentracing.SetGlobalTracer(tracer)

	globalFlusher = func() {
		if globalCloser != nil {
			globalCloser.Close()
		}
	}

	return err
}
