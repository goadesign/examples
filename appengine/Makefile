#! /usr/bin/make
#
# Makefile for goa examples in appengine
#
# Targets:
# - appengine
# This command makes necessary changes to use with GAE / Go.
# Also, please be sure to vendoring If you use this command.
#
appengine:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./vendor/github.com/goadesign/goa" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./app" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./client" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./tool" \
          -from="context" \
          -to="golang.org/x/net/context"
