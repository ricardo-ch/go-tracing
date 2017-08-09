package tracing

import (
	"context"
	"net/http"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

//CreateSpan ...
func CreateSpan(ctx context.Context, spanName string, tags *map[string]interface{}) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, spanName)

	setTags(span, tags)

	return span, ctx
}

//CreateSpanFromClientContext ...
func CreateSpanFromClientContext(r *http.Request, spanName string, tags *map[string]interface{}) (opentracing.Span, context.Context) {
	tracer := GetGlobalTracer()

	wireContext, _ := tracer.Extract(
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(r.Header),
	)

	// create span
	span := tracer.StartSpan(spanName, ext.RPCServerOption(wireContext))

	setTags(span, tags)

	// store span in context
	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	childCtx := opentracing.ContextWithSpan(ctx, span)

	return span, childCtx
}

//SetSpanError ...
func SetSpanError(span opentracing.Span, err error) {
	setTags(span, &map[string]interface{}{"error": err})
}

func setTags(span opentracing.Span, tags *map[string]interface{}) {
	if tags != nil {
		for k, v := range *tags {
			span.SetTag(k, v)
		}
	}
}

// InjectIntoCarrier returns a textMapCarrier, basically a map[string]string,
//  which can be used to transmit a span context to another service with ExtractFromCarrier
func InjectIntoCarrier(ctx context.Context) opentracing.TextMapCarrier{
	carrier := opentracing.TextMapCarrier{}

	// Retrieve the Span from context
	if span := opentracing.SpanFromContext(ctx); span != nil {
		// We are going to use this span in a client request, so mark as such.
		ext.SpanKindProducer.Set(span)
		// Retrieve tracer
		tracer := opentracing.GlobalTracer()
		// Inject the Span context into the outgoing HTTP Request
		tracer.Inject(
			span.Context(),
			opentracing.TextMap,
			carrier,
		)
	}
	return carrier
}

// ExtractFromCarrier returns a span with context passed  by the carrier
// ctx should not already have span in it
func ExtractFromCarrier(ctx context.Context, carrier opentracing.TextMapCarrier, spanName string, tags *map[string]interface{}) (opentracing.Span, context.Context) {
	tracer := GetGlobalTracer()

	wireContext, _ := tracer.Extract(
		opentracing.TextMap,
		carrier,
	)

	span := tracer.StartSpan(spanName, ext.RPCServerOption(wireContext))
	setTags(span, tags)

	// store span in context
	if ctx == nil {
		ctx = context.Background()
	}
	childCtx := opentracing.ContextWithSpan(ctx, span)

	return span, childCtx
}

//InjectSpan ...
func InjectSpan(r *http.Request) *http.Request {
	// Retrieve the Span from context
	if span := opentracing.SpanFromContext(r.Context()); span != nil {
		// We are going to use this span in a client request, so mark as such.
		ext.SpanKindRPCClient.Set(span)
		// Retrieve tracer
		tracer := opentracing.GlobalTracer()
		// Inject the Span context into the outgoing HTTP Request
		tracer.Inject(
			span.Context(),
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
	}
	return r
}