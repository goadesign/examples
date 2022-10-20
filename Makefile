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
GIT_ROOT=$(shell git rev-parse --show-toplevel)

export GO111MODULE=on

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	honnef.co/go/tools/cmd/staticcheck \
	goa.design/goa/v3/cmd/goa@v3

.phony: all depend lint test build clean

all: check-goa gen lint test
	@echo DONE!

ci: depend all

# Install protoc
PROTOC_VERSION=21.7
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
	GIT_ROOT:=$(subst \,/,$(GIT_ROOT))
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
		rm -rf $(PROTOC) && \
		echo "`protoc --version`"
	@echo go mod graph

lint:
	@echo LINTING CODE...
ifneq ($(GOOS),windows)
	@if [ "`staticcheck ./... | grep -v ".pb.go" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi
endif

gen:
	@# NOTE: not all command line tools are generated
	@echo GENERATING CODE...
	@goa version
	@cd $(GIT_ROOT)
	@rm -rf "basic/cmd/calc-cli"
	@rm -rf "cellar/cmd/cellar-cli"
	@rm -rf "cookies/cmd/"
	@rm -rf "encodings/text/cmd"
	@rm -rf "error/cmd"
	@rm -rf "files/cmd"
	@rm -rf "multipart/cmd"
	@rm -rf "security/hierarchy/cmd"
	@rm -rf "security/multiauth/cmd"
	@rm -rf "streaming/cmd/chatter"
	@rm -rf "tus/cmd/upload-cli"
	@rm -rf "upload_download/cmd/upload_download-cli"
	@echo "basic [1/13]"
	@goa gen goa.design/examples/basic/design -o "basic"
	@goa example goa.design/examples/basic/design -o "basic"
	@echo "cellar [2/13]"
	@goa gen goa.design/examples/cellar/design -o "cellar"
	@goa example goa.design/examples/cellar/design -o "cellar"
	@echo "cookies [3/13]"
	@goa gen goa.design/examples/cookies/design -o "cookies"
	@goa example goa.design/examples/cookies/design -o "cookies"
	@echo "encodings/cbor [4/13]"
	@goa gen goa.design/examples/encodings/cbor/design -o "encodings/cbor"
	@goa example goa.design/examples/encodings/cbor/design -o "encodings/cbor"
	@echo "encodings/text [5/13]"
	@goa gen goa.design/examples/encodings/text/design -o "encodings/text"
	@goa example goa.design/examples/encodings/text/design -o "encodings/text"
	@echo "error [6/13]"
	@goa gen goa.design/examples/error/design -o "error"
	@goa example goa.design/examples/error/design -o "error"
	@echo "files [7/13]"
	@goa gen goa.design/examples/files/design -o "files"
	@goa example goa.design/examples/files/design -o "files"
	@echo "multipart [8/13]"
	@goa gen goa.design/examples/multipart/design -o "multipart"
	@goa example goa.design/examples/multipart/design -o "multipart"
	@echo "security/hierarchy [9/13]"
	@goa gen goa.design/examples/security/hierarchy/design -o "security/hierarchy"
	@goa example goa.design/examples/security/hierarchy/design -o "security/hierarchy"
	@echo "security/multiauth [10/13]"
	@goa gen goa.design/examples/security/multiauth/design -o "security/multiauth"
	@goa example goa.design/examples/security/multiauth/design -o "security/multiauth"
	@echo "streaming [11/13]"
	@goa gen goa.design/examples/streaming/design -o "streaming"
	@goa example goa.design/examples/streaming/design -o "streaming"
	@echo "tus [12/13]"
	@goa gen goa.design/examples/tus/design -o "tus"
	@goa example goa.design/examples/tus/design -o "tus"
	@echo "upload_download [13/13]"
	@goa gen goa.design/examples/upload_download/design -o "upload_download"
	@goa example goa.design/examples/upload_download/design -o "upload_download"
	@go mod tidy -compat=1.19

build:
	@cd "$(GIT_ROOT)/basic" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GIT_ROOT)/cellar" && \
		go build ./cmd/cellar && go build ./cmd/cellar-cli
	@cd "$(GIT_ROOT)/cookies" && \
		go build ./cmd/session && go build ./cmd/session-cli
	@cd "$(GIT_ROOT)/encodings/cbor" && \
		go build ./cmd/concat && go build ./cmd/concat-cli
	@cd "$(GIT_ROOT)/encodings/text" && \
		go build ./cmd/text && go build ./cmd/text-cli
	@cd "$(GIT_ROOT)/error" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GIT_ROOT)/files" && \
		go build ./cmd/openapi && go build ./cmd/openapi-cli
	@cd "$(GIT_ROOT)/multipart" && \
		go build ./cmd/resume && go build ./cmd/resume-cli
	@cd "$(GIT_ROOT)/security/hierarchy" && \
		go build ./cmd/hierarchy && go build ./cmd/hierarchy-cli
	@cd "$(GIT_ROOT)/security/multiauth" && \
		go build ./cmd/multi_auth && go build ./cmd/multi_auth-cli
	@cd "$(GIT_ROOT)/streaming" && \
		go build ./cmd/chatter && go build ./cmd/chatter-cli
	@cd "$(GIT_ROOT)/tracing" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GIT_ROOT)/tus" && \
		go build ./cmd/upload && go build ./cmd/upload-cli
	@cd "$(GIT_ROOT)/upload_download" && \
		go build ./cmd/upload_download && go build ./cmd/upload_download-cli

test:
	@echo TESTING...
	@go test ./... > /dev/null
