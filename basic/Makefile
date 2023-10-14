#! /usr/bin/make
#
# Makefile for Goa examples
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "gen" invokes the "goa" tool to generate the examples source code
# - "build" compiles the example microservices and client CLIs
# - "clean" deletes the output of "build"
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
GO_FILES=$(shell find . -type f -name '*.go')
MODULE=$(shell go list -m)
APP=calc

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	github.com/hashicorp/go-getter \
	github.com/cheggaaa/pb \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/golang/protobuf/proto \
	goa.design/goa/... \
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	honnef.co/go/tools/cmd/staticcheck

.phony: all depend lint test build clean

all: gen lint test
	@echo DONE!

# Install protoc
GOOS=$(shell go env GOOS)
PROTOC_VERSION=3.6.1
ifeq ($(GOOS),linux)
PROTOC=protoc-$(PROTOC_VERSION)-linux-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
GOBIN=$(GOPATH)/bin
else
	ifeq ($(GOOS),darwin)
PROTOC=protoc-$(PROTOC_VERSION)-osx-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
GOBIN=$(GOPATH)/bin
	else
		ifeq ($(GOOS),windows)
PROTOC=protoc-$(PROTOC_VERSION)-win32
PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
GOBIN="$(GOPATH)\bin"
		endif
	endif
endif
depend:
	@echo INSTALLING DEPENDENCIES...
	@go get -v $(DEPEND)
	@go install github.com/hashicorp/go-getter/cmd/go-getter && \
		go-getter https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC).zip $(PROTOC) && \
		cp $(PROTOC_EXEC) $(GOBIN) && \
		rm -r $(PROTOC)
	@go install github.com/golang/protobuf/protoc-gen-go
	@go get -t -v ./...

lint:
	@echo LINTING CODE...
	@if [ "`goimports -l $(GO_FILES) | grep -v .pb.go | tee /dev/stderr`" ]; then \
		echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
	fi
	@if [ "`staticcheck -checks all,-ST1000,-ST1001 ./... | grep -v ".pb.go" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi

.PHONY: gen
gen:
	@# NOTE: not all command line tools are generated
	@echo GENERATING CODE...
	goa gen     "$(MODULE)/design" && \
	goa example "$(MODULE)/design"

build:
	@go build "./cmd/$(APP)" && go build "./cmd/$(APP)-cli"

clean:
	@rm -rf "./cmd/$(APP)" "./cmd/$(APP)-cli"

test:
	@echo TESTING...
	@go test ./... > /dev/null

check-freshness:
	@if [ "`git diff | wc -l`" -gt "0" ]; then \
	        echo "[ERROR] generated code not in-sync with design:"; \
	        echo; \
	        git status -s; \
	        git --no-pager diff; \
	        echo; \
	        exit 1; \
	fi
