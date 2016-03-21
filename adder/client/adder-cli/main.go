package main

import (
	"fmt"
	"github.com/goadesign/examples/adder/client"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "adder-cli",
		Short: "CLI client for the adder service",
	}
	c := client.New(nil)
	c.UserAgent = "adder-cli/1.0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:8080", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout, defaults to 20s")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "add",
		Short: "add returns the sum of the left and right parameters in the response body",
	}
	tmp1 := new(AddOperandsCommand)
	sub = &cobra.Command{
		Use:   "operands",
		Short: "",
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub)
	command.AddCommand(sub)
	app.AddCommand(command)

}
