package muxproto

import "context"

func WithContext(parent context.Context, node Node) context.Context {
	return context.WithValue(parent, httpContextKey, node)
}

func FromContext(ctx context.Context) Node {
	val := ctx.Value(httpContextKey)
	node, _ := val.(Node)

	return node
}

var httpContextKey = contextKey{name: "tunnel-http-context-key"}

type contextKey struct {
	name string
}
