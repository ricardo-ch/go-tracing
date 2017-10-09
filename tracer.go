package tracing

import (
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/pkg/errors"
)

var globalColector zipkin.Collector

//GetGlobalTracer ...
func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

//SetGlobalTracer ...
func SetGlobalTracer(apiName string, zipkinURL string) error {
	tracer, err := createTracer(apiName, zipkinURL)
	if err != nil {
		return errors.Wrap(err, "SetGlobalTracer")
	}
	opentracing.SetGlobalTracer(tracer)
	return nil
}

func createTracer(apiName string, zipkinURL string) (opentracing.Tracer, error) {
	collector, err := createCollector(zipkinURL)
	if err != nil {
		return nil, errors.Wrap(err, "createTracer")
	}

	tracer, err := zipkin.NewTracer(
		createRecorder(collector, apiName),
		zipkin.ClientServerSameSpan(true),
	)

	if err != nil {
		return nil, errors.Wrap(err, "createTracer")
	}

	return tracer, nil
}

func createCollector(zipkinURL string) (zipkin.Collector, error) {
	url := zipkinURL + "/api/v1/spans"
	collector, err := zipkin.NewHTTPCollector(url)

	if err != nil {
		return nil, errors.Wrap(err, "createCollector")
	}
	globalColector = collector
	return collector, nil
}

func createRecorder(collector zipkin.Collector, apiName string) zipkin.SpanRecorder {
	return zipkin.NewRecorder(collector,
		false,
		"0.0.0.0:0",
		apiName)
}

func FlushCollector() {
	if globalColector != nil {
		globalColector.Close()
	}
}
