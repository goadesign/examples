// Code generated by goa v3.13.1, DO NOT EDIT.
//
// hierarchy HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	apikeyservicec "goa.design/examples/security/hierarchy/gen/http/api_key_service/client"
	defaultservicec "goa.design/examples/security/hierarchy/gen/http/default_service/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `default-service default
api-key-service (default|secure)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` default-service default --username "Sit doloremque inventore eius." --password "Culpa alias."` + "\n" +
		os.Args[0] + ` api-key-service default --key "Aut delectus ipsam."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		defaultServiceFlags = flag.NewFlagSet("default-service", flag.ContinueOnError)

		defaultServiceDefaultFlags        = flag.NewFlagSet("default", flag.ExitOnError)
		defaultServiceDefaultUsernameFlag = defaultServiceDefaultFlags.String("username", "REQUIRED", "")
		defaultServiceDefaultPasswordFlag = defaultServiceDefaultFlags.String("password", "REQUIRED", "")

		apiKeyServiceFlags = flag.NewFlagSet("api-key-service", flag.ContinueOnError)

		apiKeyServiceDefaultFlags   = flag.NewFlagSet("default", flag.ExitOnError)
		apiKeyServiceDefaultKeyFlag = apiKeyServiceDefaultFlags.String("key", "REQUIRED", "")

		apiKeyServiceSecureFlags     = flag.NewFlagSet("secure", flag.ExitOnError)
		apiKeyServiceSecureTokenFlag = apiKeyServiceSecureFlags.String("token", "REQUIRED", "")
	)
	defaultServiceFlags.Usage = defaultServiceUsage
	defaultServiceDefaultFlags.Usage = defaultServiceDefaultUsage

	apiKeyServiceFlags.Usage = apiKeyServiceUsage
	apiKeyServiceDefaultFlags.Usage = apiKeyServiceDefaultUsage
	apiKeyServiceSecureFlags.Usage = apiKeyServiceSecureUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "default-service":
			svcf = defaultServiceFlags
		case "api-key-service":
			svcf = apiKeyServiceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "default-service":
			switch epn {
			case "default":
				epf = defaultServiceDefaultFlags

			}

		case "api-key-service":
			switch epn {
			case "default":
				epf = apiKeyServiceDefaultFlags

			case "secure":
				epf = apiKeyServiceSecureFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "default-service":
			c := defaultservicec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "default":
				endpoint = c.Default()
				data, err = defaultservicec.BuildDefaultPayload(*defaultServiceDefaultUsernameFlag, *defaultServiceDefaultPasswordFlag)
			}
		case "api-key-service":
			c := apikeyservicec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "default":
				endpoint = c.Default()
				data, err = apikeyservicec.BuildDefaultPayload(*apiKeyServiceDefaultKeyFlag)
			case "secure":
				endpoint = c.Secure()
				data, err = apikeyservicec.BuildSecurePayload(*apiKeyServiceSecureTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// default-serviceUsage displays the usage of the default-service command and
// its subcommands.
func defaultServiceUsage() {
	fmt.Fprintf(os.Stderr, `Service is the default_service service interface.
Usage:
    %[1]s [globalflags] default-service COMMAND [flags]

COMMAND:
    default: The default service default_method is secured with basic authentication

Additional help:
    %[1]s default-service COMMAND --help
`, os.Args[0])
}
func defaultServiceDefaultUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] default-service default -username STRING -password STRING

The default service default_method is secured with basic authentication
    -username STRING: 
    -password STRING: 

Example:
    %[1]s default-service default --username "Sit doloremque inventore eius." --password "Culpa alias."
`, os.Args[0])
}

// api-key-serviceUsage displays the usage of the api-key-service command and
// its subcommands.
func apiKeyServiceUsage() {
	fmt.Fprintf(os.Stderr, `The svc service is secured with API key based authentication
Usage:
    %[1]s [globalflags] api-key-service COMMAND [flags]

COMMAND:
    default: Default implements default.
    secure: This method requires a valid JWT token.

Additional help:
    %[1]s api-key-service COMMAND --help
`, os.Args[0])
}
func apiKeyServiceDefaultUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api-key-service default -key STRING

Default implements default.
    -key STRING: 

Example:
    %[1]s api-key-service default --key "Aut delectus ipsam."
`, os.Args[0])
}

func apiKeyServiceSecureUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api-key-service secure -token STRING

This method requires a valid JWT token.
    -token STRING: 

Example:
    %[1]s api-key-service secure --token "Quasi ex."
`, os.Args[0])
}
