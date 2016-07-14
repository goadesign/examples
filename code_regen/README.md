goa Code Regeneration
=====================


This example shows how to use a helper script
and Makefile to regenerate your code
without disrupting the modifications you have made.

Using the methods below,
you can modify your design,
regenerate your application,
and find your code existing within the freshly generated goa code.
This helper pack is by no means necessary
for goa app regeneration.
Under normal circumstances goa will not touch or update
the files in the root application directory and
considers these files completely under your control.
However, if you delete one of the controller files,
then goa will generate it because it does not exist.
This helper pack uses that behavior
to backup, regenerate, and restore your files.


__Note__: This API development process encouraged by this helper pack
does not necessarily follow best practices.
It is best used during early iteration,
as a learning tool,
or to a gain deeper understanding into
goa and the translations between design and code.


#### Installation

All you have to do is drop the `Makefile` and `restore.py` into your app dirrectory


#### Usage

Type `make regen`

This is what happens:

1. Backup the current go files in the root directory.
1. `rm *.go`
1. `goagen` the app like new.
1. Restore your code from the backups

This process is enabled by having [start/end] tags.
These are generated automatically for the controller stubs.
You may also use two special tags `import` and `extra`.
`import` will appear as the first block in the import clause
and enables the maintenance of additional imports.
`extra` will be appended to the end of the file
and allows extra functions to be added to the file.


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

The Makefile has the following commands:

- `make`     : regen
- `make regen` : backup clean gen restore
- `make backup` : cp *.go -> *.go.backup
- `make clean`: rm *.go
- `make gen` : goagen
- `make restore` : restore.py
- `make clean_backups` : rm *.backup
- `make clean_all` : clean clean_backups



