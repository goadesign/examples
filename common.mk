# Common variables
GO_FILES=$(shell find . -type f -name '*.go')
MODULE=$(shell head -n1 go.mod | cut -d ' ' -f2)

# Dependencies
DEPEND=\
	github.com/hashicorp/go-getter \
	github.com/cheggaaa/pb \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/golang/protobuf/proto \
	goa.design/goa/... \
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	honnef.co/go/tools/cmd/staticcheck

# Protoc setup
GOOS=$(shell go env GOOS)
PROTOC_VERSION=28.3
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

# Common targets
.PHONY: all depend lint test build clean gen check-freshness

all: gen lint test
	@echo DONE!

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
	@if [ "`staticcheck ./... | grep -v ".pb.go" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi

gen:
	@echo GENERATING CODE...
	@goa gen "$(MODULE)/design"
ifneq ($(SKIP_GOA_EXAMPLE),true)
	@goa example "$(MODULE)/design"
endif

build:
	@go build "./cmd/$(APP)" && go build "./cmd/$(APP)-cli"

clean:
	@rm -f "./$(APP)" "./$(APP)-cli"

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
