install:
	go install github.com/bufbuild/buf/cmd/buf

generate: build
build:
	buf generate

lint:
	buf lint

format:
	buf format proto -w
