goa Controller Code Regeneration
================================

This example contains a Makefile and associated script which makes it possible to easily
apply the latest scaffold to an existing controller file.

### Context

goa generates two types of code:

* The goagen owned code which cannot be edited and is completely regenerated each time. Non
  generated code use the goagen owned code by importing the generated package.
  The goagen owned code includes: the `app` package, the `client` package and the `tool/cli`
  package.

* The scaffold code that is generated once as a convenience. This code should be treated
  identically to non generated code. goagen won't override such files if they already exist.
  The scaffold code includes: the `main` package files: `main.go`,  all the
  `<controller name>.go` files as well as the `tool/<name of api>-cli/main.go` file.

### Strategy

Sometimes it may be convenient to retrieve the scaffold that would be generated if a given
controller file did not already exist - for example after having updated the design to
include new actions.

The makefile included in this example moves the existing controller files so that goagen
re-generates the scaffold. It then applies the existing changes back to the newly generated
files.

The code for the existing controller methods is automatically copied over. Any code that
appears between two `extra` tags is also copied over, see the example below.

### Usage

1. Copy the files `Makefile` and `restore.py` from this example into the service root directory (the directory containing the generated controller files).
2. `cd` into the service root directory.
3. Run `make`

### Example

Given the file below containing the user written `Checkin` controller implementation:

```Go
package main

import (
  "github.com/goadesign/goa"
  "nginx-checkin/app"
)

...

// Create runs the create action.
func (c *CheckinController) Create(ctx *app.CreateCheckinContext) error {

        // CheckinController_Create: start_implement

        fmt.Println("Creating Server")

        // CheckinController_Create: end_implement

        res := &app.GoaVerdvermFrisbyServer{}
        return ctx.OK(res)
}

...

// extra start_implement
func hello() {
  fmt.Println("hello")
}
// extra end_implement
```

Running `make` causes the file to get replaced with a newly generated scaffold 
where both the content of the controller methods and the code in between the
`extra start_implement` and `extra end_implement` tags is copied over.
