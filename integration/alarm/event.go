package alarm

import (
	"context"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/storage"
)

type eventSendTask struct {
	evt   *model.Event
	sub   *model.Subscriber
	store storage.Storer
	cli   Client
}

func (tsk *eventSendTask) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dongs := tsk.sub.Dong

	title, body := tsk.store.EventDong(ctx, tsk.evt)
	tsk.cli.Dong(ctx, dongs, title, body)

	// TODO implement me
	panic("implement me")
}
