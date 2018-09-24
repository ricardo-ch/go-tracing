EXAMPLE_APP ?= httpServer

.PHONY: build
build:
	 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo

.PHONY: test
test:
	go test -v `go list ./... | grep -v /examples/`


docker-compose:
	@CGO_ENABLED=0 GOOS=linux go build -o ./app -a -ldflags '-s' -installsuffix cgo examples/$(EXAMPLE_APP)/main.go
	@docker-compose build
	@docker-compose up