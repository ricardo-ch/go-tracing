package tracing

import (
	"fmt"
	"net/http"
	"github.com/go-kit/kit/endpoint"
	"context"
)

// HTTPMiddleware returns a Middleware that injects an OpenTracing Span found in
// context into the HTTP Headers.
func HTTPMiddleware(operationName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tags := make(map[string]interface{})

		// set whole URL as tag
		tags["url"] = req.URL.String()

		// set request parameters as tags
		parameters := req.URL.Query()
		for key, values := range parameters {
			if len(values) == 1 {
				tags[key] = values[0]
			} else {
				for i, value := range values {
					tags[fmt.Sprintf("%s[%d]", key, i)] = value
				}
			}
		}

		span, ctx := CreateSpanFromClientContext(req, operationName, &tags)
		defer span.Finish()

		// update request context to include our new span
		req = req.WithContext(ctx)

		// next middleware or actual request handler
		next.ServeHTTP(w, req)
	})
}

// GotKitEndpointMiddleWare returns a gokit.Middleware which change the behavior of a gokit.endpoint
// it had tracing capability
func GotKitEndpointMiddleWare(operationName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			span, ctx := CreateSpan(ctx, operationName, nil)
			defer span.Finish()
			return next(ctx, request)
		}
	}
}
