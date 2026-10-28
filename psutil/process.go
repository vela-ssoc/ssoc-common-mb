package psutil

import (
	"context"
	"time"

	psprocess "github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	PID        int32                     `json:"pid"`
	Name       string                    `json:"name"`
	Cwd        string                    `json:"cwd"`
	Exe        string                    `json:"exe"`
	Cmdline    string                    `json:"cmdline"`
	Username   string                    `json:"username"`
	CPUPercent float64                   `json:"cpu_percent"`
	NumThreads int32                     `json:"num_threads"`
	NumFDs     int32                     `json:"num_fds"`
	CreatedAt  time.Time                 `json:"created_at"`
	Memory     *psprocess.MemoryInfoStat `json:"memory"`
}

func ProcessesWithContext(ctx context.Context) ([]*ProcessInfo, error) {
	procs, err := psprocess.ProcessesWithContext(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]*ProcessInfo, 0, len(procs))
	for _, proc := range procs {
		info, err1 := ProcessInfoLookup(ctx, proc)
		if err1 != nil {
			continue
		}
		ret = append(ret, info)
	}

	return ret, nil
}

func ProcessInfoLookup(ctx context.Context, proc *psprocess.Process) (*ProcessInfo, error) {
	info := new(ProcessInfo)
	info.PID = proc.Pid
	name, err := proc.Name()
	if err != nil { // 可能进程已销毁
		return nil, err
	}
	info.Name = name
	info.Cwd, _ = proc.CwdWithContext(ctx)
	info.Exe, _ = proc.ExeWithContext(ctx)
	info.Cmdline, _ = proc.CmdlineWithContext(ctx)
	info.Username, _ = proc.UsernameWithContext(ctx)
	info.CPUPercent, _ = proc.CPUPercentWithContext(ctx)
	info.NumThreads, _ = proc.NumThreadsWithContext(ctx)
	info.NumFDs, _ = proc.NumFDsWithContext(ctx)
	info.Memory, _ = proc.MemoryInfoWithContext(ctx)
	if mill, err1 := proc.CreateTimeWithContext(ctx); err1 == nil {
		info.CreatedAt = time.UnixMilli(mill)
	}

	return info, nil
}
