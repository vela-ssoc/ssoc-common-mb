package pterm

import "github.com/gorilla/websocket"

type Stdin struct {
	conn *websocket.Conn
}

func (s *Stdin) Read() (*Input, error) {
	in := new(Input)
	if err := s.conn.ReadJSON(in); err != nil {
		return nil, err
	}

	return in, nil
}

func NewIO(conn *websocket.Conn) (*Stdin, *Stdout) {
	return &Stdin{conn: conn}, &Stdout{conn: conn}
}
