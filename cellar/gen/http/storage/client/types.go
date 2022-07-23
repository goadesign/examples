// Code generated by goa v3.7.13, DO NOT EDIT.
//
// storage HTTP client types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"unicode/utf8"

	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "storage" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery" json:"winery" xml:"winery"`
	// Vintage of bottle
	Vintage uint32 `form:"vintage" json:"vintage" xml:"vintage"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// MultiUpdateRequestBody is the type of the "storage" service "multi_update"
// endpoint HTTP request body.
type MultiUpdateRequestBody struct {
	// Array of bottle info that matches the ids attribute
	Bottles []*BottleRequestBody `form:"bottles" json:"bottles" xml:"bottles"`
}

// ListResponseBody is the type of the "storage" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredBottleResponse

// ShowResponseBody is the type of the "storage" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryResponseBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponseBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "storage" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing bottle
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredBottleResponse is used to define fields on response body types.
type StoredBottleResponse struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryResponse `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponse `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// WineryResponse is used to define fields on response body types.
type WineryResponse struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentResponse is used to define fields on response body types.
type ComponentResponse struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// WineryResponseBody is used to define fields on response body types.
type WineryResponseBody struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentResponseBody is used to define fields on response body types.
type ComponentResponseBody struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// WineryRequestBody is used to define fields on request body types.
type WineryRequestBody struct {
	// Name of winery
	Name string `form:"name" json:"name" xml:"name"`
	// Region of winery
	Region string `form:"region" json:"region" xml:"region"`
	// Country of winery
	Country string `form:"country" json:"country" xml:"country"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentRequestBody is used to define fields on request body types.
type ComponentRequestBody struct {
	// Grape varietal
	Varietal string `form:"varietal" json:"varietal" xml:"varietal"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// BottleRequestBody is used to define fields on request body types.
type BottleRequestBody struct {
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery" json:"winery" xml:"winery"`
	// Vintage of bottle
	Vintage uint32 `form:"vintage" json:"vintage" xml:"vintage"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "storage" service.
func NewAddRequestBody(p *storage.Bottle) *AddRequestBody {
	body := &AddRequestBody{
		Name:        p.Name,
		Vintage:     p.Vintage,
		Description: p.Description,
		Rating:      p.Rating,
	}
	if p.Winery != nil {
		body.Winery = marshalStorageWineryToWineryRequestBody(p.Winery)
	}
	if p.Composition != nil {
		body.Composition = make([]*ComponentRequestBody, len(p.Composition))
		for i, val := range p.Composition {
			body.Composition[i] = marshalStorageComponentToComponentRequestBody(val)
		}
	}
	return body
}

// NewBottleRequestBody builds the HTTP request body from the payload of the
// "multi_add" endpoint of the "storage" service.
func NewBottleRequestBody(p []*storage.Bottle) []*BottleRequestBody {
	body := make([]*BottleRequestBody, len(p))
	for i, val := range p {
		body[i] = marshalStorageBottleToBottleRequestBody(val)
	}
	return body
}

// NewMultiUpdateRequestBody builds the HTTP request body from the payload of
// the "multi_update" endpoint of the "storage" service.
func NewMultiUpdateRequestBody(p *storage.MultiUpdatePayload) *MultiUpdateRequestBody {
	body := &MultiUpdateRequestBody{}
	if p.Bottles != nil {
		body.Bottles = make([]*BottleRequestBody, len(p.Bottles))
		for i, val := range p.Bottles {
			body.Bottles[i] = marshalStorageBottleToBottleRequestBody(val)
		}
	}
	return body
}

// NewListStoredBottleCollectionOK builds a "storage" service "list" endpoint
// result from a HTTP "OK" response.
func NewListStoredBottleCollectionOK(body ListResponseBody) storageviews.StoredBottleCollectionView {
	v := make([]*storageviews.StoredBottleView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredBottleResponseToStorageviewsStoredBottleView(val)
	}

	return v
}

// NewShowStoredBottleOK builds a "storage" service "show" endpoint result from
// a HTTP "OK" response.
func NewShowStoredBottleOK(body *ShowResponseBody) *storageviews.StoredBottleView {
	v := &storageviews.StoredBottleView{
		ID:          body.ID,
		Name:        body.Name,
		Vintage:     body.Vintage,
		Description: body.Description,
		Rating:      body.Rating,
	}
	v.Winery = unmarshalWineryResponseBodyToStorageviewsWineryView(body.Winery)
	if body.Composition != nil {
		v.Composition = make([]*storageviews.ComponentView, len(body.Composition))
		for i, val := range body.Composition {
			v.Composition[i] = unmarshalComponentResponseBodyToStorageviewsComponentView(val)
		}
	}

	return v
}

// NewShowNotFound builds a storage service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *storage.NotFound {
	v := &storage.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateStoredBottleResponse runs the validations defined on
// StoredBottleResponse
func ValidateStoredBottleResponse(body *StoredBottleResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := ValidateWineryResponse(body.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := ValidateComponentResponse(e); err2 != nil {
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
	return
}

// ValidateWineryResponse runs the validations defined on WineryResponse
func ValidateWineryResponse(body *WineryResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "[a-zA-Z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "[a-zA-Z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponentResponse runs the validations defined on ComponentResponse
func ValidateComponentResponse(body *ComponentResponse) (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// ValidateWineryResponseBody runs the validations defined on WineryResponseBody
func ValidateWineryResponseBody(body *WineryResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "[a-zA-Z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "[a-zA-Z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponentResponseBody runs the validations defined on
// ComponentResponseBody
func ValidateComponentResponseBody(body *ComponentResponseBody) (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// ValidateWineryRequestBody runs the validations defined on WineryRequestBody
func ValidateWineryRequestBody(body *WineryRequestBody) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("body.region", body.Region, "[a-zA-Z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("body.country", body.Country, "[a-zA-Z '\\.]+"))
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponentRequestBody runs the validations defined on
// ComponentRequestBody
func ValidateComponentRequestBody(body *ComponentRequestBody) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", body.Varietal, "[A-Za-z' ]+"))
	if utf8.RuneCountInString(body.Varietal) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", body.Varietal, utf8.RuneCountInString(body.Varietal), 100, false))
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// ValidateBottleRequestBody runs the validations defined on BottleRequestBody
func ValidateBottleRequestBody(body *BottleRequestBody) (err error) {
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
	return
}
