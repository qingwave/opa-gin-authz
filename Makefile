.PHONY: build run test clean

build: clean
	go build -mod vendor -o bin/opa-gin-authz main.go

run:
	go run -mod vendor main.go

test:
	go test -ldflags -s -v --cover $(shell go list ./...)
	opa test -v authz

clean:
	@rm -rf bin/
	go mod tidy
	go mod vendor
