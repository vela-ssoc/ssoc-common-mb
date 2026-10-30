package crontab

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/robfig/cron/v3"
)

type Jober interface {
	Options() JobOptions
	Execute(context.Context) error
}

type JobOptions struct {
	ID        string        `json:"id"`        // 任务 ID，要求全局唯一。
	Name      string        `json:"name"`      // 任务名，用于辅助用户识别。
	Expr      string        `json:"expr"`      // 表达式（当 Sched 不存在时生效）
	Sched     cron.Schedule `json:"-"`         // 定时器（比 Expr 优先级高）
	Timeout   time.Duration `json:"timeout"`   // 任务触发后执行超时时间
	Immediate bool          `json:"immediate"` // 立即执行
}

type JobInfo struct {
	JobOptions
	Next    time.Time    `json:"next,omitzero"`
	EntryID cron.EntryID `json:"entry_id"`
}

type safeJob struct {
	job Jober
	opt JobOptions
	run atomic.Bool
	log *slog.Logger
}

func (sj *safeJob) Run() {
	args := []any{"options", sj.opt}
	sj.log.Debug("开始执行定时任务", args...)
	if exec, err := sj.exec(context.Background()); !exec {
		sj.log.Warn("任务已经正在运行中（本次触发未执行）", args...)
	} else if err != nil {
		args = append(args, "err", err)
		sj.log.Info("任务执行出错", args...)
	} else {
		sj.log.Debug("任务执行完毕", args...)
	}
}

func (sj *safeJob) exec(parent context.Context) (bool, error) {
	if !sj.run.CompareAndSwap(false, true) {
		return false, nil
	}
	defer sj.run.Store(false)

	err := sj.safeExec(parent)

	return true, err
}

func (sj *safeJob) safeExec(ctx context.Context) (err error) {
	if du := sj.opt.Timeout; du > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, du)
		defer cancel()
	}

	defer func() {
		if v := recover(); v != nil {
			stack := debug.Stack()
			err = fmt.Errorf("panicked: %v\n%s", v, stack)
			_, _ = fmt.Fprintf(os.Stderr, "%s\n%s", v, stack)
		}
	}()

	err = sj.job.Execute(ctx)

	return
}
