.PHONY: build
build:
	 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo

.PHONY: test
test:
	go test -v `go list ./... | grep -v /examples/`