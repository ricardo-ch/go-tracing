EXAMPLE_APP ?= httpServer

.PHONY: build
build:
	 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo

.PHONY: test
test:
	go test ./...

.PHONY: run-jaeger
run-jaeger:
	@wget -nc https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
	@chmod +x wait-for-it.sh
	docker-compose up -d jaeger
	@./wait-for-it.sh localhost:5778 -- echo "jaeger ready"
