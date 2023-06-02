package alarm

import (
	"context"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
)

type Alerter interface {
	Event(ctx context.Context, evt *model.Event) error
	Risk(ctx context.Context, rsk *model.Risk) error
}

func UnifyAlerter() Alerter {
	return nil
}

type unifyAlert struct{}

func (ua *unifyAlert) Event(ctx context.Context, evt *model.Event) error {
	// TODO implement me
	panic("implement me")
}

func (ua *unifyAlert) Risk(ctx context.Context, rsk *model.Risk) error {
	// TODO implement me
	panic("implement me")
}
