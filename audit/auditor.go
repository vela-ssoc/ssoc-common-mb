package audit

import (
	"context"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/logback"
)

type Auditor interface {
	Event(context.Context, *model.Event) error
	Risk(context.Context, *model.Risk) error
}

func NewAuditor(slog logback.Logger) Auditor {
	nano := time.Now().UnixNano()
	random := rand.New(rand.NewSource(nano))

	return &alarmAudit{
		slog:   slog,
		random: random,
	}
}

type alarmAudit struct {
	slog   logback.Logger
	random *rand.Rand
}

func (ala *alarmAudit) Event(ctx context.Context, evt *model.Event) error {
	if evt == nil || evt.MinionID == 0 || evt.Inet == "" {
		ala.slog.Warnf("event 缺少 minionID 或 inet，丢弃该 event: %#v", evt)
		return nil
	}
	evt.ID = 0
	evt.HaveRead = false
	now := time.Now()
	evt.CreatedAt = now
	if evt.OccurAt.IsZero() {
		evt.OccurAt = now
	}
	alert := evt.SendAlert
	if alert {
		buf := make([]byte, 16)
		ala.random.Read(buf)
		evt.Secret = hex.EncodeToString(buf)
	}

	if err := query.Event.WithContext(ctx).Create(evt); err != nil || !alert {
		return err
	}

	// TODO 发送告警

	return nil
}

func (ala *alarmAudit) Risk(ctx context.Context, rsk *model.Risk) error {
	if rsk.MinionID == 0 || rsk.Inet == "" {
		ala.slog.Warn("risk 缺少 minionID 或 inet，丢弃该 risk")
		return nil
	}

	rsk.Status = model.RSUnprocessed
	now := time.Now()
	if rsk.OccurAt.IsZero() {
		rsk.OccurAt = now
	}

	alert := rsk.SendAlert
	if alert {
		buf := make([]byte, 16)
		ala.random.Read(buf)
		rsk.Secret = hex.EncodeToString(buf)
	}
	if err := query.Risk.WithContext(ctx).Create(rsk); err != nil || !alert {
		return err
	}

	// TODO 发送告警

	return nil
}
