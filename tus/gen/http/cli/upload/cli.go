// Code generated by goa v3.21.5, DO NOT EDIT.
//
// Upload HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/tus/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	tusc "goa.design/examples/tus/gen/http/tus/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `tus (head|patch|options|post|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tus head --id "99homqugg35laha266s6" --tus-resumable "1.0.0"` + "\n" +
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
		tusFlags = flag.NewFlagSet("tus", flag.ContinueOnError)

		tusHeadFlags            = flag.NewFlagSet("head", flag.ExitOnError)
		tusHeadIDFlag           = tusHeadFlags.String("id", "REQUIRED", "IDs are generated using Xid: https://github.com/rs/xid")
		tusHeadTusResumableFlag = tusHeadFlags.String("tus-resumable", "REQUIRED", "")

		tusPatchFlags              = flag.NewFlagSet("patch", flag.ExitOnError)
		tusPatchIDFlag             = tusPatchFlags.String("id", "REQUIRED", "IDs are generated using Xid: https://github.com/rs/xid")
		tusPatchTusResumableFlag   = tusPatchFlags.String("tus-resumable", "REQUIRED", "")
		tusPatchUploadOffsetFlag   = tusPatchFlags.String("upload-offset", "REQUIRED", "")
		tusPatchUploadChecksumFlag = tusPatchFlags.String("upload-checksum", "", "")
		tusPatchStreamFlag         = tusPatchFlags.String("stream", "REQUIRED", "path to file containing the streamed request body")

		tusOptionsFlags = flag.NewFlagSet("options", flag.ExitOnError)

		tusPostFlags                 = flag.NewFlagSet("post", flag.ExitOnError)
		tusPostTusResumableFlag      = tusPostFlags.String("tus-resumable", "REQUIRED", "")
		tusPostUploadLengthFlag      = tusPostFlags.String("upload-length", "", "")
		tusPostUploadDeferLengthFlag = tusPostFlags.String("upload-defer-length", "", "")
		tusPostUploadChecksumFlag    = tusPostFlags.String("upload-checksum", "", "")
		tusPostUploadMetadataFlag    = tusPostFlags.String("upload-metadata", "", "")
		tusPostTusMaxSizeFlag        = tusPostFlags.String("tus-max-size", "", "")
		tusPostStreamFlag            = tusPostFlags.String("stream", "REQUIRED", "path to file containing the streamed request body")

		tusDeleteFlags            = flag.NewFlagSet("delete", flag.ExitOnError)
		tusDeleteIDFlag           = tusDeleteFlags.String("id", "REQUIRED", "IDs are generated using Xid: https://github.com/rs/xid")
		tusDeleteTusResumableFlag = tusDeleteFlags.String("tus-resumable", "REQUIRED", "")
	)
	tusFlags.Usage = tusUsage
	tusHeadFlags.Usage = tusHeadUsage
	tusPatchFlags.Usage = tusPatchUsage
	tusOptionsFlags.Usage = tusOptionsUsage
	tusPostFlags.Usage = tusPostUsage
	tusDeleteFlags.Usage = tusDeleteUsage

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
		case "tus":
			svcf = tusFlags
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
		case "tus":
			switch epn {
			case "head":
				epf = tusHeadFlags

			case "patch":
				epf = tusPatchFlags

			case "options":
				epf = tusOptionsFlags

			case "post":
				epf = tusPostFlags

			case "delete":
				epf = tusDeleteFlags

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
		case "tus":
			c := tusc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "head":
				endpoint = c.Head()
				data, err = tusc.BuildHeadPayload(*tusHeadIDFlag, *tusHeadTusResumableFlag)
			case "patch":
				endpoint = c.Patch()
				data, err = tusc.BuildPatchPayload(*tusPatchIDFlag, *tusPatchTusResumableFlag, *tusPatchUploadOffsetFlag, *tusPatchUploadChecksumFlag)
				if err == nil {
					data, err = tusc.BuildPatchStreamPayload(data, *tusPatchStreamFlag)
				}
			case "options":
				endpoint = c.Options()
			case "post":
				endpoint = c.Post()
				data, err = tusc.BuildPostPayload(*tusPostTusResumableFlag, *tusPostUploadLengthFlag, *tusPostUploadDeferLengthFlag, *tusPostUploadChecksumFlag, *tusPostUploadMetadataFlag, *tusPostTusMaxSizeFlag)
				if err == nil {
					data, err = tusc.BuildPostStreamPayload(data, *tusPostStreamFlag)
				}
			case "delete":
				endpoint = c.Delete()
				data, err = tusc.BuildDeletePayload(*tusDeleteIDFlag, *tusDeleteTusResumableFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// tusUsage displays the usage of the tus command and its subcommands.
func tusUsage() {
	fmt.Fprintf(os.Stderr, `The tus service exposes the methods required to implement the tus protocol
Usage:
    %[1]s [globalflags] tus COMMAND [flags]

COMMAND:
    head: Clients use the HEAD request to determine the offset at which the upload should be continued.
    patch: Clients use the PATCH method to start or resume an upload.
    options: Clients use the OPTIONS method to gather information about the Server’s current configuration.
    post: Clients use the POST method against a known upload creation URL to request a new upload resource.
    delete: Clients use the DELETE method to terminate completed and unfinished uploads allowing the Server to free up used resources.

Additional help:
    %[1]s tus COMMAND --help
`, os.Args[0])
}
func tusHeadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] tus head -id STRING -tus-resumable STRING

Clients use the HEAD request to determine the offset at which the upload should be continued.
    -id STRING: IDs are generated using Xid: https://github.com/rs/xid
    -tus-resumable STRING: 

Example:
    %[1]s tus head --id "99homqugg35laha266s6" --tus-resumable "1.0.0"
`, os.Args[0])
}

func tusPatchUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] tus patch -id STRING -tus-resumable STRING -upload-offset INT64 -upload-checksum STRING -stream STRING

Clients use the PATCH method to start or resume an upload.
    -id STRING: IDs are generated using Xid: https://github.com/rs/xid
    -tus-resumable STRING: 
    -upload-offset INT64: 
    -upload-checksum STRING: 
    -stream STRING: path to file containing the streamed request body

Example:
    %[1]s tus patch --id "hmpmrpp0ls3ru16n0vkq" --tus-resumable "1.0.0" --upload-offset 4913379198212964978 --upload-checksum "sha1 Kq5sNclPz7QV2+lfQIuc6R7oRu0=" --stream "goa.png"
`, os.Args[0])
}

func tusOptionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] tus options

Clients use the OPTIONS method to gather information about the Server’s current configuration.

Example:
    %[1]s tus options
`, os.Args[0])
}

func tusPostUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] tus post -tus-resumable STRING -upload-length INT64 -upload-defer-length INT -upload-checksum STRING -upload-metadata STRING -tus-max-size INT64 -stream STRING

Clients use the POST method against a known upload creation URL to request a new upload resource.
    -tus-resumable STRING: 
    -upload-length INT64: 
    -upload-defer-length INT: 
    -upload-checksum STRING: 
    -upload-metadata STRING: 
    -tus-max-size INT64: 
    -stream STRING: path to file containing the streamed request body

Example:
    %[1]s tus post --tus-resumable "1.0.0" --upload-length 7471840069206534621 --upload-defer-length 1 --upload-checksum "sha1 Kq5sNclPz7QV2+lfQIuc6R7oRu0=" --upload-metadata "key1 val1,key2 val2" --tus-max-size 4187081630712482631 --stream "goa.png"
`, os.Args[0])
}

func tusDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] tus delete -id STRING -tus-resumable STRING

Clients use the DELETE method to terminate completed and unfinished uploads allowing the Server to free up used resources.
    -id STRING: IDs are generated using Xid: https://github.com/rs/xid
    -tus-resumable STRING: 

Example:
    %[1]s tus delete --id "fd90p6dgjhach3rv88d9" --tus-resumable "1.0.0"
`, os.Args[0])
}
