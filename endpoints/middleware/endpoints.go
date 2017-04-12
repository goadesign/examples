package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/goadesign/goa"
)

// Private type used to store user info in context.
type key int

// Context key used to store user info.
const userKey key = 1

// User is the type that matches the data serialized into the user info header by Google Endpoints.
type User struct {
	// User email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User ID
	ID string `form:"id" json:"id" xml:"id"`
	// Token issuer
	Issuer *string `form:"issuer,omitempty" json:"issuer,omitempty" xml:"issuer,omitempty"`
}

// Endpoints returns a goa middleware that extracts the user information initialized by
// Google Cloud Endpoints and stores it in the context. Use UserInfo to retrieve it.
// See https://cloud.google.com/endpoints/docs/authenticating-users
func Endpoints() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			logger := goa.ContextLogger(ctx)
			info := req.Header.Get("X-Endpoint-API-UserInfo")
			logger.Info("endpoints middleware", "header", info)
			u := User{ID: "anonymous"}
			if info != "" {
				js, err := base64.RawStdEncoding.DecodeString(info)
				if err != nil {
					logger.Error("invalid header Base64 encoding", "err", err)
				} else {
					if err = json.Unmarshal(js, &u); err != nil {
						logger.Error("invalid header JSON", "err", err)
					}
				}
			}
			ctx = context.WithValue(ctx, userKey, &u)
			return h(ctx, rw, req)
		}
	}
}

// UserInfo returns the user information provided by Google Endpoints if any.
func UserInfo(ctx context.Context) *User {
	u := ctx.Value(userKey)
	if u == nil {
		return nil
	}
	return u.(*User)
}
