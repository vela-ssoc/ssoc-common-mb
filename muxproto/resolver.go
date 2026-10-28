package muxproto

import (
	"context"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

type TLDResolver interface {
	// TLD 获取顶级域名，例如 manager.internal
	TLD() string

	// MUX 通过主机位获取虚拟通道，不存在则返回 nil。
	// 例如完整域名是 a.manager.internal
	// TLD 是：manager.internal 主机位是：a
	// Dialer 会通 TLD 比较是否需要调用此 Resolver，
	// 如果命中 Resolver 则会提取主机位来获取 MUX 虚拟通道。
	MUX(ctx context.Context, host string) (muxconn.Muxer, error)
}

type MUXResolver struct {
	tld string
	mux muxconn.Muxer
}

func NewMUXResolver(tld string, mux muxconn.Muxer) *MUXResolver {
	return &MUXResolver{tld: tld, mux: mux}
}

func (r *MUXResolver) TLD() string                                            { return r.tld }
func (r *MUXResolver) MUX(_ context.Context, _ string) (muxconn.Muxer, error) { return r.mux, nil }

type HubResolver struct {
	hub Huber
}

func NewHubResolver(hub Huber) *HubResolver {
	return &HubResolver{hub: hub}
}

func (r *HubResolver) TLD() string { return r.hub.TLD() }

func (r *HubResolver) MUX(_ context.Context, host string) (muxconn.Muxer, error) {
	if node := r.hub.Get(host); node != nil {
		return node.MUX(), nil
	}

	return nil, nil
}
