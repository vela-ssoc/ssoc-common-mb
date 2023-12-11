package main

import (
	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		Mode:    gen.WithDefaultQuery,
		OutPath: "./dal/query",
	})

	tbl := tables()
	g.ApplyBasic(tbl...)

	g.Execute()
}

func tables() []any {
	return []any{
		model.Alert{},
		model.AuthTemp{},
		model.Broker{},
		model.BrokerBin{},
		model.BrokerStat{},
		model.Certificate{},
		model.Cmdb{},
		model.Ding{},
		model.Domain{},
		model.Dong{},
		model.Effect{},
		model.Elastic{},
		model.Email{},
		model.Emc{},
		model.Event{},
		model.Job{},
		model.JobCode{},
		model.JobPolicy{},
		model.JobReport{},
		model.LoginLock{},
		model.LoginRetry{},
		model.Minion{},
		model.MinionAccount{},
		model.MinionBin{},
		model.MinionCustomized{},
		model.MinionGroup{},
		model.MinionListen{},
		model.MinionLogon{},
		model.MinionProcess{},
		model.MinionTag{},
		model.MinionTask{},
		model.Notifier{},
		model.Oplog{},
		model.PassDNS{},
		model.PassIP{},
		model.Plate{},
		model.Purl{},
		model.Recipient{},
		model.Resignation{},
		model.Risk{},
		model.RiskDNS{},
		model.RiskFile{},
		model.RiskIP{},
		model.Rotate{},
		model.SBOMComponent{},
		model.SBOMMinion{},
		model.SBOMProject{},
		model.SBOMVuln{},
		model.Startup{},
		model.Store{},
		model.Substance{},
		model.SubstanceTask{},
		model.SysInfo{},
		model.Third{},
		model.ThirdCustomized{},
		model.User{},
		model.VIP{},
	}
}
