# GO-TRACING
[![Build Status](https://travis-ci.org/ricardo-ch/go-tracing.svg?branch=travis)](https://travis-ci.org/ricardo-ch/go-tracing)
[![Coverage Status](https://coveralls.io/repos/github/ricardo-ch/go-tracing/badge.svg?branch=master)](https://coveralls.io/github/ricardo-ch/go-tracing?branch=master)

Go-tracing provides an easy way to use zipkin tracing with only four lines of code.

## Quick start

```golang
// import the library
import "github.com/ricardo-ch/go-tracing"

// set your tracer
tracing.SetGlobalTracer(appName, "{zipkin_url}")
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
go run examples/httpServer-middleware/main.go
go run examples/httpGoKit-middleware/main.go
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
Pull requests are the way to help us here. We will be really grateful.
