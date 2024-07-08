// Code generated by goa v3.17.2, DO NOT EDIT.
//
// chatter HTTP server types
//
// Command:
// $ goa gen goa.design/examples/streaming/design

package server

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
)

// ChatSummaryResponseCollection is the type of the "chatter" service "summary"
// endpoint HTTP response body.
type ChatSummaryResponseCollection []*ChatSummaryResponse

// SubscribeResponseBody is the type of the "chatter" service "subscribe"
// endpoint HTTP response body.
type SubscribeResponseBody struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
	Action  string `form:"action" json:"action" xml:"action"`
	// Time at which the message was added
	AddedAt string `form:"added_at" json:"added_at" xml:"added_at"`
}

// HistoryResponseBodyTiny is the type of the "chatter" service "history"
// endpoint HTTP response body.
type HistoryResponseBodyTiny struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
}

// HistoryResponseBody is the type of the "chatter" service "history" endpoint
// HTTP response body.
type HistoryResponseBody struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt string `form:"sent_at" json:"sent_at" xml:"sent_at"`
}

// ChatSummaryResponse is used to define fields on response body types.
type ChatSummaryResponse struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt string `form:"sent_at" json:"sent_at" xml:"sent_at"`
}

// NewChatSummaryResponseCollection builds the HTTP response body from the
// result of the "summary" endpoint of the "chatter" service.
func NewChatSummaryResponseCollection(res chatterviews.ChatSummaryCollectionView) ChatSummaryResponseCollection {
	body := make([]*ChatSummaryResponse, len(res))
	for i, val := range res {
		body[i] = marshalChatterviewsChatSummaryViewToChatSummaryResponse(val)
	}
	return body
}

// NewSubscribeResponseBody builds the HTTP response body from the result of
// the "subscribe" endpoint of the "chatter" service.
func NewSubscribeResponseBody(res *chatter.Event) *SubscribeResponseBody {
	body := &SubscribeResponseBody{
		Message: res.Message,
		Action:  res.Action,
		AddedAt: res.AddedAt,
	}
	return body
}

// NewHistoryResponseBodyTiny builds the HTTP response body from the result of
// the "history" endpoint of the "chatter" service.
func NewHistoryResponseBodyTiny(res *chatterviews.ChatSummaryView) *HistoryResponseBodyTiny {
	body := &HistoryResponseBodyTiny{
		Message: *res.Message,
	}
	return body
}

// NewHistoryResponseBody builds the HTTP response body from the result of the
// "history" endpoint of the "chatter" service.
func NewHistoryResponseBody(res *chatterviews.ChatSummaryView) *HistoryResponseBody {
	body := &HistoryResponseBody{
		Message: *res.Message,
		Length:  res.Length,
		SentAt:  *res.SentAt,
	}
	return body
}

// NewLoginPayload builds a chatter service login endpoint payload.
func NewLoginPayload() *chatter.LoginPayload {
	v := &chatter.LoginPayload{}

	return v
}

// NewEchoerPayload builds a chatter service echoer endpoint payload.
func NewEchoerPayload(token string) *chatter.EchoerPayload {
	v := &chatter.EchoerPayload{}
	v.Token = token

	return v
}

// NewListenerPayload builds a chatter service listener endpoint payload.
func NewListenerPayload(token string) *chatter.ListenerPayload {
	v := &chatter.ListenerPayload{}
	v.Token = token

	return v
}

// NewSummaryPayload builds a chatter service summary endpoint payload.
func NewSummaryPayload(token string) *chatter.SummaryPayload {
	v := &chatter.SummaryPayload{}
	v.Token = token

	return v
}

// NewSubscribePayload builds a chatter service subscribe endpoint payload.
func NewSubscribePayload(token string) *chatter.SubscribePayload {
	v := &chatter.SubscribePayload{}
	v.Token = token

	return v
}

// NewHistoryPayload builds a chatter service history endpoint payload.
func NewHistoryPayload(view *string, token string) *chatter.HistoryPayload {
	v := &chatter.HistoryPayload{}
	v.View = view
	v.Token = token

	return v
}
