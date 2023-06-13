package alarm

import (
	"context"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/integration/devops"
	"github.com/vela-ssoc/vela-common-mb/integration/dong"
	"github.com/vela-ssoc/vela-common-mb/integration/formwork"
	"github.com/vela-ssoc/vela-common-mb/integration/ntfmatch"
	"github.com/vela-ssoc/vela-common-mb/logback"
	"github.com/vela-ssoc/vela-common-mb/taskpool"
)

type Alerter interface {
	EventSaveAndAlert(ctx context.Context, evt *model.Event) error
	RiskSaveAndAlert(ctx context.Context, rsk *model.Risk) error
}

func UnifyAlerter(rend formwork.Render,
	pool taskpool.Executor,
	match ntfmatch.Matcher,
	slog logback.Logger,
	dong dong.Client,
	dps devops.Client,
) Alerter {
	nano := time.Now().UnixNano()
	random := rand.New(rand.NewSource(nano))

	return &unifyAlert{
		rend:   rend,
		pool:   pool,
		match:  match,
		slog:   slog,
		dong:   dong,
		dps:    dps,
		random: random,
	}
}

type unifyAlert struct {
	rend   formwork.Render
	pool   taskpool.Executor
	match  ntfmatch.Matcher
	slog   logback.Logger
	dong   dong.Client
	dps    devops.Client
	random *rand.Rand
}

func (ua *unifyAlert) EventSaveAndAlert(ctx context.Context, evt *model.Event) error {
	if evt == nil {
		ua.slog.Warn("event 为 nil 不作处理")
		return nil
	}

	// 入库前处理
	now := time.Now()
	if evt.OccurAt.IsZero() {
		evt.OccurAt = now
	}
	if evt.SendAlert {
		buf := make([]byte, 16)
		ua.random.Read(buf)
		evt.Secret = hex.EncodeToString(buf)
	}

	// 入库保存
	if err := query.Event.WithContext(ctx).Create(evt); err != nil || !evt.SendAlert {
		return err
	}

	key := evt.FromCode
	sub := ua.match.Event(ctx, key)
	if sub == nil || sub.Empty() {
		ua.slog.Infof("event 事件 %s 没有订阅者", key)
		return nil
	}

	st := &sendTask{
		sub:  sub,
		rend: ua.rend,
		slog: ua.slog,
		dong: ua.dong,
		dps:  ua.dps,
	}
	ua.pool.Submit(st)

	return nil
}

func (ua *unifyAlert) RiskSaveAndAlert(ctx context.Context, rsk *model.Risk) error {
	if rsk == nil {
		ua.slog.Warn("risk 为 nil 不作处理")
		return nil
	}

	// 入库前处理
	now := time.Now()
	rsk.Status = model.RSUnprocessed
	if rsk.OccurAt.IsZero() {
		rsk.OccurAt = now
	}
	if rsk.SendAlert {
		buf := make([]byte, 16)
		ua.random.Read(buf)
		rsk.Secret = hex.EncodeToString(buf)
	}
	// 入库保存
	if err := query.Risk.WithContext(ctx).Create(rsk); err != nil || !rsk.SendAlert {
		return err
	}

	key := rsk.FromCode
	sub := ua.match.Risk(ctx, key)
	if sub == nil || sub.Empty() {
		ua.slog.Infof("风险 %s 没有订阅者", key)
		return nil
	}

	st := &sendTask{
		sub:  sub,
		rend: ua.rend,
		slog: ua.slog,
		dong: ua.dong,
		dps:  ua.dps,
	}
	ua.pool.Submit(st)

	return nil
}
