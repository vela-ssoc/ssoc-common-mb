//go:build !windows

package pterm

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

//goland:noinspection GoUnhandledErrorResult
func HandlePTY(ws *websocket.Conn, cmd string) error {
	defer ws.Close()

	if cmd == "" {
		envsh := os.Getenv("SHELL")
		cmds := []string{envsh, "fish", "zsh", "bash", "sh"}
		cmd = LookPath(cmds, "/bin/bash")
	}

	stdin, stdout := NewIO(ws)
	c := exec.Command(cmd)
	ptmx, err := pty.Start(c)
	if err != nil {
		_ = stdout.WriteErrorf("[%s]: %s", cmd, err.Error())
		return err
	}
	defer func() {
		_ = ptmx.Close()
		if proc := c.Process; proc != nil {
			_ = proc.Kill()
		}
	}()

	go func() {
		_, err1 := io.Copy(stdout, ptmx)
		_ = stdout.WriteError(err1)
		_ = ws.Close()
	}()

	var exit bool
	for !exit {
		data, err1 := stdin.Read()
		if err1 != nil {
			_ = stdout.WriteError(err1)
			return err1
		}

		switch data.Type {
		case "ping":
			_ = stdout.WritePong()
			continue
		case "exit":
			exit = true
		case "stdin":
			msg, err2 := data.AsString()
			if err2 == nil {
				_, err2 = ptmx.Write([]byte(msg))
			}
			if err2 != nil {
				_ = stdout.WriteError(err2)
				return err2
			}
		case "resize":
			msg, err2 := data.AsResize()
			if err2 == nil {
				winsz := &pty.Winsize{Rows: uint16(msg.Rows), Cols: uint16(msg.Cols)}
				err2 = pty.Setsize(ptmx, winsz)
			}
			if err2 != nil {
				_ = stdout.WriteError(err2)
				return err2
			}
		default:
			_ = stdout.WriteErrorf("接收收到无效的数据")
			exit = true
		}
	}

	return nil
}
