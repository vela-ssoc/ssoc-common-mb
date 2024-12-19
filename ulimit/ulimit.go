//go:build !windows

package ulimit

import "syscall"

// Least 调整文件描述符限制，如果当前 fd 限制小于 n 个，
// 则将限制调整为 n 个。
func Least(n uint64) error {
	lim := new(syscall.Rlimit)
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, lim); err != nil {
		return err
	}
	if lim.Cur >= n && lim.Max >= n {
		return nil
	}

	if lim.Cur < n {
		lim.Cur = n
	}
	if lim.Max < n {
		lim.Max = n
	}

	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, lim)
}
