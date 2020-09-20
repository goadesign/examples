// Code generated by goa v3.2.4, DO NOT EDIT.
//
// storage gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package client

import (
	"encoding/json"
	"fmt"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	goa "goa.design/goa/v3/pkg"
)

// BuildShowPayload builds the payload for the storage show endpoint from CLI
// flags.
func BuildShowPayload(storageShowMessage string, storageShowView string) (*storage.ShowPayload, error) {
	var err error
	var message storagepb.ShowRequest
	{
		if storageShowMessage != "" {
			err = json.Unmarshal([]byte(storageShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Earum dolorem.\"\n   }'")
			}
		}
	}
	var view *string
	{
		if storageShowView != "" {
			view = &storageShowView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &storage.ShowPayload{
		ID: message.Id,
	}
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the storage add endpoint from CLI
// flags.
func BuildAddPayload(storageAddMessage string) (*storage.Bottle, error) {
	var err error
	var message storagepb.AddRequest
	{
		if storageAddMessage != "" {
			err = json.Unmarshal([]byte(storageAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"composition\": [\n         {\n            \"percentage\": 73,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 73,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 73,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 73,\n            \"varietal\": \"Syrah\"\n         }\n      ],\n      \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n      \"name\": \"Blue\\'s Cuvee\",\n      \"rating\": 4,\n      \"vintage\": 1978,\n      \"winery\": {\n         \"country\": \"USA\",\n         \"name\": \"Longoria\",\n         \"region\": \"Central Coast, California\",\n         \"url\": \"http://www.longoriawine.com/\"\n      }\n   }'")
			}
		}
	}
	v := &storage.Bottle{
		Name:    message.Name,
		Vintage: message.Vintage,
	}
	if message.Description != "" {
		v.Description = &message.Description
	}
	if message.Rating != 0 {
		v.Rating = &message.Rating
	}
	if message.Winery != nil {
		v.Winery = protobufStoragepbWineryToStorageWinery(message.Winery)
	}
	if message.Composition != nil {
		v.Composition = make([]*storage.Component, len(message.Composition))
		for i, val := range message.Composition {
			v.Composition[i] = &storage.Component{
				Varietal: val.Varietal,
			}
			if val.Percentage != 0 {
				v.Composition[i].Percentage = &val.Percentage
			}
		}
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the storage remove endpoint from
// CLI flags.
func BuildRemovePayload(storageRemoveMessage string) (*storage.RemovePayload, error) {
	var err error
	var message storagepb.RemoveRequest
	{
		if storageRemoveMessage != "" {
			err = json.Unmarshal([]byte(storageRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Corporis quam delectus quas exercitationem alias est.\"\n   }'")
			}
		}
	}
	v := &storage.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildRatePayload builds the payload for the storage rate endpoint from CLI
// flags.
func BuildRatePayload(storageRateMessage string) (map[uint32][]string, error) {
	var err error
	var message storagepb.RateRequest
	{
		if storageRateMessage != "" {
			err = json.Unmarshal([]byte(storageRateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"field\": {\n         \"1210888915\": {\n            \"field\": [\n               \"Expedita in quam eos distinctio.\",\n               \"Ut molestiae possimus.\",\n               \"Aliquam itaque quam beatae veniam quaerat sint.\",\n               \"Error sit qui ut delectus nihil sunt.\"\n            ]\n         },\n         \"1558969343\": {\n            \"field\": [\n               \"Expedita in quam eos distinctio.\",\n               \"Ut molestiae possimus.\",\n               \"Aliquam itaque quam beatae veniam quaerat sint.\",\n               \"Error sit qui ut delectus nihil sunt.\"\n            ]\n         },\n         \"2429176308\": {\n            \"field\": [\n               \"Expedita in quam eos distinctio.\",\n               \"Ut molestiae possimus.\",\n               \"Aliquam itaque quam beatae veniam quaerat sint.\",\n               \"Error sit qui ut delectus nihil sunt.\"\n            ]\n         }\n      }\n   }'")
			}
		}
	}
	v := make(map[uint32][]string, len(message.Field))
	for key, val := range message.Field {
		tk := key
		tv := make([]string, len(val.Field))
		for i, val := range val.Field {
			tv[i] = val
		}
		v[tk] = tv
	}
	return v, nil
}

// BuildMultiAddPayload builds the payload for the storage multi_add endpoint
// from CLI flags.
func BuildMultiAddPayload(storageMultiAddMessage string) ([]*storage.Bottle, error) {
	var err error
	var message storagepb.MultiAddRequest
	{
		if storageMultiAddMessage != "" {
			err = json.Unmarshal([]byte(storageMultiAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"field\": [\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 2005,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 2005,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 2005,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 2005,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         }\n      ]\n   }'")
			}
		}
	}
	v := make([]*storage.Bottle, len(message.Field))
	for i, val := range message.Field {
		v[i] = &storage.Bottle{
			Name:    val.Name,
			Vintage: val.Vintage,
		}
		if val.Description != "" {
			v[i].Description = &val.Description
		}
		if val.Rating != 0 {
			v[i].Rating = &val.Rating
		}
		if val.Winery != nil {
			v[i].Winery = protobufStoragepbWineryToStorageWinery(val.Winery)
		}
		if val.Composition != nil {
			v[i].Composition = make([]*storage.Component, len(val.Composition))
			for j, val := range val.Composition {
				v[i].Composition[j] = &storage.Component{
					Varietal: val.Varietal,
				}
				if val.Percentage != 0 {
					v[i].Composition[j].Percentage = &val.Percentage
				}
			}
		}
	}
	return v, nil
}

// BuildMultiUpdatePayload builds the payload for the storage multi_update
// endpoint from CLI flags.
func BuildMultiUpdatePayload(storageMultiUpdateMessage string) (*storage.MultiUpdatePayload, error) {
	var err error
	var message storagepb.MultiUpdateRequest
	{
		if storageMultiUpdateMessage != "" {
			err = json.Unmarshal([]byte(storageMultiUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"bottles\": [\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 2005,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 73,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 2005,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         }\n      ],\n      \"ids\": [\n         \"Voluptas numquam et aperiam.\",\n         \"Qui aliquid sit et.\",\n         \"In est.\"\n      ]\n   }'")
			}
		}
	}
	v := &storage.MultiUpdatePayload{}
	if message.Ids != nil {
		v.Ids = make([]string, len(message.Ids))
		for i, val := range message.Ids {
			v.Ids[i] = val
		}
	}
	if message.Bottles != nil {
		v.Bottles = make([]*storage.Bottle, len(message.Bottles))
		for i, val := range message.Bottles {
			v.Bottles[i] = &storage.Bottle{
				Name:    val.Name,
				Vintage: val.Vintage,
			}
			if val.Description != "" {
				v.Bottles[i].Description = &val.Description
			}
			if val.Rating != 0 {
				v.Bottles[i].Rating = &val.Rating
			}
			if val.Winery != nil {
				v.Bottles[i].Winery = protobufStoragepbWineryToStorageWinery(val.Winery)
			}
			if val.Composition != nil {
				v.Bottles[i].Composition = make([]*storage.Component, len(val.Composition))
				for j, val := range val.Composition {
					v.Bottles[i].Composition[j] = &storage.Component{
						Varietal: val.Varietal,
					}
					if val.Percentage != 0 {
						v.Bottles[i].Composition[j].Percentage = &val.Percentage
					}
				}
			}
		}
	}

	return v, nil
}
