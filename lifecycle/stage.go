package lifecycle

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/NVIDIA/gontainer/v2"
)

type StageFunc func(ctx context.Context) error

type Stage struct {
	Name    string
	OnStart StageFunc
	OnClose StageFunc
}

type Manager struct {
	per    time.Duration
	log    *slog.Logger
	stages []Stage
}

func NewManager(log *slog.Logger, per ...time.Duration) *Manager {
	var du time.Duration
	if len(per) != 0 {
		du = per[0]
	}

	return &Manager{per: du, log: log}
}

func (m *Manager) Append(ss gontainer.Multiple[Stage]) {
	m.stages = append(m.stages, ss...)
}

func (m *Manager) OnStart(ctx context.Context) error {
	for _, st := range m.stages {
		if err := m.call(ctx, st.OnClose); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) OnClose(ctx context.Context) error {
	if ctx.Err() != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Minute)
		defer cancel()
	}

	var errs []error
	for _, st := range m.stages {
		if err := m.call(ctx, st.OnClose); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (m *Manager) call(parent context.Context, f StageFunc) error {
	if f == nil {
		return nil
	}

	ctx, cancel := m.perContext(parent)
	defer cancel()

	return f(ctx)
}

func (m *Manager) perContext(parent context.Context) (context.Context, context.CancelFunc) {
	if m.per > 0 {
		return context.WithTimeout(parent, m.per)
	}

	return context.WithTimeout(parent, 10*time.Second)
}
