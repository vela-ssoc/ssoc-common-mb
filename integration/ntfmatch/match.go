package ntfmatch

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/dal/model"
	"github.com/vela-ssoc/ssoc-common-mb/dal/query"
)

type Matcher interface {
	Event(ctx context.Context, evtCode string) *model.Subscriber
	Risk(ctx context.Context, rskCode string) *model.Subscriber
}

func NewMatch(qry *query.Query) Matcher {
	return &notifyMatch{
		qry: qry,
	}
}

type notifyMatch struct {
	qry *query.Query
}

func (nm *notifyMatch) Event(ctx context.Context, evtCode string) *model.Subscriber {
	sub := nm.slowLoad(ctx)
	return sub.Event(evtCode)
}

func (nm *notifyMatch) Risk(ctx context.Context, rskCode string) *model.Subscriber {
	sub := nm.slowLoad(ctx)
	return sub.Risk(rskCode)
}

func (nm *notifyMatch) slowLoad(ctx context.Context) model.Subscribers {
	tbl := nm.qry.Notifier
	dats, err := tbl.WithContext(ctx).Find()
	if err != nil {
		return model.Subscribers{}
	}
	subs := model.Notifiers(dats).Subscribers()

	return subs
}
