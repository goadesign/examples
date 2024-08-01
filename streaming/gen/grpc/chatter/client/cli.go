// Code generated by goa v3.18.2, DO NOT EDIT.
//
// chatter gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/streaming/design

package client

import (
	chatter "goa.design/examples/streaming/gen/chatter"
)

// BuildLoginPayload builds the payload for the chatter login endpoint from CLI
// flags.
func BuildLoginPayload(chatterLoginUser string, chatterLoginPassword string) (*chatter.LoginPayload, error) {
	var user string
	{
		user = chatterLoginUser
	}
	var password string
	{
		password = chatterLoginPassword
	}
	v := &chatter.LoginPayload{}
	v.User = user
	v.Password = password

	return v, nil
}

// BuildEchoerPayload builds the payload for the chatter echoer endpoint from
// CLI flags.
func BuildEchoerPayload(chatterEchoerToken string) (*chatter.EchoerPayload, error) {
	var token string
	{
		token = chatterEchoerToken
	}
	v := &chatter.EchoerPayload{}
	v.Token = token

	return v, nil
}

// BuildListenerPayload builds the payload for the chatter listener endpoint
// from CLI flags.
func BuildListenerPayload(chatterListenerToken string) (*chatter.ListenerPayload, error) {
	var token string
	{
		token = chatterListenerToken
	}
	v := &chatter.ListenerPayload{}
	v.Token = token

	return v, nil
}

// BuildSummaryPayload builds the payload for the chatter summary endpoint from
// CLI flags.
func BuildSummaryPayload(chatterSummaryToken string) (*chatter.SummaryPayload, error) {
	var token string
	{
		token = chatterSummaryToken
	}
	v := &chatter.SummaryPayload{}
	v.Token = token

	return v, nil
}

// BuildSubscribePayload builds the payload for the chatter subscribe endpoint
// from CLI flags.
func BuildSubscribePayload(chatterSubscribeToken string) (*chatter.SubscribePayload, error) {
	var token string
	{
		token = chatterSubscribeToken
	}
	v := &chatter.SubscribePayload{}
	v.Token = token

	return v, nil
}

// BuildHistoryPayload builds the payload for the chatter history endpoint from
// CLI flags.
func BuildHistoryPayload(chatterHistoryView string, chatterHistoryToken string) (*chatter.HistoryPayload, error) {
	var view *string
	{
		if chatterHistoryView != "" {
			view = &chatterHistoryView
		}
	}
	var token string
	{
		token = chatterHistoryToken
	}
	v := &chatter.HistoryPayload{}
	v.View = view
	v.Token = token

	return v, nil
}
