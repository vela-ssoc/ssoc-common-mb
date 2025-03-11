package alarm

import (
	"context"
	"fmt"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

type riskTask struct {
	qry   *query.Query
	unify *unifyAlert
	risk  *model.Risk
}

func (et *riskTask) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// 先入库保存
	rsk := et.risk
	tbl := et.qry.Risk
	if err := tbl.WithContext(ctx).
		Create(rsk); err != nil || !rsk.SendAlert {
		return
	}

	// 发送告警
	key := rsk.FromCode
	sub := et.unify.match.Risk(ctx, key)
	if sub == nil || sub.Empty() {
		et.unify.log.Info(fmt.Sprintf("risk 风险 %s 没有订阅者", key))
		return
	}

	// 发送告警
	if dongs := sub.Dong; len(dongs) != 0 {
		et.sendDong(ctx, dongs, key)
	}
}

func (et *riskTask) sendDong(ctx context.Context, dongs []string, key string) {
	title, body := et.unify.store.RiskDong(ctx, et.risk, et.risk.Template)
	if err := et.unify.dong.Send(ctx, dongs, []string{key}, title, body); err != nil {
		et.unify.log.Warn(fmt.Sprintf("发送风险 %s 失败：%s", dongs, err))
	} else {
		et.unify.log.Info(fmt.Sprintf("发送风险 %s 成功", dongs))
	}
}
