# App Engine Example

This example shows how to setup a goa service to run in [Google App Engine](https://cloud.google.com/appengine/).

App Engine uses Go 1.6 and thus code being deployed to it cannot take advantage of the `context` package which was introduced in Go 1.7. The Makefile in this directory leverages [gorep](https://github.com/novalagung/gorep) to replace the import statements in the generated code and in the vendored goa with `golang.org/x/net/context`.

Usage:

1. Use `goagen` to generate the service code
2. Vendor `goa`
3. Run `make appengine`
4. Deploy to App Engine