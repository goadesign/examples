// Code generated by goa v3.18.2, DO NOT EDIT.
//
// sommelier gRPC client types
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package client

import (
	"unicode/utf8"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goa "goa.design/goa/v3/pkg"
)

// NewProtoPickRequest builds the gRPC request type from the payload of the
// "pick" endpoint of the "sommelier" service.
func NewProtoPickRequest(payload *sommelier.Criteria) *sommelierpb.PickRequest {
	message := &sommelierpb.PickRequest{
		Name:   payload.Name,
		Winery: payload.Winery,
	}
	if payload.Varietal != nil {
		message.Varietal = make([]string, len(payload.Varietal))
		for i, val := range payload.Varietal {
			message.Varietal[i] = val
		}
	}
	return message
}

// NewPickResult builds the result type of the "pick" endpoint of the
// "sommelier" service from the gRPC response type.
func NewPickResult(message *sommelierpb.StoredBottleCollection) sommelierviews.StoredBottleCollectionView {
	result := make([]*sommelierviews.StoredBottleView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &sommelierviews.StoredBottleView{
			ID:          &val.Id,
			Name:        &val.Name,
			Vintage:     &val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		if val.Winery != nil {
			result[i].Winery = protobufSommelierpbWineryToSommelierviewsWineryView(val.Winery)
		}
		if val.Composition != nil {
			result[i].Composition = make([]*sommelierviews.ComponentView, len(val.Composition))
			for j, val := range val.Composition {
				result[i].Composition[j] = &sommelierviews.ComponentView{
					Varietal:   &val.Varietal,
					Percentage: val.Percentage,
				}
			}
		}
	}
	return result
}

// ValidateStoredBottleCollection runs the validations defined on
// StoredBottleCollection.
func ValidateStoredBottleCollection(message *sommelierpb.StoredBottleCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredBottle runs the validations defined on StoredBottle.
func ValidateStoredBottle(elem *sommelierpb.StoredBottle) (err error) {
	if elem.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "elem"))
	}
	if utf8.RuneCountInString(elem.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("elem.name", elem.Name, utf8.RuneCountInString(elem.Name), 100, false))
	}
	if elem.Winery != nil {
		if err2 := ValidateWinery(elem.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if elem.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("elem.vintage", elem.Vintage, 1900, true))
	}
	if elem.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("elem.vintage", elem.Vintage, 2020, false))
	}
	for _, e := range elem.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if elem.Description != nil {
		if utf8.RuneCountInString(*elem.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("elem.description", *elem.Description, utf8.RuneCountInString(*elem.Description), 2000, false))
		}
	}
	if elem.Rating != nil {
		if *elem.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.rating", *elem.Rating, 1, true))
		}
	}
	if elem.Rating != nil {
		if *elem.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.rating", *elem.Rating, 5, false))
		}
	}
	return
}

// ValidateWinery runs the validations defined on Winery.
func ValidateWinery(winery *sommelierpb.Winery) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("winery.region", winery.Region, "[a-zA-Z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("winery.country", winery.Country, "[a-zA-Z '\\.]+"))
	if winery.Url != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("winery.url", *winery.Url, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponent runs the validations defined on Component.
func ValidateComponent(elem *sommelierpb.Component) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("elem.varietal", elem.Varietal, "[A-Za-z' ]+"))
	if utf8.RuneCountInString(elem.Varietal) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("elem.varietal", elem.Varietal, utf8.RuneCountInString(elem.Varietal), 100, false))
	}
	if elem.Percentage != nil {
		if *elem.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.percentage", *elem.Percentage, 1, true))
		}
	}
	if elem.Percentage != nil {
		if *elem.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.percentage", *elem.Percentage, 100, false))
		}
	}
	return
}

// svcSommelierviewsWineryViewToSommelierpbWinery builds a value of type
// *sommelierpb.Winery from a value of type *sommelierviews.WineryView.
func svcSommelierviewsWineryViewToSommelierpbWinery(v *sommelierviews.WineryView) *sommelierpb.Winery {
	res := &sommelierpb.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		Url:     v.URL,
	}

	return res
}

// protobufSommelierpbWineryToSommelierviewsWineryView builds a value of type
// *sommelierviews.WineryView from a value of type *sommelierpb.Winery.
func protobufSommelierpbWineryToSommelierviewsWineryView(v *sommelierpb.Winery) *sommelierviews.WineryView {
	res := &sommelierviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
		URL:     v.Url,
	}

	return res
}
