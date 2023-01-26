install:
	go install github.com/bufbuild/buf/cmd/buf

generate: build
build:
	go install ./...
	buf generate

lint:
	buf lint

format:
	buf format proto -w
