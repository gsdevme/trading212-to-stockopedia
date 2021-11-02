.PHONY: all
default: build;

fmt:
	go fmt ./...

lint:
	 golangci-lint run

tests:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

run:
	./dist/trading212-to-stockopedia_darwin_amd64/trading212-to-stockopedia

build:
	goreleaser build --snapshot --skip-validate --rm-dist --single-target


all: build run