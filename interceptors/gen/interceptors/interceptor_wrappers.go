// Code generated by goa v3.19.1, DO NOT EDIT.
//
// Interceptor wrappers
//
// Command:
// $ goa gen goa.design/examples/interceptors/design

package interceptors

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// wrappedStreamServerStream is a server interceptor wrapper for the
// StreamServerStream stream.
type wrappedStreamServerStream struct {
	ctx             context.Context
	sendWithContext func(context.Context, *StreamResult) error
	recvWithContext func(context.Context) (*StreamStreamingPayload, error)
	stream          StreamServerStream
}

// wrappedStreamClientStream is a client interceptor wrapper for the
// StreamClientStream stream.
type wrappedStreamClientStream struct {
	ctx             context.Context
	sendWithContext func(context.Context, *StreamStreamingPayload) error
	recvWithContext func(context.Context) (*StreamResult, error)
	stream          StreamClientStream
}

// wrapCacheGet applies the Cache server interceptor to endpoints.
func wrapGetCache(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &CacheInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.Cache(ctx, info, endpoint)
	}
}

// wrapJWTAuthGet applies the JWTAuth server interceptor to endpoints.
func wrapGetJWTAuth(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &JWTAuthInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.JWTAuth(ctx, info, endpoint)
	}
}

// wrapJWTAuthCreate applies the JWTAuth server interceptor to endpoints.
func wrapCreateJWTAuth(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &JWTAuthInfo{
			service:    "interceptors",
			method:     "Create",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.JWTAuth(ctx, info, endpoint)
	}
}

// wrapJWTAuthStream applies the JWTAuth server interceptor to endpoints.
func wrapStreamJWTAuth(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &JWTAuthInfo{
			service:    "interceptors",
			method:     "Stream",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.JWTAuth(ctx, info, endpoint)
	}
}

// wrapRequestAuditGet applies the RequestAudit server interceptor to endpoints.
func wrapGetRequestAudit(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &RequestAuditInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.RequestAudit(ctx, info, endpoint)
	}
}

// wrapRequestAuditCreate applies the RequestAudit server interceptor to
// endpoints.
func wrapCreateRequestAudit(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &RequestAuditInfo{
			service:    "interceptors",
			method:     "Create",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.RequestAudit(ctx, info, endpoint)
	}
}

// wrapSetDeadlineGet applies the SetDeadline server interceptor to endpoints.
func wrapGetSetDeadline(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &SetDeadlineInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.SetDeadline(ctx, info, endpoint)
	}
}

// wrapSetDeadlineCreate applies the SetDeadline server interceptor to
// endpoints.
func wrapCreateSetDeadline(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &SetDeadlineInfo{
			service:    "interceptors",
			method:     "Create",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.SetDeadline(ctx, info, endpoint)
	}
}

// wrapSetDeadlineStream applies the SetDeadline server interceptor to
// endpoints.
func wrapStreamSetDeadline(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &SetDeadlineInfo{
			service:    "interceptors",
			method:     "Stream",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.SetDeadline(ctx, info, endpoint)
	}
}

// wrapTraceBidirectionalStreamStream applies the TraceBidirectionalStream
// server interceptor to endpoints.
func wrapStreamTraceBidirectionalStream(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := endpoint(ctx, req)
		if err != nil {
			return res, err
		}
		stream := res.(StreamServerStream)
		return &wrappedStreamServerStream{
			ctx: ctx,
			sendWithContext: func(ctx context.Context, req *StreamResult) error {
				info := &TraceBidirectionalStreamInfo{
					service:    "interceptors",
					method:     "Stream",
					callType:   goa.InterceptorStreamingSend,
					rawPayload: req,
				}
				_, err := i.TraceBidirectionalStream(ctx, info, func(ctx context.Context, req any) (any, error) {
					castReq, _ := req.(*StreamResult)
					return nil, stream.SendWithContext(ctx, castReq)
				})
				return err
			},
			recvWithContext: func(ctx context.Context) (*StreamStreamingPayload, error) {
				info := &TraceBidirectionalStreamInfo{
					service:  "interceptors",
					method:   "Stream",
					callType: goa.InterceptorStreamingRecv,
				}
				res, err := i.TraceBidirectionalStream(ctx, info, func(ctx context.Context, _ any) (any, error) {
					return stream.RecvWithContext(ctx)
				})
				castRes, _ := res.(*StreamStreamingPayload)
				return castRes, err
			},
			stream: stream,
		}, nil
	}
}

// wrapTraceRequestGet applies the TraceRequest server interceptor to endpoints.
func wrapGetTraceRequest(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &TraceRequestInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.TraceRequest(ctx, info, endpoint)
	}
}

// wrapTraceRequestCreate applies the TraceRequest server interceptor to
// endpoints.
func wrapCreateTraceRequest(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &TraceRequestInfo{
			service:    "interceptors",
			method:     "Create",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.TraceRequest(ctx, info, endpoint)
	}
}

// wrapClientEncodeTenantGet applies the EncodeTenant client interceptor to
// endpoints.
func wrapClientGetEncodeTenant(endpoint goa.Endpoint, i ClientInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &EncodeTenantInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.EncodeTenant(ctx, info, endpoint)
	}
}

// wrapClientEncodeTenantCreate applies the EncodeTenant client interceptor to
// endpoints.
func wrapClientCreateEncodeTenant(endpoint goa.Endpoint, i ClientInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &EncodeTenantInfo{
			service:    "interceptors",
			method:     "Create",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.EncodeTenant(ctx, info, endpoint)
	}
}

// wrapClientEncodeTenantStream applies the EncodeTenant client interceptor to
// endpoints.
func wrapClientStreamEncodeTenant(endpoint goa.Endpoint, i ClientInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &EncodeTenantInfo{
			service:    "interceptors",
			method:     "Stream",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.EncodeTenant(ctx, info, endpoint)
	}
}

// wrapClientRetryGet applies the Retry client interceptor to endpoints.
func wrapClientGetRetry(endpoint goa.Endpoint, i ClientInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &RetryInfo{
			service:    "interceptors",
			method:     "Get",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.Retry(ctx, info, endpoint)
	}
}

// wrapClientRetryCreate applies the Retry client interceptor to endpoints.
func wrapClientCreateRetry(endpoint goa.Endpoint, i ClientInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		info := &RetryInfo{
			service:    "interceptors",
			method:     "Create",
			callType:   goa.InterceptorUnary,
			rawPayload: req,
		}
		return i.Retry(ctx, info, endpoint)
	}
}

// wrapClientTraceBidirectionalStreamStream applies the
// TraceBidirectionalStream client interceptor to endpoints.
func wrapClientStreamTraceBidirectionalStream(endpoint goa.Endpoint, i ClientInterceptors) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := endpoint(ctx, req)
		if err != nil {
			return res, err
		}
		stream := res.(StreamClientStream)
		return &wrappedStreamClientStream{
			ctx: ctx,
			sendWithContext: func(ctx context.Context, req *StreamStreamingPayload) error {
				info := &TraceBidirectionalStreamInfo{
					service:    "interceptors",
					method:     "Stream",
					callType:   goa.InterceptorStreamingSend,
					rawPayload: req,
				}
				_, err := i.TraceBidirectionalStream(ctx, info, func(ctx context.Context, req any) (any, error) {
					castReq, _ := req.(*StreamStreamingPayload)
					return nil, stream.SendWithContext(ctx, castReq)
				})
				return err
			},
			recvWithContext: func(ctx context.Context) (*StreamResult, error) {
				info := &TraceBidirectionalStreamInfo{
					service:  "interceptors",
					method:   "Stream",
					callType: goa.InterceptorStreamingRecv,
				}
				res, err := i.TraceBidirectionalStream(ctx, info, func(ctx context.Context, _ any) (any, error) {
					return stream.RecvWithContext(ctx)
				})
				castRes, _ := res.(*StreamResult)
				return castRes, err
			},
			stream: stream,
		}, nil
	}
}

// Send streams instances of "StreamServerStream" after executing the applied
// interceptor.
func (w *wrappedStreamServerStream) Send(v *StreamResult) error {
	return w.SendWithContext(w.ctx, v)
}

// SendWithContext streams instances of "StreamServerStream" after executing
// the applied interceptor with context.
func (w *wrappedStreamServerStream) SendWithContext(ctx context.Context, v *StreamResult) error {
	if w.sendWithContext == nil {
		return w.stream.SendWithContext(ctx, v)
	}
	return w.sendWithContext(ctx, v)
}

// Recv reads instances of "StreamServerStream" from the stream after executing
// the applied interceptor.
func (w *wrappedStreamServerStream) Recv() (*StreamStreamingPayload, error) {
	return w.RecvWithContext(w.ctx)
}

// RecvWithContext reads instances of "StreamServerStream" from the stream
// after executing the applied interceptor with context.
func (w *wrappedStreamServerStream) RecvWithContext(ctx context.Context) (*StreamStreamingPayload, error) {
	if w.recvWithContext == nil {
		return w.stream.RecvWithContext(ctx)
	}
	return w.recvWithContext(ctx)
}

// Close closes the stream.
func (w *wrappedStreamServerStream) Close() error {
	return w.stream.Close()
}

// Send streams instances of "StreamClientStream" after executing the applied
// interceptor.
func (w *wrappedStreamClientStream) Send(v *StreamStreamingPayload) error {
	return w.SendWithContext(w.ctx, v)
}

// SendWithContext streams instances of "StreamClientStream" after executing
// the applied interceptor with context.
func (w *wrappedStreamClientStream) SendWithContext(ctx context.Context, v *StreamStreamingPayload) error {
	if w.sendWithContext == nil {
		return w.stream.SendWithContext(ctx, v)
	}
	return w.sendWithContext(ctx, v)
}

// Recv reads instances of "StreamClientStream" from the stream after executing
// the applied interceptor.
func (w *wrappedStreamClientStream) Recv() (*StreamResult, error) {
	return w.RecvWithContext(w.ctx)
}

// RecvWithContext reads instances of "StreamClientStream" from the stream
// after executing the applied interceptor with context.
func (w *wrappedStreamClientStream) RecvWithContext(ctx context.Context) (*StreamResult, error) {
	if w.recvWithContext == nil {
		return w.stream.RecvWithContext(ctx)
	}
	return w.recvWithContext(ctx)
}

// Close closes the stream.
func (w *wrappedStreamClientStream) Close() error {
	return w.stream.Close()
}
