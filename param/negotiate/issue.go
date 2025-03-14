package negotiate

import (
	"encoding/json"

	"github.com/vela-ssoc/ssoc-common-mb/profile"
	"github.com/vela-ssoc/ssoc-common-mba/ciphertext"
)

type Issue struct {
	Name     string               `json:"name"`     // broker 名字
	Passwd   []byte               `json:"passwd"`   // 通信加解密密钥
	Server   profile.BrokerServer `json:"server"`   // broker 服务监听配置
	Logger   profile.Logger       `json:"logger"`   // 日志配置
	Database profile.Database     `json:"database"` // 数据库配置
}

// String fmt.Stringer
func (iss Issue) String() string {
	dat, _ := json.MarshalIndent(iss, "", "    ")
	return string(dat)
}

// Encrypt 将 Issue 加密。
func (iss Issue) Encrypt() ([]byte, error) {
	return ciphertext.EncryptJSON(iss)
}

func (iss *Issue) Decrypt(enc []byte) error {
	return ciphertext.DecryptJSON(enc, iss)
}
