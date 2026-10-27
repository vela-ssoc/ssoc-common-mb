package muxproto

import "net/http"

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

type BrokerAuthResponse[T any] struct {
	Code   int    `json:"code"`
	Text   string `json:"text"`
	Config T      `json:"config"`
}

func NewAcceptedBrokerAuthResponse[T any](c T) BrokerAuthResponse[T] {
	return BrokerAuthResponse[T]{Code: http.StatusAccepted, Config: c}
}

func (r BrokerAuthResponse[T]) Err() error {
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
