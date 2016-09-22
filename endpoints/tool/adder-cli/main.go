package main

import (
	"fmt"
	"github.com/goadesign/examples/endpoints/client"
	"github.com/goadesign/examples/endpoints/tool/cli"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "adder-cli",
		Short: `CLI client for the adder service`,
	}

	// Create client struct
	httpClient := newHTTPClient()
	c := client.New(goaclient.HTTPClientDoer(httpClient))

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "goa-endpoints.appspot.com", "API hostname")
	app.PersistentFlags().DurationVarP(&httpClient.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Register signer flags
	var key, format string
	app.PersistentFlags().StringVar(&key, "key", "", "API key used for authentication")
	app.PersistentFlags().StringVar(&format, "format", "Bearer %s", "Format used to create auth header or query from key")
	var token, typ string
	app.PersistentFlags().StringVar(&token, "token", "", "Token used for authentication")
	app.PersistentFlags().StringVar(&typ, "token-type", "Bearer", "Token type used for authentication")

	// Parse flags and setup signers
	app.ParseFlags(os.Args)
	source := &goaclient.StaticTokenSource{
		StaticToken: &goaclient.StaticToken{Type: typ, Value: token},
	}
	apiKeySigner := newAPIKeySigner(key, format)
	jwtSigner := newJWTSigner(source)

	// Initialize API client
	c.SetAPIKeySigner(apiKeySigner)
	c.SetJWTSigner(jwtSigner)
	c.UserAgent = "adder-cli/0"

	// Register API commands
	cli.RegisterCommands(app, c)

	// Execute!
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}

// newAPIKeySigner returns the request signer used for authenticating
// against the api_key security scheme.
func newAPIKeySigner(key, format string) goaclient.Signer {
	return &goaclient.APIKeySigner{
		SignQuery: true,
		KeyName:   "key",
		KeyValue:  key,
		Format:    format,
	}

}

// newJWTSigner returns the request signer used for authenticating
// against the jwt security scheme.
func newJWTSigner(source goaclient.TokenSource) goaclient.Signer {
	return &goaclient.OAuth2Signer{
		TokenSource: source,
	}

}
