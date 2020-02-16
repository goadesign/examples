package multiauth

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type authInfo struct {
	user   string
	claims jwt.MapClaims
	key    string
}

func (auth authInfo) String() string {
	if auth.user != "" {
		return fmt.Sprintf("AuthInfo: Username + Password")
	} else if auth.claims != nil {
		return fmt.Sprintf("AuthInfo: JWT/OAuth")
	} else if auth.key != "" {
		return fmt.Sprintf("AuthInfo: API")
	} else {
		return fmt.Sprintf("AuthInfo: none")
	}
}

type ctxValue int

const (
	ctxValueClaims ctxValue = iota
)

// contextWithAuthInfo adds the given JWT claims to the context and returns it.
func contextWithAuthInfo(ctx context.Context, auth authInfo) context.Context {
	return context.WithValue(ctx, ctxValueClaims, auth)
}

// contextAuthInfo returns the jwt.MapClaims from the given context.
func contextAuthInfo(ctx context.Context) (auth authInfo) {
	auth, _ = ctx.Value(ctxValueClaims).(authInfo)
	return
}
