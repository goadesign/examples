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
GOA:=$(shell goa version 2> /dev/null)
GOOS=$(shell go env GOOS)
GOPATH=$(shell go env GOPATH)

export GO111MODULE=on

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	honnef.co/go/tools/cmd/staticcheck

.phony: all depend lint test build clean

all: check-goa gen lint test
	@echo DONE!

ci: depend all

# Install protoc
PROTOC_VERSION=3.14.0
UNZIP=unzip
ifeq ($(GOOS),linux)
	PROTOC=protoc-$(PROTOC_VERSION)-linux-x86_64
	PROTOC_EXEC=$(PROTOC)/bin/protoc
endif
ifeq ($(GOOS),darwin)
	PROTOC=protoc-$(PROTOC_VERSION)-osx-x86_64
	PROTOC_EXEC=$(PROTOC)/bin/protoc
endif
ifeq ($(GOOS),windows)
	PROTOC=protoc-$(PROTOC_VERSION)-win32
	PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
	GOPATH:=$(subst \,/,$(GOPATH))
endif

check-goa:
ifdef GOA
	go mod download
	@echo $(GOA)
else
	go get -u goa.design/goa/v3@v3
	go get -u goa.design/goa/v3/...@v3
	go mod download
	@echo $(GOA)
endif

# Note: the steps below rely on curl and tar which are available
# on both Linux and Windows 10 (build>=17603).
depend:
	@echo INSTALLING DEPENDENCIES...
	@go mod download
	@env go get -v $(DEPEND)
	@echo INSTALLING PROTOC...
	@mkdir $(PROTOC)
	@cd $(PROTOC); \
	curl -O -L https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC).zip; \
	$(UNZIP) $(PROTOC).zip
	@cp $(PROTOC_EXEC) $(GOPATH)/bin && \
		rm -r $(PROTOC) && \
		echo "`protoc --version`"
	@echo go mod graph

lint:
	@echo LINTING CODE...
	@if [ "`goimports -l $(GO_FILES) | grep -v .pb.go | tee /dev/stderr`" ]; then \
		echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
	fi
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi
	@if [ "`staticcheck -checks all,-ST1000,-ST1001,-ST1021 ./... | grep -v ".pb.go" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi

gen:
	@# NOTE: not all command line tools are generated
	@echo GENERATING CODE...
	@goa version
	@rm -rf "$(GOPATH)/src/goa.design/examples/basic/cmd/calc-cli"
	@rm -rf "$(GOPATH)/src/goa.design/examples/cellar/cmd/cellar-cli"
	@rm -rf "$(GOPATH)/src/goa.design/examples/cookies/cmd/"
	@rm -rf "$(GOPATH)/src/goa.design/examples/encodings/text/cmd"
	@rm -rf "$(GOPATH)/src/goa.design/examples/error/cmd"
	@rm -rf "$(GOPATH)/src/goa.design/examples/files/cmd"
	@rm -rf "$(GOPATH)/src/goa.design/examples/multipart/cmd"
	@rm -rf "$(GOPATH)/src/goa.design/examples/security/cmd"
	@rm -rf "$(GOPATH)/src/goa.design/examples/streaming/cmd/chatter"
	@rm -rf "$(GOPATH)/src/goa.design/examples/tus/cmd/upload-cli"
	@rm -rf "$(GOPATH)/src/goa.design/examples/upload_download/cmd/upload_download-cli"
	@echo "basic [1/12]"
	@goa gen goa.design/examples/basic/design -o "$(GOPATH)/src/goa.design/examples/basic"
	@goa example goa.design/examples/basic/design -o "$(GOPATH)/src/goa.design/examples/basic"
	@echo "cellar [2/12]"
	@goa gen goa.design/examples/cellar/design -o "$(GOPATH)/src/goa.design/examples/cellar"
	@goa example goa.design/examples/cellar/design -o "$(GOPATH)/src/goa.design/examples/cellar"
	@echo "cookies [3/12]"
	@goa gen goa.design/examples/cookies/design -o "$(GOPATH)/src/goa.design/examples/cookies"
	@goa example goa.design/examples/cookies/design -o "$(GOPATH)/src/goa.design/examples/cookies"
	@echo "encodings/cbor [4/12]"
	@goa gen goa.design/examples/encodings/cbor/design -o "$(GOPATH)/src/goa.design/examples/encodings/cbor"
	@goa example goa.design/examples/encodings/cbor/design -o "$(GOPATH)/src/goa.design/examples/encodings/cbor"
	@echo "encodings/text [5/12]"
	@goa gen goa.design/examples/encodings/text/design -o "$(GOPATH)/src/goa.design/examples/encodings/text"
	@goa example goa.design/examples/encodings/text/design -o "$(GOPATH)/src/goa.design/examples/encodings/text"
	@echo "error [6/12]"
	@goa gen goa.design/examples/error/design -o "$(GOPATH)/src/goa.design/examples/error"
	@goa example goa.design/examples/error/design -o "$(GOPATH)/src/goa.design/examples/error"
	@echo "files [7/12]"
	@goa gen goa.design/examples/files/design -o "$(GOPATH)/src/goa.design/examples/files"
	@goa example goa.design/examples/files/design -o "$(GOPATH)/src/goa.design/examples/files"
	@echo "multipart [8/12]"
	@goa gen goa.design/examples/multipart/design -o "$(GOPATH)/src/goa.design/examples/multipart"
	@goa example goa.design/examples/multipart/design -o "$(GOPATH)/src/goa.design/examples/multipart"
	@echo "security [9/12]"
	@goa gen goa.design/examples/security/design -o "$(GOPATH)/src/goa.design/examples/security"
	@goa example goa.design/examples/security/design -o "$(GOPATH)/src/goa.design/examples/security"
	@echo "streaming [10/12]"
	@goa gen goa.design/examples/streaming/design -o "$(GOPATH)/src/goa.design/examples/streaming"
	@goa example goa.design/examples/streaming/design -o "$(GOPATH)/src/goa.design/examples/streaming"
	@echo "tus [11/12]"
	@goa gen goa.design/examples/tus/design -o "$(GOPATH)/src/goa.design/examples/tus"
	@goa example goa.design/examples/tus/design -o "$(GOPATH)/src/goa.design/examples/tus"
	@echo "upload_download [12/12]"
	@goa gen goa.design/examples/upload_download/design -o "$(GOPATH)/src/goa.design/examples/upload_download"
	@goa example goa.design/examples/upload_download/design -o "$(GOPATH)/src/goa.design/examples/upload_download"
	@go mod tidy -compat=1.17

build:
	@cd "$(GOPATH)/src/goa.design/examples/basic" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GOPATH)/src/goa.design/examples/cellar" && \
		go build ./cmd/cellar && go build ./cmd/cellar-cli
	@cd "$(GOPATH)/src/goa.design/examples/cookies" && \
		go build ./cmd/session && go build ./cmd/session-cli
	@cd "$(GOPATH)/src/goa.design/examples/encodings/cbor" && \
		go build ./cmd/concat && go build ./cmd/concat-cli
	@cd "$(GOPATH)/src/goa.design/examples/encodings/text" && \
		go build ./cmd/text && go build ./cmd/text-cli
	@cd "$(GOPATH)/src/goa.design/examples/error" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GOPATH)/src/goa.design/examples/files" && \
		go build ./cmd/openapi && go build ./cmd/openapi-cli
	@cd "$(GOPATH)/src/goa.design/examples/multipart" && \
		go build ./cmd/resume && go build ./cmd/resume-cli
	@cd "$(GOPATH)/src/goa.design/examples/security" && \
		go build ./cmd/multi_auth && go build ./cmd/multi_auth-cli
	@cd "$(GOPATH)/src/goa.design/examples/streaming" && \
		go build ./cmd/chatter && go build ./cmd/chatter-cli
	@cd "$(GOPATH)/src/goa.design/examples/tracing" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GOPATH)/src/goa.design/examples/tus" && \
		go build ./cmd/upload && go build ./cmd/upload-cli
	@cd "$(GOPATH)/src/goa.design/examples/upload_download" && \
		go build ./cmd/upload_download && go build ./cmd/upload_download-cli

test:
	@echo TESTING...
	@go test ./... > /dev/null
