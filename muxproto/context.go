package muxproto

import "context"

func WithContext(parent context.Context, node Node) context.Context {
	return context.WithValue(parent, muxServerNodeKey, node)
}

func FromContext(ctx context.Context) Node {
	val := ctx.Value(muxServerNodeKey)
	node, _ := val.(Node)

	return node
}

var muxServerNodeKey = contextKey{name: "tunnel-http-context-key"}

type contextKey struct {
	name string
}
