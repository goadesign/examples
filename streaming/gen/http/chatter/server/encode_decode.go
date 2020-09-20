// Code generated by goa v3.2.4, DO NOT EDIT.
//
// chatter HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package server

import (
	"context"
	"net/http"
	"strings"

	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeLoginResponse returns an encoder for responses returned by the chatter
// login endpoint.
func EncodeLoginResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginRequest returns a decoder for requests sent to the chatter login
// endpoint.
func DecodeLoginRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		payload := NewLoginPayload()
		user, pass, ok := r.BasicAuth()
		if !ok {
			return nil, goa.MissingFieldError("Authorization", "header")
		}
		payload.User = user
		payload.Password = pass

		return payload, nil
	}
}

// EncodeLoginError returns an encoder for errors returned by the login chatter
// endpoint.
func EncodeLoginError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(chatter.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewLoginUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeEchoerRequest returns a decoder for requests sent to the chatter
// echoer endpoint.
func DecodeEchoerRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewEchoerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeEchoerError returns an encoder for errors returned by the echoer
// chatter endpoint.
func EncodeEchoerError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(chatter.InvalidScopes)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewEchoerInvalidScopesResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(chatter.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewEchoerUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeListenerRequest returns a decoder for requests sent to the chatter
// listener endpoint.
func DecodeListenerRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListenerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeListenerError returns an encoder for errors returned by the listener
// chatter endpoint.
func EncodeListenerError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(chatter.InvalidScopes)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListenerInvalidScopesResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(chatter.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListenerUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeSummaryRequest returns a decoder for requests sent to the chatter
// summary endpoint.
func DecodeSummaryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSummaryPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeSummaryError returns an encoder for errors returned by the summary
// chatter endpoint.
func EncodeSummaryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(chatter.InvalidScopes)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSummaryInvalidScopesResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(chatter.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSummaryUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeSubscribeRequest returns a decoder for requests sent to the chatter
// subscribe endpoint.
func DecodeSubscribeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSubscribePayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeSubscribeError returns an encoder for errors returned by the subscribe
// chatter endpoint.
func EncodeSubscribeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(chatter.InvalidScopes)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSubscribeInvalidScopesResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(chatter.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSubscribeUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeHistoryRequest returns a decoder for requests sent to the chatter
// history endpoint.
func DecodeHistoryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			view  *string
			token string
			err   error
		)
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewHistoryPayload(view, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeHistoryError returns an encoder for errors returned by the history
// chatter endpoint.
func EncodeHistoryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(chatter.InvalidScopes)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHistoryInvalidScopesResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(chatter.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHistoryUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalChatterviewsChatSummaryViewToChatSummaryResponse builds a value of
// type *ChatSummaryResponse from a value of type *chatterviews.ChatSummaryView.
func marshalChatterviewsChatSummaryViewToChatSummaryResponse(v *chatterviews.ChatSummaryView) *ChatSummaryResponse {
	res := &ChatSummaryResponse{
		Message: *v.Message,
		Length:  v.Length,
		SentAt:  *v.SentAt,
	}

	return res
}
