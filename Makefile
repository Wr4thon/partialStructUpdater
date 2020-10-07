all: clean download generate test build
.PHONY: all

download:
	go mod download

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
	go build -installsuffix cgo -o partialstructupdater ./cmd/main.go

generate:
	go generate ./...

test:
	go test -short -cover -v -coverprofile=coverage.out -covermode=atomic ./...

clean:
	go clean
	go mod tidy

docker:
	docker run --rm -v ~/.netrc:/root/.netrc -v $(shell pwd):/src -w /src golang:1.15 make build
