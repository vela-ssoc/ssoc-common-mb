package negotiate

import "github.com/vela-ssoc/vela-common-mba/netutil"

// Hide broker 二进制隐写数据。
type Hide struct {
	ID      int64             `json:"id"`      // Broker 节点的 ID
	Secret  string            `json:"secret"`  // 密钥
	Semver  string            `json:"semver"`  // 版本号
	Servers netutil.Addresses `json:"servers"` // 中心端地址
}
