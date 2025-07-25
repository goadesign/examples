// Code generated by goa v3.21.5, DO NOT EDIT.
//
// upload_download HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/upload_download/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	updownc "goa.design/examples/upload_download/gen/http/updown/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `updown (upload|download)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` updown upload --dir "upload" --content-type "multipart/form-data; boundary=goa" --stream "goa.png"` + "\n" +
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
		updownFlags = flag.NewFlagSet("updown", flag.ContinueOnError)

		updownUploadFlags           = flag.NewFlagSet("upload", flag.ExitOnError)
		updownUploadDirFlag         = updownUploadFlags.String("dir", "REQUIRED", "Dir is the relative path to the file directory where the uploaded content is saved.")
		updownUploadContentTypeFlag = updownUploadFlags.String("content-type", "multipart/form-data; boundary=goa", "")
		updownUploadStreamFlag      = updownUploadFlags.String("stream", "REQUIRED", "path to file containing the streamed request body")

		updownDownloadFlags = flag.NewFlagSet("download", flag.ExitOnError)
		updownDownloadPFlag = updownDownloadFlags.String("p", "REQUIRED", "Path to downloaded file.")
	)
	updownFlags.Usage = updownUsage
	updownUploadFlags.Usage = updownUploadUsage
	updownDownloadFlags.Usage = updownDownloadUsage

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
		case "updown":
			svcf = updownFlags
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
		case "updown":
			switch epn {
			case "upload":
				epf = updownUploadFlags

			case "download":
				epf = updownDownloadFlags

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
		case "updown":
			c := updownc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "upload":
				endpoint = c.Upload()
				data, err = updownc.BuildUploadPayload(*updownUploadDirFlag, *updownUploadContentTypeFlag)
				if err == nil {
					data, err = updownc.BuildUploadStreamPayload(data, *updownUploadStreamFlag)
				}
			case "download":
				endpoint = c.Download()
				data = *updownDownloadPFlag
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// updownUsage displays the usage of the updown command and its subcommands.
func updownUsage() {
	fmt.Fprintf(os.Stderr, `Service updown demonstrates how to implement upload and download of files in
	Goa without having to load the entire content in memory first.
	The upload method uses SkipRequestBodyEncodeDecode to delegate reading the HTTP
	request body to the service logic. This alleviates the need for loading the
	full body content in memory first to decode it into a data structure. Note that
	using SkipRequestBodyDecode is incompatible with gRPC and can only be used on
	methods that only define a HTTP transport mapping. This example implementation
	leverages package "mime/multipart" to read the request body.
	Similarly the download method uses SkipResponseBodyEncodeDecode to stream the 
	file to the client without requiring to load the complete content in memory
	first. As with SkipRequestBodyDecode using SkipResponseBodyEncodeDecode is
	incompatible with gRPC.
Usage:
    %[1]s [globalflags] updown COMMAND [flags]

COMMAND:
    upload: Upload implements upload.
    download: Download implements download.

Additional help:
    %[1]s updown COMMAND --help
`, os.Args[0])
}
func updownUploadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] updown upload -dir STRING -content-type STRING -stream STRING

Upload implements upload.
    -dir STRING: Dir is the relative path to the file directory where the uploaded content is saved.
    -content-type STRING: 
    -stream STRING: path to file containing the streamed request body

Example:
    %[1]s updown upload --dir "upload" --content-type "multipart/form-data; boundary=goa" --stream "goa.png"
`, os.Args[0])
}

func updownDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] updown download -p STRING

Download implements download.
    -p STRING: Path to downloaded file.

Example:
    %[1]s updown download --p "Hic accusamus delectus voluptatum architecto."
`, os.Args[0])
}
