package muxproto

type BrokerAuthRequest struct {
	Secret     string   `json:"secret"   validate:"required"` // broker 密钥
	Version    string   `json:"version"  validate:"required"` // broker 版本号
	Inet       string   `json:"inet"     validate:"required"`
	Goos       string   `json:"goos"     validate:"required"`
	Goarch     string   `json:"goarch"   validate:"required"`
	PID        int      `json:"pid"`
	Hostname   string   `json:"hostname"`
	Workdir    string   `json:"workdir"`
	Executable string   `json:"executable"`
	Args       []string `json:"args"`
}

type BrokerAuthResponse struct {
	Code   int    `json:"code"`
	Text   string `json:"text"`
	Config any    `json:"config"`
}

func (r BrokerAuthResponse) Err() error {
	if r.Code/100 == 2 {
		return nil
	}

	return &AuthError{Code: r.Code, Text: r.Text}
}

type AuthError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (a *AuthError) Error() string {
	return a.Text
}
