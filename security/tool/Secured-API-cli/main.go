package main

import (
	"fmt"
	"github.com/goadesign/examples/security/client"
	"github.com/goadesign/examples/security/tool/cli"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "Secured API-cli",
		Short: `CLI client for the Secured API service`,
	}

	// Create client struct
	c := client.New(newHTTPClient())

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Register signer flags
	var user, pass string
	app.PersistentFlags().StringVar(&user, "user", "", "Username used for authentication")
	app.PersistentFlags().StringVar(&pass, "pass", "", "Password used for authentication")
	var key, format string
	app.PersistentFlags().StringVar(&key, "key", "", "API key used for authentication")
	app.PersistentFlags().StringVar(&format, "format", "Bearer %s", "Format used to create auth header or query from key")
	var token, typ string
	app.PersistentFlags().StringVar(&token, "token", "", "Token used for authentication")
	app.PersistentFlags().StringVar(&typ, "token-type", "Bearer", "Token type used for authentication")
	source := &goaclient.StaticTokenSource{
		StaticToken: &goaclient.StaticToken{Type: typ, Value: token},
	}

	// Parse flags and setup signers
	app.ParseFlags(os.Args)
	apiKeySigner := newAPIKeySigner(key, format)
	basicAuthSigner := newBasicAuthSigner(user, pass)
	jwtSigner := newJWTSigner(key, format)
	oauth2Signer := newOauth2Signer(source)

	// Initialize API client
	c.SetAPIKeySigner(apiKeySigner)
	c.SetBasicAuthSigner(basicAuthSigner)
	c.SetJWTSigner(jwtSigner)
	c.SetOauth2Signer(oauth2Signer)
	c.UserAgent = "Secured API-cli/0"

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
		SignQuery: false,
		KeyName:   "X-Shared-Secret",
		KeyValue:  key,
		Format:    format,
	}

}

// newBasicAuthSigner returns the request signer used for authenticating
// against the basic_auth security scheme.
func newBasicAuthSigner(user, pass string) goaclient.Signer {
	return &goaclient.BasicSigner{
		Username: user,
		Password: pass,
	}

}

// newJWTSigner returns the request signer used for authenticating
// against the jwt security scheme.
func newJWTSigner(key, format string) goaclient.Signer {
	return &goaclient.APIKeySigner{
		SignQuery: false,
		KeyName:   "Authorization",
		KeyValue:  key,
		Format:    format,
	}

}

// newOauth2Signer returns the request signer used for authenticating
// against the oauth2 security scheme.
func newOauth2Signer(source goaclient.TokenSource) goaclient.Signer {
	return &goaclient.OAuth2Signer{
		TokenSource: source,
	}

}
