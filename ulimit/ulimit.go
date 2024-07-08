package ulimit

import "syscall"

// Set 在 go <= 1.18 中，会自动将 fd 限制为 4096 个，
// 该方式用于修改默认限制。
func Set(n uint64) error {
	lim := &syscall.Rlimit{Cur: n, Max: n}
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, lim)
}
