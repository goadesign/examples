goa Code Regeneration
=====================

Note: This depends on a
[PR in progress](https://github.com/goadesign/goa/pull/593) for changing TDB to the tag pairs.

---

This example shows how to regenerate your apps code
without disrupting the modifications you have made.
This allows you to modify your design, after the initial generation,
and keep the logic you have already implemented.

__All you have to do is drop the `Makefile` and `restore.py` into your app dirrectory__

and type `make regen`

The process that happens is:

1. Backup the current go files
1. Remove and regenerate the app
1. Restore the modifications from the backups

This process is enabled by having [start/end] tags.
These are genreated automatically for the controller stubs.
You may also use two special tags `import` and `extra`.
`import` will appear as the first block in the import clause
and enables the maintenance of additional imports.
`extra` will be appended to the end of the file
and allows extra functions to be added to the file.


```Go
package main

import (
  // import start_implement
  "fmt"
  // import end_implement
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

The Makefile also wraps the goagen workflow.

- `make gen` : goagen
- `make`     : go build .
- `make run` : ./<app>
- `make build` : go build .
- `make install` : go install .
- `make regen` : backup, gen, restore
- `make backup` : cp *.go -> *.go.backup
- `make restore` : restore.py
- `make clean`: rm *.go
- `make clean_backups` : rm *.backup
- `make clean_all` : clean, clean_backups



