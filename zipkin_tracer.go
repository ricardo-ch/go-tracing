package tracing

import (
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/pkg/errors"
	"net/url"
)

func setGlobalZipkinTracer(apiName string, zipkinURL string) error {
	tracer, err := createZipkinTracer(apiName, zipkinURL)
	if err != nil {
		return errors.Wrap(err, "SetGlobalTracer")
	}
	opentracing.SetGlobalTracer(tracer)
	return nil
}

func createZipkinTracer(apiName string, zipkinURL string) (opentracing.Tracer, error) {
	collector, err := createCollector(zipkinURL)
	if err != nil {
		return nil, errors.Wrap(err, "createZipkinTracer")
	}

	tracer, err := zipkin.NewTracer(
		createRecorder(collector, apiName),
		zipkin.ClientServerSameSpan(true),
	)

	if err != nil {
		return nil, errors.Wrap(err, "createZipkinTracer")
	}

	return tracer, nil
}

func createCollector(zipkinURL string) (zipkin.Collector, error) {
	relativeEndPointURL, _ := url.Parse("api/v1/spans")
	serviceURL, err := url.Parse(zipkinURL)
	if err != nil {
		return nil, err
	}
	absoluteEndPointURL := serviceURL.ResolveReference(relativeEndPointURL)

	collector, err := zipkin.NewHTTPCollector(absoluteEndPointURL.String())

	if err != nil {
		return nil, errors.Wrap(err, "createCollector")
	}

	globalFlusher = func() {
		if collector != nil {
			collector.Close()
		}
	}

	return collector, nil
}

func createRecorder(collector zipkin.Collector, apiName string) zipkin.SpanRecorder {
	return zipkin.NewRecorder(collector,
		false,
		"0.0.0.0:0",
		apiName)
}
