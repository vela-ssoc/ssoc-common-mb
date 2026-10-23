//go:build !windows

package pterm

import (
	"errors"

	"github.com/gorilla/websocket"
)

func HandlePTY(ws *websocket.Conn, cmd string) error {
	return errors.ErrUnsupported
}
