// Code generated by goa v3.1.1, DO NOT EDIT.
//
// storage HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	storage "goa.design/examples/cellar/gen/storage"
	goa "goa.design/goa/v3/pkg"
)

// BuildShowPayload builds the payload for the storage show endpoint from CLI
// flags.
func BuildShowPayload(storageShowID string, storageShowView string) (*storage.ShowPayload, error) {
	var err error
	var id string
	{
		id = storageShowID
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
	v := &storage.ShowPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the storage add endpoint from CLI
// flags.
func BuildAddPayload(storageAddBody string) (*storage.Bottle, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(storageAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"composition\": [\n         {\n            \"percentage\": 92,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 92,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 92,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 92,\n            \"varietal\": \"Syrah\"\n         }\n      ],\n      \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n      \"name\": \"Blue\\'s Cuvee\",\n      \"rating\": 1,\n      \"vintage\": 1941,\n      \"winery\": {\n         \"country\": \"USA\",\n         \"name\": \"Longoria\",\n         \"region\": \"Central Coast, California\",\n         \"url\": \"http://www.longoriawine.com/\"\n      }\n   }'")
		}
		if body.Winery == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
		}
		if utf8.RuneCountInString(body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
		}
		if body.Winery != nil {
			if err2 := ValidateWineryRequestBody(body.Winery); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", body.Vintage, 1900, true))
		}
		if body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", body.Vintage, 2020, false))
		}
		for _, e := range body.Composition {
			if e != nil {
				if err2 := ValidateComponentRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if body.Description != nil {
			if utf8.RuneCountInString(*body.Description) > 2000 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
			}
		}
		if body.Rating != nil {
			if *body.Rating < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
			}
		}
		if body.Rating != nil {
			if *body.Rating > 5 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &storage.Bottle{
		Name:        body.Name,
		Vintage:     body.Vintage,
		Description: body.Description,
		Rating:      body.Rating,
	}
	if body.Winery != nil {
		v.Winery = marshalWineryRequestBodyToStorageWinery(body.Winery)
	}
	if body.Composition != nil {
		v.Composition = make([]*storage.Component, len(body.Composition))
		for i, val := range body.Composition {
			v.Composition[i] = marshalComponentRequestBodyToStorageComponent(val)
		}
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the storage remove endpoint from
// CLI flags.
func BuildRemovePayload(storageRemoveID string) (*storage.RemovePayload, error) {
	var id string
	{
		id = storageRemoveID
	}
	v := &storage.RemovePayload{}
	v.ID = id

	return v, nil
}

// BuildMultiAddPayload builds the payload for the storage multi_add endpoint
// from CLI flags.
func BuildMultiAddPayload(storageMultiAddBody string) ([]*storage.Bottle, error) {
	var err error
	var body []*BottleRequestBody
	{
		err = json.Unmarshal([]byte(storageMultiAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'[\n      {\n         \"composition\": [\n            {\n               \"percentage\": 92,\n               \"varietal\": \"Syrah\"\n            },\n            {\n               \"percentage\": 92,\n               \"varietal\": \"Syrah\"\n            }\n         ],\n         \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"rating\": 1,\n         \"vintage\": 1984,\n         \"winery\": {\n            \"country\": \"USA\",\n            \"name\": \"Longoria\",\n            \"region\": \"Central Coast, California\",\n            \"url\": \"http://www.longoriawine.com/\"\n         }\n      },\n      {\n         \"composition\": [\n            {\n               \"percentage\": 92,\n               \"varietal\": \"Syrah\"\n            },\n            {\n               \"percentage\": 92,\n               \"varietal\": \"Syrah\"\n            }\n         ],\n         \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"rating\": 1,\n         \"vintage\": 1984,\n         \"winery\": {\n            \"country\": \"USA\",\n            \"name\": \"Longoria\",\n            \"region\": \"Central Coast, California\",\n            \"url\": \"http://www.longoriawine.com/\"\n         }\n      }\n   ]'")
		}
	}
	v := make([]*storage.Bottle, len(body))
	for i, val := range body {
		v[i] = marshalBottleRequestBodyToStorageBottle(val)
	}
	return v, nil
}

// BuildMultiUpdatePayload builds the payload for the storage multi_update
// endpoint from CLI flags.
func BuildMultiUpdatePayload(storageMultiUpdateBody string, storageMultiUpdateIds string) (*storage.MultiUpdatePayload, error) {
	var err error
	var body MultiUpdateRequestBody
	{
		err = json.Unmarshal([]byte(storageMultiUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"bottles\": [\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 1,\n            \"vintage\": 1984,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 1,\n            \"vintage\": 1984,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 1,\n            \"vintage\": 1984,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 92,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 1,\n            \"vintage\": 1984,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         }\n      ]\n   }'")
		}
		if body.Bottles == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("bottles", "body"))
		}
		for _, e := range body.Bottles {
			if e != nil {
				if err2 := ValidateBottleRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var ids []string
	{
		err = json.Unmarshal([]byte(storageMultiUpdateIds), &ids)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for ids, example of valid JSON:\n%s", "'[\n      \"Aut rem vel veritatis.\",\n      \"Animi nulla aut aut.\"\n   ]'")
		}
	}
	v := &storage.MultiUpdatePayload{}
	if body.Bottles != nil {
		v.Bottles = make([]*storage.Bottle, len(body.Bottles))
		for i, val := range body.Bottles {
			v.Bottles[i] = marshalBottleRequestBodyToStorageBottle(val)
		}
	}
	v.Ids = ids

	return v, nil
}
