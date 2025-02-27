package profile

type ManagerConfig struct {
	Active   string        `json:"active"   yaml:"active"`   // 多环境配置
	Server   ManagerServer `json:"server"   yaml:"server"`   // HTTP 服务
	Database Database      `json:"database" yaml:"database"` // 数据库
	Logger   Logger        `json:"logger"   yaml:"logger"`   // 日志
	Cmdb2    Cmdb2         `json:"cmdb2"    yaml:"cmdb2"`    // CMDB2
	Oauth    Oauth         `json:"oauth"    yaml:"oauth"`
}

type BrokerConfig struct {
	Server   BrokerServer `json:"server"`   // Broker 服务配置
	Database Database     `json:"database"` // 数据库连接信息
	Logger   Logger       `json:"logger"`   // 全局日志
}

type Oauth struct {
	CAS          string `json:"cas"           yaml:"cas"` // 旧版认证接口
	URL          string `json:"url"           yaml:"url"`
	ClientID     string `json:"client_id"     yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	RedirectURL  string `json:"redirect_url"  yaml:"redirect_url"`
}

type Cmdb2 struct {
	URL       string `json:"url"        yaml:"url"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
}
