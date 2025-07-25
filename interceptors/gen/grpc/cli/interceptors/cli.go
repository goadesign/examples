// Code generated by goa v3.21.5, DO NOT EDIT.
//
// interceptors gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package cli

import (
	"flag"
	"fmt"
	"os"

	interceptorsc "goa.design/examples/interceptors/gen/grpc/interceptors/client"
	interceptors "goa.design/examples/interceptors/gen/interceptors"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `interceptors get
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` interceptors get --message '{
      "auth": "Exercitationem vitae ipsum molestiae.",
      "recordID": "b39b771d-1348-4566-bde1-e935e7a128c7",
      "spanID": "b39b771d-1348-4566-bde1-e935e7a128c7",
      "tenantID": "b39b771d-1348-4566-bde1-e935e7a128c7",
      "traceID": "b39b771d-1348-4566-bde1-e935e7a128c7"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	cc *grpc.ClientConn,
	inter interceptors.ClientInterceptors,
	opts ...grpc.CallOption,
) (goa.Endpoint, any, error) {
	var (
		interceptorsFlags = flag.NewFlagSet("interceptors", flag.ContinueOnError)

		interceptorsGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		interceptorsGetMessageFlag = interceptorsGetFlags.String("message", "", "")
	)
	interceptorsFlags.Usage = interceptorsUsage
	interceptorsGetFlags.Usage = interceptorsGetUsage

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
		case "interceptors":
			svcf = interceptorsFlags
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
		case "interceptors":
			switch epn {
			case "get":
				epf = interceptorsGetFlags

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
		case "interceptors":
			c := interceptorsc.NewClient(cc, opts...)
			switch epn {
			case "get":
				endpoint = c.Get()
				endpoint = interceptors.WrapGetClientEndpoint(endpoint, inter)
				data, err = interceptorsc.BuildGetPayload(*interceptorsGetMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// interceptorsUsage displays the usage of the interceptors command and its
// subcommands.
func interceptorsUsage() {
	fmt.Fprintf(os.Stderr, `The interceptors service demonstrates a comprehensive use of interceptors combining
	authentication, tenant validation, caching, audit logging, and retry mechanisms. It showcases
	both client-side and server-side interceptors working together to provide a robust service.
Usage:
    %[1]s [globalflags] interceptors COMMAND [flags]

COMMAND:
    get: Get retrieves a record by ID with all interceptors in action

Additional help:
    %[1]s interceptors COMMAND --help
`, os.Args[0])
}
func interceptorsGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] interceptors get -message JSON

Get retrieves a record by ID with all interceptors in action
    -message JSON: 

Example:
    %[1]s interceptors get --message '{
      "auth": "Exercitationem vitae ipsum molestiae.",
      "recordID": "b39b771d-1348-4566-bde1-e935e7a128c7",
      "spanID": "b39b771d-1348-4566-bde1-e935e7a128c7",
      "tenantID": "b39b771d-1348-4566-bde1-e935e7a128c7",
      "traceID": "b39b771d-1348-4566-bde1-e935e7a128c7"
   }'
`, os.Args[0])
}
