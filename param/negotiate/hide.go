package negotiate

import "github.com/vela-ssoc/vela-common-mba/netutil"

// Hide broker 二进制隐写数据。
type Hide struct {
	Secret  string            `json:"secret"`  // 密钥
	Servers netutil.Addresses `json:"servers"` // 中心端地址
}
