package alarm

import (
	"context"
	"fmt"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

type eventTask struct {
	qry   *query.Query
	unify *unifyAlert
	event *model.Event
}

func (et *eventTask) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// 先入库保存
	evt := et.event
	tbl := et.qry.Event
	if err := tbl.WithContext(ctx).
		Create(et.event); err != nil || !evt.SendAlert {
		return
	}

	// 发送告警
	key := evt.FromCode
	sub := et.unify.match.Event(ctx, key)
	if sub == nil || sub.Empty() {
		et.unify.log.Info(fmt.Sprintf("event 事件 %s 没有订阅者", key))
		return
	}

	// 发送告警
	if dongs := sub.Dong; len(dongs) != 0 {
		et.sendDong(ctx, dongs, key)
	}
	if devs := sub.Devops; len(devs) != 0 {
		et.sendDevops(ctx, devs)
	}
}

func (et *eventTask) sendDong(ctx context.Context, dongs []string, key string) {
	title, body := et.unify.store.EventDong(ctx, et.event, "")
	if err := et.unify.dong.Send(ctx, dongs, []string{key}, title, body); err != nil {
		et.unify.log.Warn(fmt.Sprintf("发送事件 %s 失败：%s", dongs, err))
	} else {
		et.unify.log.Info("发送事件成功")
	}
}

func (et *eventTask) sendDevops(ctx context.Context, devs []*model.Devops) {
	// title, body := st.rend.RiskDong(ctx, st.dat, st.dat)
	err := et.unify.dps.Send(ctx, "告警", "内容", devs)
	et.unify.log.Info(fmt.Sprintf("发送 devops 事件 %s 结果：%v", devs, err))
}
