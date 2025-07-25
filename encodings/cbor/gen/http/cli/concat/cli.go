// Code generated by goa v3.21.5, DO NOT EDIT.
//
// concat HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	concatc "goa.design/examples/encodings/cbor/gen/http/concat/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `concat concat
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` concat concat --a "In eos." --b "Dolorum reiciendis rerum quae."` + "\n" +
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
		concatFlags = flag.NewFlagSet("concat", flag.ContinueOnError)

		concatConcatFlags = flag.NewFlagSet("concat", flag.ExitOnError)
		concatConcatAFlag = concatConcatFlags.String("a", "REQUIRED", "Left operand")
		concatConcatBFlag = concatConcatFlags.String("b", "REQUIRED", "Right operand")
	)
	concatFlags.Usage = concatUsage
	concatConcatFlags.Usage = concatConcatUsage

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
		case "concat":
			svcf = concatFlags
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
		case "concat":
			switch epn {
			case "concat":
				epf = concatConcatFlags

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
		case "concat":
			c := concatc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "concat":
				endpoint = c.Concat()
				data, err = concatc.BuildConcatPayload(*concatConcatAFlag, *concatConcatBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// concatUsage displays the usage of the concat command and its subcommands.
func concatUsage() {
	fmt.Fprintf(os.Stderr, `The concat service performs operations on strings.
	
	The service uses the CBOR binary serialization standard to encode responses.
	It supports reading requests encoded with CBOR, JSON, XML or GOB.
	
Usage:
    %[1]s [globalflags] concat COMMAND [flags]

COMMAND:
    concat: Concat implements concat.

Additional help:
    %[1]s concat COMMAND --help
`, os.Args[0])
}
func concatConcatUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] concat concat -a STRING -b STRING

Concat implements concat.
    -a STRING: Left operand
    -b STRING: Right operand

Example:
    %[1]s concat concat --a "In eos." --b "Dolorum reiciendis rerum quae."
`, os.Args[0])
}
