.PHONY=fmt
BINARY_NAME=eeyore
GOPATH := $(shell go env GOPATH)
GOBIN ?= $(GOPATH)/bin
PREFIX=~/.local/bin
BUILD_DIR=./build

GO := $(shell command -v go 2> /dev/null)

GOLANGCILINT := $(GOBIN)/golangci-lint
$(GOLANGCILINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.0

RICHGO := $(GOBIN)/richgo
$(RICHGO):
	@go install github.com/kyoh86/richgo@v0.3.6

GOSEC := $(GOBIN)/gosec
$(GOSEC):
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.13.1

GOFUMPT := $(GOBIN)/gofumpt
$(GOFUMPT):
	$(GO) install mvdan.cc/gofumpt@v0.1.1

GOCRITIC := $(GOBIN)/gocritic
$(GOCRITIC):
	$(GO) install github.com/go-critic/go-critic/cmd/gocritic@v0.6.5

dependencies: $(GOLANGCILINT) $(RICHGO) $(GOSEC) $(GOFUMPT) $(GOCRITIC)

fmt: $(GOFUMPT) $(GOCRITIC)
	$(GO) fmt ./...
	goimports -w .
	$(GOFUMPT) -s -w .
	$(GOCRITIC) check ./...


lint: $(GOLANGCILINT)
	golangci-lint run

test: $(RICHGO)
	@$(RICHGO) test -v ./...

security: $(GOSEC)
	gosec -quiet ./...

check: dependencies
	pre-commit run -a

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

install:
	test -f $(BUILD_DIR)/$(BINARY_NAME)
	mkdir -p $(PREFIX)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(PREFIX)

uninstall:
	rm $(PREFIX)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)
