// Code generated by goa v3.15.1, DO NOT EDIT.
//
// chatter gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/streaming/design

package server

import (
	"context"
	"strings"

	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeLoginResponse encodes responses from the "chatter" service "login"
// endpoint.
func EncodeLoginResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "login", "string", v)
	}
	resp := NewProtoLoginResponse(result)
	return resp, nil
}

// DecodeLoginRequest decodes requests sent to "chatter" service "login"
// endpoint.
func DecodeLoginRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		user     string
		password string
		err      error
	)
	{
		if vals := md.Get("user"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("user", "metadata"))
		} else {
			user = vals[0]
		}
		if vals := md.Get("password"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("password", "metadata"))
		} else {
			password = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.LoginPayload
	{
		payload = NewLoginPayload(user, password)
	}
	return payload, nil
}

// EncodeEchoerResponse encodes responses from the "chatter" service "echoer"
// endpoint.
func EncodeEchoerResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "echoer", "string", v)
	}
	resp := NewProtoEchoerResponse(result)
	return resp, nil
}

// DecodeEchoerRequest decodes requests sent to "chatter" service "echoer"
// endpoint.
func DecodeEchoerRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.EchoerPayload
	{
		payload = NewEchoerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeListenerResponse encodes responses from the "chatter" service
// "listener" endpoint.
func EncodeListenerResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoListenerResponse()
	return resp, nil
}

// DecodeListenerRequest decodes requests sent to "chatter" service "listener"
// endpoint.
func DecodeListenerRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.ListenerPayload
	{
		payload = NewListenerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeSummaryResponse encodes responses from the "chatter" service "summary"
// endpoint.
func EncodeSummaryResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(chatterviews.ChatSummaryCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "summary", "chatterviews.ChatSummaryCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoChatSummaryCollection(result)
	return resp, nil
}

// DecodeSummaryRequest decodes requests sent to "chatter" service "summary"
// endpoint.
func DecodeSummaryRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.SummaryPayload
	{
		payload = NewSummaryPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeSubscribeResponse encodes responses from the "chatter" service
// "subscribe" endpoint.
func EncodeSubscribeResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*chatter.Event)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "subscribe", "*chatter.Event", v)
	}
	resp := NewProtoSubscribeResponse(result)
	return resp, nil
}

// DecodeSubscribeRequest decodes requests sent to "chatter" service
// "subscribe" endpoint.
func DecodeSubscribeRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.SubscribePayload
	{
		payload = NewSubscribePayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// EncodeHistoryResponse encodes responses from the "chatter" service "history"
// endpoint.
func EncodeHistoryResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*chatterviews.ChatSummary)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "history", "*chatterviews.ChatSummary", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoHistoryResponse(result)
	return resp, nil
}

// DecodeHistoryRequest decodes requests sent to "chatter" service "history"
// endpoint.
func DecodeHistoryRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		view  *string
		token string
		err   error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.HistoryPayload
	{
		payload = NewHistoryPayload(view, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}
