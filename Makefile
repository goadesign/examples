#! /usr/bin/make
#
# Makefile for goa examples
#
# Generates and builds all examples.
#
DIRS=$(wildcard */.)
DIRS+=$(wildcard xray/*/.)
ROOT=$(shell pwd)
NOBUILD="gopherjs/." "code_regen/." "xray/." "appengine/." "files/."
DEPEND=\
	github.com/ajg/form             \
	github.com/dgrijalva/jwt-go     \
	github.com/goadesign/goa/...    \
	github.com/goadesign/oauth2     \
	github.com/tylerb/graceful      \
	gopkg.in/yaml.v2

.PHONY : all
all: depend build

depend:
	@go get -v $(DEPEND)

build:
	@for d in $(DIRS); do \
		skip=$$(echo 0); \
		for e in $(NOBUILD); do \
			if test $$e = $$d; then \
				skip=$$(echo 1); \
				break; \
			fi; \
		done; \
		if test "$$skip" = "1"; then continue; fi; \
		echo $$d && cd $$d && go generate > /dev/null && go build && cd $(ROOT); \
	done
