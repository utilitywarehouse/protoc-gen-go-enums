BUILDENV :=
BUILDENV += CGO_ENABLED=0
BUILDENV += GO111MODULE=on
BUILDENV += GOPRIVATE="github.com/utilitywarehouse/*"

LINTER_EXE := golangci-lint
LINTER := $(GOPATH)/bin/$(LINTER_EXE)

$(LINTER):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(GOPATH)/bin

.PHONY: lint
lint: $(LINTER)
	$(LINTER) run ./...

.PHONY: install
install:
	$(BUILDENV) go install ./...

.PHONY: test
test:
	$(BUILDENV) CGO_ENABLED=1 go test -v -cover -race ./...

.PHONY: all
all: $(LINTER) lint test install

.PHONY: generate
generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	$(BUILDENV) go generate ./...
