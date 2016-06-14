#! /usr/bin/make
#
# Makefile for goa examples
#
# Generates and builds all examples.
#
DIRS=$(wildcard */.)
DEPEND=\
	github.com/goadesign/goa        \
	github.com/goadesign/goa/goagen \
	github.com/tylerb/graceful      \
	gopkg.in/yaml.v2                \
	github.com/ajg/form

.PHONY : all
all: depend build

depend:
	@go get -v $(DEPEND)

build:
	@for d in $(DIRS); do \
		echo $$d && cd $$d && go generate > /dev/null && go build && cd ..; \
	done
