package profile

type ManagerServer struct {
	Addr    string   `json:"addr"`                             // 监听地址
	Cert    string   `json:"cert"    validate:"lte=500"`       // 证书路径
	Pkey    string   `json:"pkey"    validate:"lte=500"`       // 私钥路径
	Static  string   `json:"static"  validate:"lte=500"`       // 静态资源路径
	Vhosts  []string `json:"vhosts"  validate:"dive,required"` // 虚拟主机头
	Session duration `json:"session"`                          // 用户 session 有效期
	CDN     string   `json:"cdn"`                              // 数据库文件缓存目录
}

type BrokerServer struct {
	Addr string `json:"addr"`                     // 监听地址
	Cert string `json:"cert" validate:"lte=5000"` // 证书内容
	Pkey string `json:"pkey" validate:"lte=5000"` // 私钥内容
	CDN  string `json:"cdn"`                      // 数据库文件缓存目录
}
