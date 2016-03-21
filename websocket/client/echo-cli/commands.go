package main

import (
	"github.com/goadesign/examples/websocket/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// ConnectEchoCommand is the command line data structure for the connect action of echo
	ConnectEchoCommand struct {
		// Initial message to echo
		Initial string
	}
)

// Run establishes a websocket connection for the ConnectEchoCommand command.
func (cmd *ConnectEchoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/echo"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.UseLogger(context.Background(), logger)
	ws, err := c.ConnectEcho(ctx, path, cmd.Initial)
	if err != nil {
		goa.Error(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ConnectEchoCommand) RegisterFlags(cc *cobra.Command) {
	var tmp2 string
	cc.Flags().StringVar(&cmd.Initial, "initial", tmp2, "Initial message to echo")
}
