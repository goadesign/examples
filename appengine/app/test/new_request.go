package test

import (
	"context"
	"io"
	"net/http"
)

type NewRequestFunc func(string, string, io.Reader) (*http.Request, error)

type keyType string

const newRequestFuncKey keyType = "NewRequestFunc"

func WithNewRequestFunc(ctx context.Context, f NewRequestFunc) context.Context {
	return context.WithValue(ctx, newRequestFuncKey, f)
}

func getNewRequestFunc(ctx context.Context) NewRequestFunc {
	v := ctx.Value(newRequestFuncKey)
	if v == nil {
		return http.NewRequest
	}
	return v.(NewRequestFunc)
}

func NewRequestWith(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return getNewRequestFunc(ctx)(method, url, body)
}
