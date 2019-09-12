// Code generated by goa v3.0.6, DO NOT EDIT.
//
// sommelier HTTP server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package server

import (
	"unicode/utf8"

	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goa "goa.design/goa/v3/pkg"
)

// PickRequestBody is the type of the "sommelier" service "pick" endpoint HTTP
// request body.
type PickRequestBody struct {
	// Name of bottle to pick
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery *string `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
}

// StoredBottleResponseCollection is the type of the "sommelier" service "pick"
// endpoint HTTP response body.
type StoredBottleResponseCollection []*StoredBottleResponse

// PickNoCriteriaResponseBody is the type of the "sommelier" service "pick"
// endpoint HTTP response body for the "no_criteria" error.
type PickNoCriteriaResponseBody string

// PickNoMatchResponseBody is the type of the "sommelier" service "pick"
// endpoint HTTP response body for the "no_match" error.
type PickNoMatchResponseBody string

// StoredBottleResponse is used to define fields on response body types.
type StoredBottleResponse struct {
	// ID is the unique id of the bottle.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryResponseTiny `form:"winery" json:"winery" xml:"winery"`
	// Vintage of bottle
	Vintage uint32 `form:"vintage" json:"vintage" xml:"vintage"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponse `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// WineryResponseTiny is used to define fields on response body types.
type WineryResponseTiny struct {
	// Name of winery
	Name string `form:"name" json:"name" xml:"name"`
}

// ComponentResponse is used to define fields on response body types.
type ComponentResponse struct {
	// Grape varietal
	Varietal string `form:"varietal" json:"varietal" xml:"varietal"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// NewStoredBottleResponseCollection builds the HTTP response body from the
// result of the "pick" endpoint of the "sommelier" service.
func NewStoredBottleResponseCollection(res sommelierviews.StoredBottleCollectionView) StoredBottleResponseCollection {
	body := make([]*StoredBottleResponse, len(res))
	for i, val := range res {
		body[i] = &StoredBottleResponse{
			ID:          *val.ID,
			Name:        *val.Name,
			Vintage:     *val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		if val.Winery != nil {
			body[i].Winery = marshalSommelierviewsWineryViewToWineryResponseTiny(val.Winery)
		}
		if val.Composition != nil {
			body[i].Composition = make([]*ComponentResponse, len(val.Composition))
			for j, val := range val.Composition {
				body[i].Composition[j] = marshalSommelierviewsComponentViewToComponentResponse(val)
			}
		}
	}
	return body
}

// NewPickNoCriteriaResponseBody builds the HTTP response body from the result
// of the "pick" endpoint of the "sommelier" service.
func NewPickNoCriteriaResponseBody(res sommelier.NoCriteria) PickNoCriteriaResponseBody {
	body := PickNoCriteriaResponseBody(res)
	return body
}

// NewPickNoMatchResponseBody builds the HTTP response body from the result of
// the "pick" endpoint of the "sommelier" service.
func NewPickNoMatchResponseBody(res sommelier.NoMatch) PickNoMatchResponseBody {
	body := PickNoMatchResponseBody(res)
	return body
}

// NewPickCriteria builds a sommelier service pick endpoint payload.
func NewPickCriteria(body *PickRequestBody) *sommelier.Criteria {
	v := &sommelier.Criteria{
		Name:   body.Name,
		Winery: body.Winery,
	}
	if body.Varietal != nil {
		v.Varietal = make([]string, len(body.Varietal))
		for i, val := range body.Varietal {
			v.Varietal[i] = val
		}
	}
	return v
}

// ValidateStoredBottleResponse runs the validations defined on
// StoredBottleResponse
func ValidateStoredBottleResponse(body *StoredBottleResponse) (err error) {
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if utf8.RuneCountInString(body.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
	}
	if body.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", body.Vintage, 1900, true))
	}
	if body.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", body.Vintage, 2020, false))
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

// ValidateComponentResponse runs the validations defined on ComponentResponse
func ValidateComponentResponse(body *ComponentResponse) (err error) {
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
