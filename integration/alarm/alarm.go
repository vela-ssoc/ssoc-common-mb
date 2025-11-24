package alarm

import (
	"context"
	"encoding/hex"
	"log/slog"
	"math/rand"
	"time"

	"github.com/vela-ssoc/ssoc-common-mb/dal/model"
	"github.com/vela-ssoc/ssoc-common-mb/dal/query"
	"github.com/vela-ssoc/ssoc-common-mb/gopool"
	"github.com/vela-ssoc/ssoc-common-mb/integration/devops"
	"github.com/vela-ssoc/ssoc-common-mb/integration/dong/v2"
	"github.com/vela-ssoc/ssoc-common-mb/integration/ntfmatch"
	"github.com/vela-ssoc/ssoc-common-mb/storage/v2"
)

type Alerter interface {
	EventSaveAndAlert(ctx context.Context, evt *model.Event) error
	RiskSaveAndAlert(ctx context.Context, rsk *model.Risk) error
}

func UnifyAlerter(store storage.Storer,
	match ntfmatch.Matcher,
	log *slog.Logger,
	dong dong.Client,
	dps devops.Client,
	qry *query.Query,
) Alerter {
	nano := time.Now().UnixNano()
	random := rand.New(rand.NewSource(nano))

	return &unifyAlert{
		store:  store,
		pool:   gopool.New(1024),
		match:  match,
		log:    log,
		dong:   dong,
		dps:    dps,
		random: random,
		qry:    qry,
	}
}

type unifyAlert struct {
	store  storage.Storer
	pool   gopool.Pool
	match  ntfmatch.Matcher
	log    *slog.Logger
	dong   dong.Client
	dps    devops.Client
	random *rand.Rand
	qry    *query.Query
}

func (ua *unifyAlert) EventSaveAndAlert(ctx context.Context, evt *model.Event) error {
	if evt == nil {
		ua.log.Warn("event 为 nil 不作处理")
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

	task := &eventTask{
		unify: ua,
		event: evt,
		qry:   ua.qry,
	}
	ua.pool.Go(task.Run)

	return nil
}

func (ua *unifyAlert) RiskSaveAndAlert(ctx context.Context, rsk *model.Risk) error {
	if rsk == nil {
		ua.log.Warn("risk 为 nil 不作处理")
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

	task := &riskTask{
		qry:   ua.qry,
		unify: ua,
		risk:  rsk,
	}
	ua.pool.Go(task.Run)

	return nil
}
