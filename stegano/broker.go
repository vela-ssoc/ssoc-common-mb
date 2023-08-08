package stegano

type BHide struct {
	ID      int64    `json:"id"`      // Broker 节点的 ID
	Secret  string   `json:"secret"`  // 密钥
	Semver  string   `json:"semver"`  // 版本号
	Servers []string `json:"servers"` // 中心端地址
}
