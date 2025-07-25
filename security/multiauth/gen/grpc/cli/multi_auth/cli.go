// Code generated by goa v3.21.5, DO NOT EDIT.
//
// multi_auth gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design

package cli

import (
	"flag"
	"fmt"
	"os"

	securedservicec "goa.design/examples/security/multiauth/gen/grpc/secured_service/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `secured-service (signin|secure|doubly-secure|also-doubly-secure)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` secured-service signin --username "user" --password "password"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	cc *grpc.ClientConn,
	opts ...grpc.CallOption,
) (goa.Endpoint, any, error) {
	var (
		securedServiceFlags = flag.NewFlagSet("secured-service", flag.ContinueOnError)

		securedServiceSigninFlags        = flag.NewFlagSet("signin", flag.ExitOnError)
		securedServiceSigninUsernameFlag = securedServiceSigninFlags.String("username", "REQUIRED", "")
		securedServiceSigninPasswordFlag = securedServiceSigninFlags.String("password", "REQUIRED", "")

		securedServiceSecureFlags       = flag.NewFlagSet("secure", flag.ExitOnError)
		securedServiceSecureMessageFlag = securedServiceSecureFlags.String("message", "", "")
		securedServiceSecureTokenFlag   = securedServiceSecureFlags.String("token", "REQUIRED", "")

		securedServiceDoublySecureFlags       = flag.NewFlagSet("doubly-secure", flag.ExitOnError)
		securedServiceDoublySecureMessageFlag = securedServiceDoublySecureFlags.String("message", "", "")
		securedServiceDoublySecureTokenFlag   = securedServiceDoublySecureFlags.String("token", "REQUIRED", "")

		securedServiceAlsoDoublySecureFlags          = flag.NewFlagSet("also-doubly-secure", flag.ExitOnError)
		securedServiceAlsoDoublySecureMessageFlag    = securedServiceAlsoDoublySecureFlags.String("message", "", "")
		securedServiceAlsoDoublySecureOauthTokenFlag = securedServiceAlsoDoublySecureFlags.String("oauth-token", "", "")
		securedServiceAlsoDoublySecureTokenFlag      = securedServiceAlsoDoublySecureFlags.String("token", "", "")
	)
	securedServiceFlags.Usage = securedServiceUsage
	securedServiceSigninFlags.Usage = securedServiceSigninUsage
	securedServiceSecureFlags.Usage = securedServiceSecureUsage
	securedServiceDoublySecureFlags.Usage = securedServiceDoublySecureUsage
	securedServiceAlsoDoublySecureFlags.Usage = securedServiceAlsoDoublySecureUsage

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
		case "secured-service":
			svcf = securedServiceFlags
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
		case "secured-service":
			switch epn {
			case "signin":
				epf = securedServiceSigninFlags

			case "secure":
				epf = securedServiceSecureFlags

			case "doubly-secure":
				epf = securedServiceDoublySecureFlags

			case "also-doubly-secure":
				epf = securedServiceAlsoDoublySecureFlags

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
		case "secured-service":
			c := securedservicec.NewClient(cc, opts...)
			switch epn {
			case "signin":
				endpoint = c.Signin()
				data, err = securedservicec.BuildSigninPayload(*securedServiceSigninUsernameFlag, *securedServiceSigninPasswordFlag)
			case "secure":
				endpoint = c.Secure()
				data, err = securedservicec.BuildSecurePayload(*securedServiceSecureMessageFlag, *securedServiceSecureTokenFlag)
			case "doubly-secure":
				endpoint = c.DoublySecure()
				data, err = securedservicec.BuildDoublySecurePayload(*securedServiceDoublySecureMessageFlag, *securedServiceDoublySecureTokenFlag)
			case "also-doubly-secure":
				endpoint = c.AlsoDoublySecure()
				data, err = securedservicec.BuildAlsoDoublySecurePayload(*securedServiceAlsoDoublySecureMessageFlag, *securedServiceAlsoDoublySecureOauthTokenFlag, *securedServiceAlsoDoublySecureTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// securedServiceUsage displays the usage of the secured-service command and
// its subcommands.
func securedServiceUsage() {
	fmt.Fprintf(os.Stderr, `The secured service exposes endpoints that require valid authorization credentials.
Usage:
    %[1]s [globalflags] secured-service COMMAND [flags]

COMMAND:
    signin: Creates a valid JWT
    secure: This action is secured with the jwt scheme
    doubly-secure: This action is secured with the jwt scheme and also requires an API key query string.
    also-doubly-secure: This action is secured with the jwt scheme and also requires an API key header.

Additional help:
    %[1]s secured-service COMMAND --help
`, os.Args[0])
}
func securedServiceSigninUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] secured-service signin -username STRING -password STRING

Creates a valid JWT
    -username STRING: 
    -password STRING: 

Example:
    %[1]s secured-service signin --username "user" --password "password"
`, os.Args[0])
}

func securedServiceSecureUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] secured-service secure -message JSON -token STRING

This action is secured with the jwt scheme
    -message JSON: 
    -token STRING: 

Example:
    %[1]s secured-service secure --message '{
      "fail": true
   }' --token "Veniam quis nulla officiis id rerum."
`, os.Args[0])
}

func securedServiceDoublySecureUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] secured-service doubly-secure -message JSON -token STRING

This action is secured with the jwt scheme and also requires an API key query string.
    -message JSON: 
    -token STRING: 

Example:
    %[1]s secured-service doubly-secure --message '{
      "key": "abcdef12345"
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func securedServiceAlsoDoublySecureUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] secured-service also-doubly-secure -message JSON -oauth-token STRING -token STRING

This action is secured with the jwt scheme and also requires an API key header.
    -message JSON: 
    -oauth-token STRING: 
    -token STRING: 

Example:
    %[1]s secured-service also-doubly-secure --message '{
      "key": "abcdef12345",
      "password": "password",
      "username": "user"
   }' --oauth-token "Magni perferendis unde itaque." --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}
