// Code generated by goa v3.14.5, DO NOT EDIT.
//
// storage gRPC client types
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package client

import (
	"unicode/utf8"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goa "goa.design/goa/v3/pkg"
)

// NewProtoListRequest builds the gRPC request type from the payload of the
// "list" endpoint of the "storage" service.
func NewProtoListRequest() *storagepb.ListRequest {
	message := &storagepb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the "storage"
// service from the gRPC response type.
func NewListResult(message *storagepb.StoredBottleCollection) storageviews.StoredBottleCollectionView {
	result := make([]*storageviews.StoredBottleView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &storageviews.StoredBottleView{
			ID:          &val.Id,
			Name:        &val.Name,
			Vintage:     &val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		if val.Winery != nil {
			result[i].Winery = protobufStoragepbWineryToStorageviewsWineryView(val.Winery)
		}
		if val.Composition != nil {
			result[i].Composition = make([]*storageviews.ComponentView, len(val.Composition))
			for j, val := range val.Composition {
				result[i].Composition[j] = &storageviews.ComponentView{
					Varietal:   &val.Varietal,
					Percentage: val.Percentage,
				}
			}
		}
	}
	return result
}

// NewProtoShowRequest builds the gRPC request type from the payload of the
// "show" endpoint of the "storage" service.
func NewProtoShowRequest(payload *storage.ShowPayload) *storagepb.ShowRequest {
	message := &storagepb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the "storage"
// service from the gRPC response type.
func NewShowResult(message *storagepb.ShowResponse) *storageviews.StoredBottleView {
	result := &storageviews.StoredBottleView{
		ID:          &message.Id,
		Name:        &message.Name,
		Vintage:     &message.Vintage,
		Description: message.Description,
		Rating:      message.Rating,
	}
	if message.Winery != nil {
		result.Winery = protobufStoragepbWineryToStorageviewsWineryView(message.Winery)
	}
	if message.Composition != nil {
		result.Composition = make([]*storageviews.ComponentView, len(message.Composition))
		for i, val := range message.Composition {
			result.Composition[i] = &storageviews.ComponentView{
				Varietal:   &val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "storage" service from the gRPC error response type.
func NewShowNotFoundError(message *storagepb.ShowNotFoundError) *storage.NotFound {
	er := &storage.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewProtoAddRequest builds the gRPC request type from the payload of the
// "add" endpoint of the "storage" service.
func NewProtoAddRequest(payload *storage.Bottle) *storagepb.AddRequest {
	message := &storagepb.AddRequest{
		Name:        payload.Name,
		Vintage:     payload.Vintage,
		Description: payload.Description,
		Rating:      payload.Rating,
	}
	if payload.Winery != nil {
		message.Winery = svcStorageWineryToStoragepbWinery(payload.Winery)
	}
	if payload.Composition != nil {
		message.Composition = make([]*storagepb.Component, len(payload.Composition))
		for i, val := range payload.Composition {
			message.Composition[i] = &storagepb.Component{
				Varietal:   val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "storage"
// service from the gRPC response type.
func NewAddResult(message *storagepb.AddResponse) string {
	result := message.Field
	return result
}

// NewProtoRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "storage" service.
func NewProtoRemoveRequest(payload *storage.RemovePayload) *storagepb.RemoveRequest {
	message := &storagepb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewProtoRateRequest builds the gRPC request type from the payload of the
// "rate" endpoint of the "storage" service.
func NewProtoRateRequest(payload map[uint32][]string) *storagepb.RateRequest {
	message := &storagepb.RateRequest{}
	message.Field = make(map[uint32]*storagepb.ArrayOfString, len(payload))
	for key, val := range payload {
		tk := key
		tv := &storagepb.ArrayOfString{}
		tv.Field = make([]string, len(val))
		for i, val := range val {
			tv.Field[i] = val
		}
		message.Field[tk] = tv
	}
	return message
}

// NewProtoMultiAddRequest builds the gRPC request type from the payload of the
// "multi_add" endpoint of the "storage" service.
func NewProtoMultiAddRequest(payload []*storage.Bottle) *storagepb.MultiAddRequest {
	message := &storagepb.MultiAddRequest{}
	message.Field = make([]*storagepb.Bottle, len(payload))
	for i, val := range payload {
		message.Field[i] = &storagepb.Bottle{
			Name:        val.Name,
			Vintage:     val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		if val.Winery != nil {
			message.Field[i].Winery = svcStorageWineryToStoragepbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*storagepb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &storagepb.Component{
					Varietal:   val.Varietal,
					Percentage: val.Percentage,
				}
			}
		}
	}
	return message
}

// NewMultiAddResult builds the result type of the "multi_add" endpoint of the
// "storage" service from the gRPC response type.
func NewMultiAddResult(message *storagepb.MultiAddResponse) []string {
	result := make([]string, len(message.Field))
	for i, val := range message.Field {
		result[i] = val
	}
	return result
}

// NewProtoMultiUpdateRequest builds the gRPC request type from the payload of
// the "multi_update" endpoint of the "storage" service.
func NewProtoMultiUpdateRequest(payload *storage.MultiUpdatePayload) *storagepb.MultiUpdateRequest {
	message := &storagepb.MultiUpdateRequest{}
	if payload.Ids != nil {
		message.Ids = make([]string, len(payload.Ids))
		for i, val := range payload.Ids {
			message.Ids[i] = val
		}
	}
	if payload.Bottles != nil {
		message.Bottles = make([]*storagepb.Bottle, len(payload.Bottles))
		for i, val := range payload.Bottles {
			message.Bottles[i] = &storagepb.Bottle{
				Name:        val.Name,
				Vintage:     val.Vintage,
				Description: val.Description,
				Rating:      val.Rating,
			}
			if val.Winery != nil {
				message.Bottles[i].Winery = svcStorageWineryToStoragepbWinery(val.Winery)
			}
			if val.Composition != nil {
				message.Bottles[i].Composition = make([]*storagepb.Component, len(val.Composition))
				for j, val := range val.Composition {
					message.Bottles[i].Composition[j] = &storagepb.Component{
						Varietal:   val.Varietal,
						Percentage: val.Percentage,
					}
				}
			}
		}
	}
	return message
}

// ValidateStoredBottleCollection runs the validations defined on
// StoredBottleCollection.
func ValidateStoredBottleCollection(message *storagepb.StoredBottleCollection) (err error) {
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
func ValidateStoredBottle(elem *storagepb.StoredBottle) (err error) {
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
func ValidateWinery(winery *storagepb.Winery) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("winery.region", winery.Region, "[a-zA-Z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("winery.country", winery.Country, "[a-zA-Z '\\.]+"))
	if winery.Url != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("winery.url", *winery.Url, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponent runs the validations defined on Component.
func ValidateComponent(elem *storagepb.Component) (err error) {
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

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *storagepb.ShowResponse) (err error) {
	if message.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "message"))
	}
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if message.Description != nil {
		if utf8.RuneCountInString(*message.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", *message.Description, utf8.RuneCountInString(*message.Description), 2000, false))
		}
	}
	if message.Rating != nil {
		if *message.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", *message.Rating, 1, true))
		}
	}
	if message.Rating != nil {
		if *message.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", *message.Rating, 5, false))
		}
	}
	return
}

// svcStorageviewsWineryViewToStoragepbWinery builds a value of type
// *storagepb.Winery from a value of type *storageviews.WineryView.
func svcStorageviewsWineryViewToStoragepbWinery(v *storageviews.WineryView) *storagepb.Winery {
	res := &storagepb.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		Url:     v.URL,
	}

	return res
}

// protobufStoragepbWineryToStorageviewsWineryView builds a value of type
// *storageviews.WineryView from a value of type *storagepb.Winery.
func protobufStoragepbWineryToStorageviewsWineryView(v *storagepb.Winery) *storageviews.WineryView {
	res := &storageviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
		URL:     v.Url,
	}

	return res
}

// protobufStoragepbWineryToStorageWinery builds a value of type
// *storage.Winery from a value of type *storagepb.Winery.
func protobufStoragepbWineryToStorageWinery(v *storagepb.Winery) *storage.Winery {
	res := &storage.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.Url,
	}

	return res
}

// svcStorageWineryToStoragepbWinery builds a value of type *storagepb.Winery
// from a value of type *storage.Winery.
func svcStorageWineryToStoragepbWinery(v *storage.Winery) *storagepb.Winery {
	res := &storagepb.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		Url:     v.URL,
	}

	return res
}
