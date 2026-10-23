package pterm

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Stdout struct {
	conn *websocket.Conn
}

func (s *Stdout) Write(p []byte) (int, error) {
	n := len(p)
	if err := s.write("", string(p)); err != nil {
		return 0, err
	}

	return n, nil
}

func (s *Stdout) WriteError(err error) error {
	return s.write("error", err.Error())
}

func (s *Stdout) WriteErrorf(msg string, args ...any) error {
	return s.write("error", fmt.Sprintf(msg, args...))
}

func (s *Stdout) WritePong() error {
	return s.write("pong", "")
}

func (s *Stdout) write(typ, data string) error {
	if typ == "" {
		typ = "stdout"
	}
	msg := &responseData{Type: typ, Data: data}

	return s.conn.WriteJSON(msg)
}
