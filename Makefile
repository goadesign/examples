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

export GO111MODULE=on

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	honnef.co/go/tools/cmd/staticcheck \
	github.com/golang/protobuf/protoc-gen-go

.phony: all depend lint test build clean

all: check-goa gen lint test
	@echo DONE!

travis: depend all check-freshness

# Install protoc
PROTOC_VERSION=3.12.3
ifeq ($(GOOS),linux)
PROTOC=protoc-$(PROTOC_VERSION)-linux-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
else
	ifeq ($(GOOS),darwin)
PROTOC=protoc-$(PROTOC_VERSION)-osx-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
	else
		ifeq ($(GOOS),windows)
PROTOC=protoc-$(PROTOC_VERSION)-win32
PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
GOPATH:=$(subst \,/,$(GOPATH))
		endif
	endif
endif

check-goa:
ifdef GOA
	@echo $(GOA)
else
	go get -u goa.design/goa/v3
	go get -u goa.design/goa/v3/...
endif

depend:
	@echo INSTALLING DEPENDENCIES...
	@go mod download
	@env go get -v $(DEPEND)
	@echo installing protoc
	go-getter https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC).zip $(PROTOC)
	@cp $(PROTOC_EXEC) $(GOPATH)/bin && \
		rm -r $(PROTOC) && \
		echo "`protoc --version`"
	@echo done installing dependencies

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
	@rm -rf "$(GOPATH)/src/goa.design/examples/basic/cmd"             && \
	rm -rf "$(GOPATH)/src/goa.design/examples/cellar/cmd/cellar-cli"  && \
	rm -rf "$(GOPATH)/src/goa.design/examples/encodings/text/cmd"     && \
	rm -rf "$(GOPATH)/src/goa.design/examples/error/cmd"              && \
	rm -rf "$(GOPATH)/src/goa.design/examples/multipart/cmd"          && \
	rm -rf "$(GOPATH)/src/goa.design/examples/security/cmd"           && \
	echo "1/8 basic"                                                                                  && \
	goa gen     goa.design/examples/basic/design     -o "$(GOPATH)/src/goa.design/examples/basic"     && \
	goa example goa.design/examples/basic/design     -o "$(GOPATH)/src/goa.design/examples/basic"     && \
	echo "2/8 cellar"                                                                                 && \
	goa gen     goa.design/examples/cellar/design    -o "$(GOPATH)/src/goa.design/examples/cellar"    && \
	goa example goa.design/examples/cellar/design    -o "$(GOPATH)/src/goa.design/examples/cellar"    && \
	echo "3/8 encodings/cbor"                                                                         && \
	goa gen     goa.design/examples/encodings/cbor/design -o "$(GOPATH)/src/goa.design/examples/encodings/cbor" && \
	goa example goa.design/examples/encodings/cbor/design -o "$(GOPATH)/src/goa.design/examples/encodings/cbor" && \
	echo "4/8 encodings/text"                                                                         && \
	goa gen     goa.design/examples/encodings/text/design -o "$(GOPATH)/src/goa.design/examples/encodings/text" && \
	goa example goa.design/examples/encodings/text/design -o "$(GOPATH)/src/goa.design/examples/encodings/text" && \
	echo "5/8 error"                                                                                  && \
	goa gen     goa.design/examples/error/design     -o "$(GOPATH)/src/goa.design/examples/error"     && \
	goa example goa.design/examples/error/design     -o "$(GOPATH)/src/goa.design/examples/error"     && \
	echo "6/8 multipart"                                                                              && \
	goa gen     goa.design/examples/multipart/design -o "$(GOPATH)/src/goa.design/examples/multipart" && \
	goa example goa.design/examples/multipart/design -o "$(GOPATH)/src/goa.design/examples/multipart" && \
	echo "7/8 security"                                                                               && \
	goa gen     goa.design/examples/security/design  -o "$(GOPATH)/src/goa.design/examples/security"  && \
	goa example goa.design/examples/security/design  -o "$(GOPATH)/src/goa.design/examples/security"  && \
	echo "8/8 streaming"                                                                              && \
	goa gen     goa.design/examples/streaming/design -o "$(GOPATH)/src/goa.design/examples/streaming" && \
	goa example goa.design/examples/streaming/design -o "$(GOPATH)/src/goa.design/examples/streaming"

build:
	@cd "$(GOPATH)/src/goa.design/examples/basic" && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd "$(GOPATH)/src/goa.design/examples/cellar" && \
		go build ./cmd/cellar && go build ./cmd/cellar-cli
	@cd "$(GOPATH)/src/goa.design/examples/encodings/cbor" && \
		go build ./cmd/cbor && go build ./cmd/cbor-cli
	@cd "$(GOPATH)/src/goa.design/examples/encodings/text" && \
		go build ./cmd/text && go build ./cmd/text-cli
	@cd "$(GOPATH)/src/goa.design/examples/error" && \
		go build ./cmd/divider && go build ./cmd/divider-cli
	@cd "$(GOPATH)/src/goa.design/examples/multipart" && \
		go build ./cmd/resume && go build ./cmd/resume-cli
	@cd "$(GOPATH)/src/goa.design/examples/security" && \
		go build ./cmd/multi_auth && go build ./cmd/multi_auth-cli
	@cd "$(GOPATH)/src/goa.design/examples/streaming" && \
		go build ./cmd/chatter && go build ./cmd/chatter-cli
	@cd "$(GOPATH)/src/goa.design/examples/tracing" && \
		go build ./cmd/calc && go build ./cmd/calc-cli

clean:
	@cd "$(GOPATH)/src/goa.design/examples/basic" && \
		rm -f calc calc-cli
	@cd "$(GOPATH)/src/goa.design/examples/cellar" && \
		 rm -f cellar cellar-cli
	@cd "$(GOPATH)/src/goa.design/examples/encodings/cbor" && \
		 rm -f cbor cbor-cli
	@cd "$(GOPATH)/src/goa.design/examples/encodings/text" && \
		 rm -f text text-cli
	@cd "$(GOPATH)/src/goa.design/examples/error" && \
		 rm -f divider divider-cli
	@cd "$(GOPATH)/src/goa.design/examples/multipart" && \
		 rm -f resume resume-cli
	@cd "$(GOPATH)/src/goa.design/examples/security" && \
		 rm -f multi_auth multi_auth-cli
	@cd "$(GOPATH)/src/goa.design/examples/streaming" && \
		 rm -f chatter chatter-cli
	@cd "$(GOPATH)/src/goa.design/examples/tracing" && \
		 rm -f calc calc-cli

test:
	@echo TESTING...
	@go test ./... > /dev/null

check-freshness:
	@if [ "`git diff -- . ':(exclude)go.mod' ':(exclude)go.sum' | wc -l`" -gt "0" ]; then \
	        echo "[ERROR] generated code not in-sync with design:"; \
	        echo; \
	        git status -s; \
	        git --no-pager diff; \
	        echo; \
	        exit 1; \
	fi
