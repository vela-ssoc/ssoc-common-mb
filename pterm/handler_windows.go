//go:build windows

package pterm

import (
	"context"
	"errors"
	"io"
	"os/exec"

	"github.com/UserExistsError/conpty"
	"github.com/gorilla/websocket"
)

func HandlePTY(ws *websocket.Conn, cmd string) error {
	defer ws.Close()
	if !conpty.IsConPtyAvailable() {
		return errors.ErrUnsupported
	}
	if cmd == "" {
		for _, str := range []string{
			"pwsh.exe",
			"C:\\Program Files\\PowerShell\\7\\pwsh.exe",
			"powershell.exe",
			"cmd.exe"} {
			if val, err := exec.LookPath(str); err == nil && val != "" {
				cmd = val
				break
			}
		}
		if cmd == "" {
			cmd = "C:\\WINDOWS\\system32\\cmd.exe"
		}
	}

	stdin, stdout := NewIO(ws)
	cpty, err := conpty.Start(cmd)
	if err != nil {
		_ = stdout.WriteErrorf("[%s]: %s", cmd, err.Error())
		return err
	}
	defer cpty.Close()

	go func() {
		_, err1 := io.Copy(stdout, cpty)
		_ = stdout.WriteError(err1)
		_ = ws.Close()
	}()
	go func() {
		// 监听程序是否退出。
		if _, err1 := cpty.Wait(context.Background()); err1 != nil {
			_ = stdout.WriteError(err1)
		}
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
				_, err2 = cpty.Write([]byte(msg))
			}
			if err2 != nil {
				_ = stdout.WriteError(err2)
				return err2
			}
		case "resize":
			msg, err2 := data.AsResize()
			if err2 == nil {
				err2 = cpty.Resize(msg.Cols, msg.Rows)
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
