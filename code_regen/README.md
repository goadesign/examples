Goa Code Regeneration
=====================

This example shows how to regenerate your apps code
without disrupting the modifications you have made.

__All you have to do is drop the `Makefile` and `restore.py` into your app dirrectory__

and type `make regen`

The process that happens is:

1. Backup the current go files
1. Remove and regenerate the app
1. Restore the modifications from the backups

This process is enabled by having [start/end] tags.
[PR in progress](https://github.com/goadesign/goa/pull/593) for adding the tags automatically.



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

```

The Makefile also wraps the goagen workflow.

- `make gen`
- `make regen`
- `make backup`
- `make restore`
- `make clean`
- `make clean_backups`
- `make clean_all`



