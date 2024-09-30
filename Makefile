BINARY_NAME=./bin/client
GO_MAIN=./main.go

build:
	@go build 

run:
	@go build
	@go install
	@lockbox -h sjdoo.zapto.org


dev:
	@go build
	@go install
	@lockbox -h "localhost:1337" -u "test"

clean:
	@rm lockbox

install:
	@go build
	@go install

.PHONY: build run install clean
