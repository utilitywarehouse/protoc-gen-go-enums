BUF_VERSION := v1.0.0-rc12

install:
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go install ./...

build:
	buf generate

ci:
	go install ./...
	buf generate
