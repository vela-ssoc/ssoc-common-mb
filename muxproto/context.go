package muxproto

import "context"

func WithContext(parent context.Context, p Peer) context.Context {
	return context.WithValue(parent, httpContextKey, p)
}

func FromContext(ctx context.Context) Peer {
	val := ctx.Value(httpContextKey)
	p, _ := val.(Peer)

	return p
}

var httpContextKey = contextKey{name: "tunnel-http-context-key"}

type contextKey struct {
	name string
}
