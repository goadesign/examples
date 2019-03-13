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
DIRS=$(shell go list -f {{.Dir}} goa.design/examples/...)

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

.phony: all depend lint test build clean

all: depend lint gen test
	@echo DONE!

travis: all check-freshness

# Install protoc
GOOS=$(shell go env GOOS)
PROTOC_VERSION="3.6.1"
ifeq ($(GOOS),linux)
PROTOC="protoc-$(PROTOC_VERSION)-linux-x86_64"
PROTOC_EXEC="$(PROTOC)/bin/protoc"
GOBIN="$(GOPATH)/bin"
IMPORTS_PATH=/*.go
else
	ifeq ($(GOOS),darwin)
PROTOC="protoc-$(PROTOC_VERSION)-osx-x86_64"
PROTOC_EXEC="$(PROTOC)/bin/protoc"
GOBIN="$(GOPATH)/bin"
IMPORTS_PATH=/*.go
	else
		ifeq ($(GOOS),windows)
PROTOC="protoc-$(PROTOC_VERSION)-win32"
PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
GOBIN="$(GOPATH)\bin"
IMPORTS_PATH=\*.go
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
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d$(IMPORTS_PATH) | grep -v '.pb.go' | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi

gen:
	@# NOTE: not all command line tools are generated
	@echo GENERATING CODE...
	@rm -rf $(GOPATH)/src/goa.design/examples/basic/cmd             && \
	rm -rf $(GOPATH)/src/goa.design/examples/cellar/cmd/cellar-cli  && \
	rm -rf $(GOPATH)/src/goa.design/examples/encodings/cmd          && \
	rm -rf $(GOPATH)/src/goa.design/examples/error/cmd              && \
	rm -rf $(GOPATH)/src/goa.design/examples/multipart/cmd          && \
	rm -rf $(GOPATH)/src/goa.design/examples/security/cmd           && \
	goa gen     goa.design/examples/basic/design     -o $(GOPATH)/src/goa.design/examples/basic     && \
	goa example goa.design/examples/basic/design     -o $(GOPATH)/src/goa.design/examples/basic     && \
	goa gen     goa.design/examples/cellar/design    -o $(GOPATH)/src/goa.design/examples/cellar    && \
	goa example goa.design/examples/cellar/design    -o $(GOPATH)/src/goa.design/examples/cellar    && \
	goa gen     goa.design/examples/encodings/design -o $(GOPATH)/src/goa.design/examples/encodings && \
	goa example goa.design/examples/encodings/design -o $(GOPATH)/src/goa.design/examples/encodings && \
	goa gen     goa.design/examples/error/design     -o $(GOPATH)/src/goa.design/examples/error     && \
	goa example goa.design/examples/error/design     -o $(GOPATH)/src/goa.design/examples/error     && \
	goa gen     goa.design/examples/multipart/design -o $(GOPATH)/src/goa.design/examples/multipart && \
	goa example goa.design/examples/multipart/design -o $(GOPATH)/src/goa.design/examples/multipart && \
	goa gen     goa.design/examples/security/design  -o $(GOPATH)/src/goa.design/examples/security  && \
	goa example goa.design/examples/security/design  -o $(GOPATH)/src/goa.design/examples/security  && \
	goa gen     goa.design/examples/streaming/design -o $(GOPATH)/src/goa.design/examples/streaming && \
	goa example goa.design/examples/streaming/design -o $(GOPATH)/src/goa.design/examples/streaming

build:
	@cd $(GOPATH)/src/goa.design/examples/basic && \
		go build ./cmd/calc && go build ./cmd/calc-cli
	@cd $(GOPATH)/src/goa.design/examples/cellar && \
		go build ./cmd/cellar && go build ./cmd/cellar-cli
	@cd $(GOPATH)/src/goa.design/examples/encodings && \
		go build ./cmd/encodings && go build ./cmd/encodings-cli
	@cd $(GOPATH)/src/goa.design/examples/error && \
		go build ./cmd/divider && go build ./cmd/divider-cli
	@cd $(GOPATH)/src/goa.design/examples/multipart && \
		go build ./cmd/resume && go build ./cmd/resume-cli
	@cd $(GOPATH)/src/goa.design/examples/security && \
		go build ./cmd/multi_auth && go build ./cmd/multi_auth-cli
	@cd $(GOPATH)/src/goa.design/examples/streaming && \
		go build ./cmd/chatter && go build ./cmd/chatter-cli
	@cd $(GOPATH)/src/goa.design/examples/tracing && \
		go build ./cmd/calc && go build ./cmd/calc-cli

clean:
	@cd $(GOPATH)/src/goa.design/examples/basic && \
		rm -f calc calc-cli
	@cd $(GOPATH)/src/goa.design/examples/cellar && \
		 rm -f cellar cellar-cli
	@cd $(GOPATH)/src/goa.design/examples/encodings && \
		 rm -f encodings encodings-cli
	@cd $(GOPATH)/src/goa.design/examples/error && \
		 rm -f divider divider-cli
	@cd $(GOPATH)/src/goa.design/examples/multipart && \
		 rm -f resume resume-cli
	@cd $(GOPATH)/src/goa.design/examples/security && \
		 rm -f multi_auth multi_auth-cli
	@cd $(GOPATH)/src/goa.design/examples/streaming && \
		 rm -f chatter chatter-cli
	@cd $(GOPATH)/src/goa.design/examples/tracing && \
		 rm -f calc calc-cli

test:
	@echo TESTING...
	@go test ./... > /dev/null

check-freshness:
	@if [ "`git status -s | wc -l`" -gt "0" ]; then \
	        echo "[ERROR] generated code not in-sync with design:"; \
	        echo; \
	        git status -s; \
	        git --no-pager diff; \
	        echo; \
	        exit 1; \
	fi
