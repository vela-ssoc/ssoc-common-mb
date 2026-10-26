package bininfo

import (
	"fmt"
	"io"
	"runtime/debug"
	"time"
)

type Info struct {
	GOOS        string           `json:"goos"`                  // 操作系统
	GOARCH      string           `json:"goarch"`                // 架构
	Version     string           `json:"version"`               // 版本号
	Revision    string           `json:"revision"`              // 修订版本
	Username    string           `json:"username"`              // 用户名
	Workdir     string           `json:"workdir"`               // 工作目录
	Module      string           `json:"module"`                // 包路径
	CommittedAt time.Time        `json:"committed_at,omitzero"` // VCS 提交时间
	BuildInfo   *debug.BuildInfo `json:"build_info,omitzero"`   // BuildInfo
}

func (inf Info) ANSI(w io.Writer, tpl string) (int, error) {
	return fmt.Fprintf(w, tpl, inf.Version, inf.Module, inf.GOOS,
		inf.GOARCH, inf.Workdir, inf.CommittedAt, inf.Module, inf.Revision)
}
