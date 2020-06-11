package tracing

import (
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var globalFlusher func()

func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

type tracerConfig struct {
	TracingService string
	AppName        string
	TracingHost    string
}

type tracerOption func(config *tracerConfig)

// Pass UsingJaeger's result as argument to SetGlobalTracer to set Jaeger as your tracing system
// This is the default behavior
func UsingJaeger() tracerOption {
	return func(config *tracerConfig) {
		config.TracingService = "jaeger"
	}
}

func SetGlobalTracer(options ...tracerOption) error {
	config := &tracerConfig{}
	for _, option := range options {
		option(config)
	}

	setGlobalJaegerTracer()
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
