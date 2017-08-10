package middleware

import (
	"fmt"
	"net/http"

	"github.com/ricardo-ch/go-tracing"
)

// TracingMiddleware returns a Middleware that injects an OpenTracing Span found in
// context into the HTTP Headers.
func TracingMiddleware(operationName string, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// set request parameters as tags
		tags := make(map[string]interface{})
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

		span, ctx := tracing.CreateSpanFromClientContext(req, operationName, &tags)
		defer span.Finish()

		// update request context to include our new span
		req = req.WithContext(ctx)

		// next middleware or actual request handler
		next(w, req)
	})
}
