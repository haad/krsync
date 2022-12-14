 .PHONY: build install clean test integration modverify modtidy

VERSION=`egrep -o '[0-9]+\.[0-9a-z.\-]+' version.go`
GIT_SHA=`git rev-parse --short HEAD || echo`

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

build:
	@echo "Building krsync..."
	@mkdir -p bin
	@go build -ldflags "-X cmd.GitSHA=${GIT_SHA}" -o bin/krsync .

install:
	@echo "Installing krsync..."
	@install -c bin/krsync /usr/local/bin/krsync

clean:
	@rm -f bin/*

test:
	@echo "Running tests..."
	#@go test `go list ./... | grep -v vendor/`

fmt:
	@gofmt -l -w -s $(SRC)

integration: modtidy build test
	@echo "Running integration tests..."
	#bash integration/run.sh
	#@docker run -it --rm -v $(pwd):/go/src/github.com/haad/confd golang:1.18.2 /go/src/github.com/haad/confd/integration/run.sh

modtidy:
	@go mod tidy

modverify:
	@go mod verify
