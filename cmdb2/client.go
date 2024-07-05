package cmdb2

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/vela-ssoc/vela-common-mba/netutil"
)

var (
	ciTypeServers = []string{"server", "vserver", "vmware_vserver", "public_cloud", "docker"}
	ciTypeVIPs    = []string{"vip"}
)

type Configurer interface {
	Configure(ctx context.Context) (*Config, error)
}

type Config struct {
	URL       string
	AccessKey string
	SecretKey string
}

func (c Config) appendURL(path string) string {
	return c.URL + path
}

func NewClient(cfg Configurer, cli netutil.HTTPClient) Client {
	return Client{
		cfg: cfg,
		cli: cli,
	}
}

type Client struct {
	cfg Configurer
	cli netutil.HTTPClient
}

func (c Client) Servers(ctx context.Context, ips []string, page, size int) ([]*Server, error) {
	if size <= 0 {
		size = 30
	}

	ret := make([]*Server, 0, size)
	if err := c.getCITypes(ctx, ciTypeServers, ips, page, size, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c Client) VIPs(ctx context.Context, page, size int) ([]*VIP, error) {
	if size <= 0 {
		size = 30
	}

	ret := make([]*VIP, 0, size)
	if err := c.getCITypes(ctx, ciTypeVIPs, nil, page, size, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c Client) getCITypes(ctx context.Context, ciTypes, ips []string, page, size int, ret any) error {
	types := strings.Join(ciTypes, ";")
	args := "_type:(" + types + ")"
	if len(ips) != 0 {
		inets := strings.Join(ips, ";")
		args = args + ",private_ip:(" + inets + ")"
	}
	quires := url.Values{"q": []string{args}}

	if page > 0 {
		quires.Set("page", strconv.Itoa(page))
	}
	if size > 0 {
		quires.Set("count", strconv.Itoa(size))
	}

	return c.getCI(ctx, quires, ret)
}

func (c Client) getCI(ctx context.Context, quires url.Values, ret any) error {
	cfg, err := c.cfg.Configure(ctx)
	if err != nil {
		return err
	}
	quires.Set("_key", cfg.AccessKey)
	quires.Set("_secret", cfg.SecretKey)
	addr := cfg.appendURL("/cmdb2/api/v0.1/ci/s?" + quires.Encode())

	return c.cli.JSON(ctx, http.MethodGet, addr, nil, &result{Records: ret}, nil)
}
