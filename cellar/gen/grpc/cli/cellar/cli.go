// Code generated by goa v3.19.0, DO NOT EDIT.
//
// cellar gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package cli

import (
	"flag"
	"fmt"
	"os"

	sommelierc "goa.design/examples/cellar/gen/grpc/sommelier/client"
	storagec "goa.design/examples/cellar/gen/grpc/storage/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `sommelier pick
storage (list|show|add|remove|rate|multi-add|multi-update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sommelier pick --message '{
      "name": "Blue\'s Cuvee",
      "varietal": [
         "pinot noir",
         "merlot",
         "cabernet franc"
      ],
      "winery": "longoria"
   }'` + "\n" +
		os.Args[0] + ` storage list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		sommelierFlags = flag.NewFlagSet("sommelier", flag.ContinueOnError)

		sommelierPickFlags       = flag.NewFlagSet("pick", flag.ExitOnError)
		sommelierPickMessageFlag = sommelierPickFlags.String("message", "", "")

		storageFlags = flag.NewFlagSet("storage", flag.ContinueOnError)

		storageListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		storageShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		storageShowMessageFlag = storageShowFlags.String("message", "", "")
		storageShowViewFlag    = storageShowFlags.String("view", "", "")

		storageAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		storageAddMessageFlag = storageAddFlags.String("message", "", "")

		storageRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		storageRemoveMessageFlag = storageRemoveFlags.String("message", "", "")

		storageRateFlags       = flag.NewFlagSet("rate", flag.ExitOnError)
		storageRateMessageFlag = storageRateFlags.String("message", "", "")

		storageMultiAddFlags       = flag.NewFlagSet("multi-add", flag.ExitOnError)
		storageMultiAddMessageFlag = storageMultiAddFlags.String("message", "", "")

		storageMultiUpdateFlags       = flag.NewFlagSet("multi-update", flag.ExitOnError)
		storageMultiUpdateMessageFlag = storageMultiUpdateFlags.String("message", "", "")
	)
	sommelierFlags.Usage = sommelierUsage
	sommelierPickFlags.Usage = sommelierPickUsage

	storageFlags.Usage = storageUsage
	storageListFlags.Usage = storageListUsage
	storageShowFlags.Usage = storageShowUsage
	storageAddFlags.Usage = storageAddUsage
	storageRemoveFlags.Usage = storageRemoveUsage
	storageRateFlags.Usage = storageRateUsage
	storageMultiAddFlags.Usage = storageMultiAddUsage
	storageMultiUpdateFlags.Usage = storageMultiUpdateUsage

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
		case "sommelier":
			svcf = sommelierFlags
		case "storage":
			svcf = storageFlags
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
		case "sommelier":
			switch epn {
			case "pick":
				epf = sommelierPickFlags

			}

		case "storage":
			switch epn {
			case "list":
				epf = storageListFlags

			case "show":
				epf = storageShowFlags

			case "add":
				epf = storageAddFlags

			case "remove":
				epf = storageRemoveFlags

			case "rate":
				epf = storageRateFlags

			case "multi-add":
				epf = storageMultiAddFlags

			case "multi-update":
				epf = storageMultiUpdateFlags

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
		case "sommelier":
			c := sommelierc.NewClient(cc, opts...)
			switch epn {
			case "pick":
				endpoint = c.Pick()
				data, err = sommelierc.BuildPickPayload(*sommelierPickMessageFlag)
			}
		case "storage":
			c := storagec.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
			case "show":
				endpoint = c.Show()
				data, err = storagec.BuildShowPayload(*storageShowMessageFlag, *storageShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = storagec.BuildAddPayload(*storageAddMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = storagec.BuildRemovePayload(*storageRemoveMessageFlag)
			case "rate":
				endpoint = c.Rate()
				data, err = storagec.BuildRatePayload(*storageRateMessageFlag)
			case "multi-add":
				endpoint = c.MultiAdd()
				data, err = storagec.BuildMultiAddPayload(*storageMultiAddMessageFlag)
			case "multi-update":
				endpoint = c.MultiUpdate()
				data, err = storagec.BuildMultiUpdatePayload(*storageMultiUpdateMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
} // sommelierUsage displays the usage of the sommelier command and its
// subcommands.
func sommelierUsage() {
	fmt.Fprintf(os.Stderr, `The sommelier service retrieves bottles given a set of criteria.
Usage:
    %[1]s [globalflags] sommelier COMMAND [flags]

COMMAND:
    pick: Pick implements pick.

Additional help:
    %[1]s sommelier COMMAND --help
`, os.Args[0])
}
func sommelierPickUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sommelier pick -message JSON

Pick implements pick.
    -message JSON: 

Example:
    %[1]s sommelier pick --message '{
      "name": "Blue\'s Cuvee",
      "varietal": [
         "pinot noir",
         "merlot",
         "cabernet franc"
      ],
      "winery": "longoria"
   }'
`, os.Args[0])
}

// storageUsage displays the usage of the storage command and its subcommands.
func storageUsage() {
	fmt.Fprintf(os.Stderr, `The storage service makes it possible to view, add or remove wine bottles.
Usage:
    %[1]s [globalflags] storage COMMAND [flags]

COMMAND:
    list: List all stored bottles
    show: Show bottle by ID
    add: Add new bottle and return its ID.
    remove: Remove bottle from storage
    rate: Rate bottles by IDs
    multi-add: Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    multi-update: Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.

Additional help:
    %[1]s storage COMMAND --help
`, os.Args[0])
}
func storageListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage list

List all stored bottles

Example:
    %[1]s storage list
`, os.Args[0])
}

func storageShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show -message JSON -view STRING

Show bottle by ID
    -message JSON: 
    -view STRING: 

Example:
    %[1]s storage show --message '{
      "id": "Earum dolorem."
   }' --view "default"
`, os.Args[0])
}

func storageAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage add -message JSON

Add new bottle and return its ID.
    -message JSON: 

Example:
    %[1]s storage add --message '{
      "composition": [
         {
            "percentage": 73,
            "varietal": "Syrah"
         },
         {
            "percentage": 73,
            "varietal": "Syrah"
         },
         {
            "percentage": 73,
            "varietal": "Syrah"
         },
         {
            "percentage": 73,
            "varietal": "Syrah"
         }
      ],
      "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
      "name": "Blue\'s Cuvee",
      "rating": 4,
      "vintage": 1978,
      "winery": {
         "country": "USA",
         "name": "Longoria",
         "region": "Central Coast, California",
         "url": "http://www.longoriawine.com/"
      }
   }'
`, os.Args[0])
}

func storageRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage remove -message JSON

Remove bottle from storage
    -message JSON: 

Example:
    %[1]s storage remove --message '{
      "id": "Corporis quam delectus quas exercitationem alias est."
   }'
`, os.Args[0])
}

func storageRateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage rate -message JSON

Rate bottles by IDs
    -message JSON: 

Example:
    %[1]s storage rate --message '{
      "field": {
         "1210888915": {
            "field": [
               "Expedita in quam eos distinctio.",
               "Ut molestiae possimus.",
               "Aliquam itaque quam beatae veniam quaerat sint.",
               "Error sit qui ut delectus nihil sunt."
            ]
         },
         "1558969343": {
            "field": [
               "Expedita in quam eos distinctio.",
               "Ut molestiae possimus.",
               "Aliquam itaque quam beatae veniam quaerat sint.",
               "Error sit qui ut delectus nihil sunt."
            ]
         },
         "2429176308": {
            "field": [
               "Expedita in quam eos distinctio.",
               "Ut molestiae possimus.",
               "Aliquam itaque quam beatae veniam quaerat sint.",
               "Error sit qui ut delectus nihil sunt."
            ]
         }
      }
   }'
`, os.Args[0])
}

func storageMultiAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage multi-add -message JSON

Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    -message JSON: 

Example:
    %[1]s storage multi-add --message '{
      "field": [
         {
            "composition": [
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 2005,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 2005,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 2005,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 2005,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         }
      ]
   }'
`, os.Args[0])
}

func storageMultiUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage multi-update -message JSON

Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.
    -message JSON: 

Example:
    %[1]s storage multi-update --message '{
      "bottles": [
         {
            "composition": [
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 2005,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 73,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 2005,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         }
      ],
      "ids": [
         "Voluptas numquam et aperiam.",
         "Qui aliquid sit et.",
         "In est."
      ]
   }'
`, os.Args[0])
}
