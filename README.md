# GO-TRACING

Go-tracing provides an easy way to use zipkin tracing with only four lines of code.

## Quick start

```
// set your tracer
tracing.SetGlobalTracer(appName, "{zipkin_host}")
defer tracing.FlushCollector()

// define a trace
span, ctx := tracing.CreateSpan(ctx, "{span_name}", nil)
defer span.Finish()
```

## Examples

```
docker run -d -p 9411:9411 openzipkin/zipkin
go run examples/basic/main.go
go run examples/httpServer/main.go
```

To watch traces you just have to hit http://localhost:9411/

## Features

 - Create span from nothing
 - Create span from context
 - Extract/Inject span from/to httpRequest
 - Extract/Inject span from/to a map[string]string (textMapCarrier)
 - Declare an error span

## License
go-tracing is licensed under the MIT license. (http://opensource.org/licenses/MIT)

## Contributing
Pull requests are the way to help us here. We will be really gratefull.