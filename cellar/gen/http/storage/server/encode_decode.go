// Code generated by goa v3.14.2, DO NOT EDIT.
//
// storage HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the storage
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(storageviews.StoredBottleCollection)
		enc := encoder(ctx, w)
		body := NewStoredBottleResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeShowResponse returns an encoder for responses returned by the storage
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*storageviews.StoredBottle)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body any
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "tiny":
			body = NewShowResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the storage show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id   string
			view *string
			err  error

			params = mux.Vars(r)
		)
		id = params["id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []any{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id, view)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show storage
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *storage.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddResponse returns an encoder for responses returned by the storage
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the storage add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddBottle(&body)

		return payload, nil
	}
}

// EncodeRemoveResponse returns an encoder for responses returned by the
// storage remove endpoint.
func EncodeRemoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveRequest returns a decoder for requests sent to the storage
// remove endpoint.
func DecodeRemoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewRemovePayload(id)

		return payload, nil
	}
}

// EncodeRateResponse returns an encoder for responses returned by the storage
// rate endpoint.
func EncodeRateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeRateRequest returns a decoder for requests sent to the storage rate
// endpoint.
func DecodeRateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			query map[uint32][]string
			err   error
		)
		{
			queryRaw := r.URL.Query()
			if len(queryRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("query", "query string"))
			}
			for keyRaw, valRaw := range queryRaw {
				if strings.HasPrefix(keyRaw, "query[") {
					if query == nil {
						query = make(map[uint32][]string)
					}
					var keya uint32
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseUint(keyaRaw, 10, 32)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("query", keyaRaw, "unsigned integer"))
						}
						keya = uint32(v)
					}
					query[keya] = valRaw
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := query

		return payload, nil
	}
}

// EncodeMultiAddResponse returns an encoder for responses returned by the
// storage multi_add endpoint.
func EncodeMultiAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMultiAddRequest returns a decoder for requests sent to the storage
// multi_add endpoint.
func DecodeMultiAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var payload []*storage.Bottle
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewStorageMultiAddDecoder returns a decoder to decode the multipart request
// for the "storage" service "multi_add" endpoint.
func NewStorageMultiAddDecoder(mux goahttp.Muxer, storageMultiAddDecoderFn StorageMultiAddDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v any) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(*[]*storage.Bottle)
			if err := storageMultiAddDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}

// EncodeMultiUpdateResponse returns an encoder for responses returned by the
// storage multi_update endpoint.
func EncodeMultiUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeMultiUpdateRequest returns a decoder for requests sent to the storage
// multi_update endpoint.
func DecodeMultiUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var payload *storage.MultiUpdatePayload
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewStorageMultiUpdateDecoder returns a decoder to decode the multipart
// request for the "storage" service "multi_update" endpoint.
func NewStorageMultiUpdateDecoder(mux goahttp.Muxer, storageMultiUpdateDecoderFn StorageMultiUpdateDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v any) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(**storage.MultiUpdatePayload)
			if err := storageMultiUpdateDecoderFn(mr, p); err != nil {
				return err
			}

			var (
				ids []string
				err error
			)
			ids = r.URL.Query()["ids"]
			if ids == nil {
				err = goa.MergeErrors(err, goa.MissingFieldError("ids", "query string"))
			}
			if err != nil {
				return err
			}
			(*p).Ids = ids
			return nil
		})
	}
}

// marshalStorageviewsStoredBottleViewToStoredBottleResponseTiny builds a value
// of type *StoredBottleResponseTiny from a value of type
// *storageviews.StoredBottleView.
func marshalStorageviewsStoredBottleViewToStoredBottleResponseTiny(v *storageviews.StoredBottleView) *StoredBottleResponseTiny {
	res := &StoredBottleResponseTiny{
		ID:   *v.ID,
		Name: *v.Name,
	}
	if v.Winery != nil {
		res.Winery = marshalStorageviewsWineryViewToWineryResponseTiny(v.Winery)
	}

	return res
}

// marshalStorageviewsWineryViewToWineryResponseTiny builds a value of type
// *WineryResponseTiny from a value of type *storageviews.WineryView.
func marshalStorageviewsWineryViewToWineryResponseTiny(v *storageviews.WineryView) *WineryResponseTiny {
	res := &WineryResponseTiny{
		Name: *v.Name,
	}

	return res
}

// marshalStorageviewsWineryViewToWineryResponseBodyTiny builds a value of type
// *WineryResponseBodyTiny from a value of type *storageviews.WineryView.
func marshalStorageviewsWineryViewToWineryResponseBodyTiny(v *storageviews.WineryView) *WineryResponseBodyTiny {
	res := &WineryResponseBodyTiny{
		Name: *v.Name,
	}

	return res
}

// marshalStorageviewsComponentViewToComponentResponseBody builds a value of
// type *ComponentResponseBody from a value of type *storageviews.ComponentView.
func marshalStorageviewsComponentViewToComponentResponseBody(v *storageviews.ComponentView) *ComponentResponseBody {
	if v == nil {
		return nil
	}
	res := &ComponentResponseBody{
		Varietal:   *v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// unmarshalWineryRequestBodyToStorageWinery builds a value of type
// *storage.Winery from a value of type *WineryRequestBody.
func unmarshalWineryRequestBodyToStorageWinery(v *WineryRequestBody) *storage.Winery {
	res := &storage.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		URL:     v.URL,
	}

	return res
}

// unmarshalComponentRequestBodyToStorageComponent builds a value of type
// *storage.Component from a value of type *ComponentRequestBody.
func unmarshalComponentRequestBodyToStorageComponent(v *ComponentRequestBody) *storage.Component {
	if v == nil {
		return nil
	}
	res := &storage.Component{
		Varietal:   *v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// unmarshalBottleRequestBodyToStorageBottle builds a value of type
// *storage.Bottle from a value of type *BottleRequestBody.
func unmarshalBottleRequestBodyToStorageBottle(v *BottleRequestBody) *storage.Bottle {
	res := &storage.Bottle{
		Name:        *v.Name,
		Vintage:     *v.Vintage,
		Description: v.Description,
		Rating:      v.Rating,
	}
	res.Winery = unmarshalWineryRequestBodyToStorageWinery(v.Winery)
	if v.Composition != nil {
		res.Composition = make([]*storage.Component, len(v.Composition))
		for i, val := range v.Composition {
			res.Composition[i] = unmarshalComponentRequestBodyToStorageComponent(val)
		}
	}

	return res
}
