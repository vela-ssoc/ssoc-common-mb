package pterm

import "github.com/gorilla/websocket"

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

func (s *Stdout) WriteType(typ, data string) error {
	return s.write(typ, data)
}

func (s *Stdout) write(typ, data string) error {
	if typ == "" {
		typ = "stdout"
	}
	msg := &responseData{Type: typ, Data: data}

	return s.conn.WriteJSON(msg)
}
