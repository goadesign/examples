# App Engine Example

This example shows how to setup a goa service to run in
[Google App Engine](https://cloud.google.com/appengine/) *standard* environment.

App Engine standard environment uses Go 1.6 and thus code being deployed to it
cannot take advantage of the `context` package which was introduced in Go 1.7.

The Makefile in this directory leverages
[gorep](https://github.com/novalagung/gorep) to replace the import statements in
the generated code and in the vendored goa with `golang.org/x/net/context`. The
Makefile also contains other convenience targets for maintaining a goa project,
see below.

## Running the Hello World Example

### Prerequesites

Compiling and running applications on App Engine requires the use of the App
Engine Go SDK, follow the
[instructions](https://cloud.google.com/appengine/docs/standard/go/download) on
the App Engine website to install it.

Deploying the example in this directory requires creating a Google App Engine
project, refer to the App Engine
[documentation](https://cloud.google.com/appengine/docs/standard/go/) for
information on how to do that.

### Building

If you forked this repository you must first update the Makefile so that the
`REPO` variable points to the right place:

```
##### Convenience targets ######

REPO:=github.com/goadesign/examples/appengine <- change to match your path
```

Once the path is correct run:

```bash
$ make example
```

This target installs `glide` if necessary then runs it, runs `goagen bootstrap`
and finally runs `gorep` on both the vendored goa and the generated files.

### Running Locally

Assuming the App Engine Go SDK is properly installed and that in particular the
`goapp` tool is properly installed and in the PATH, run:

```bash
$ make local
```

This runs `goapp serve`. At this point the hello world service should be running
locally.

### Deploying

To deploy the example to App Engine first make sure the `projectName` variable
in the Makefile is properly set:

```
GAE_PROJECT:=projectName <- change
```

then run

```bash
make deploy
```

this uses `goapp` to deploy the example to the App Engine project. At this
point the service should be running in App Engine!

### Updating the DSL

The Makefile also includes a `gen` target which calls `goagen` and is useful
when working on a project after having modified the design. The target generates
the following directories:

- app
- client
- tool
- swagger

```bash
$ make gen
```

The Makefile also includes a `main` target which generates the scaffold:

- controller
- main.go

```bash
$ make main
```

