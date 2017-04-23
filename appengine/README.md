# App Engine Example

This example shows how to setup a goa service to run in [Google App Engine](https://cloud.google.com/appengine/).

## Usage

### Execute example code:   
Please updated part of Makefile if forking

```
##### Convenient command ######

REPO:=github.com/"your-repo"/examples/appengine <- change here
```

```bash
$ make for_example
$ make local
```

### When updating the DSL:

If you want to update the following

- app
- client
- tool
- swagger

```bash
$ make gen
```

---

If you want to update the following

- controller
- main.go
- app
- client
- tool
- swagger

```bash
$ make main
```

### When deploying:
Please updated part of Makefile 

```bash
##### Convenient command ######

REPO:=~~~
GAE_PROJECT:=projectName <- change here
```

## About rewriting the contents of vendor

App Engine uses Go 1.6 and thus code being deployed to it cannot take advantage of the `context` package which was introduced in Go 1.7. The Makefile in this directory leverages [gorep](https://github.com/novalagung/gorep) to replace the import statements in the generated code and in the vendored goa with `golang.org/x/net/context`.