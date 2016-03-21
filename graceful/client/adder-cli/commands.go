package main

import (
	"fmt"
	"github.com/goadesign/examples/graceful/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// AddOperandsCommand is the command line data structure for the add action of operands
	AddOperandsCommand struct {
	}
)

// Run makes the HTTP request corresponding to the AddOperandsCommand command.
func (cmd *AddOperandsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.UseLogger(context.Background(), logger)
	resp, err := c.AddOperands(ctx, path)
	if err != nil {
		goa.Error(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddOperandsCommand) RegisterFlags(cc *cobra.Command) {
}
