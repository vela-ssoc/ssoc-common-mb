package entity

import (
	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	tbls := List()
	return db.AutoMigrate(tbls...)
}

func List() []any {
	return []any{
		model.AlertServer{},
		model.AuthTemp{},
		model.Broker{},
		model.BrokerBin{},
		model.BrokerStat{},
		model.Certificate{},
		model.Cmdb{},
		model.Cmdb2{},
		model.Domain{},
		model.Dong{},
		model.EagleEyeData{},
		model.Effect{},
		model.Elastic{},
		model.Email{},
		model.Emc{},
		model.Event{},
		model.GridChunk{},
		model.GridFile{},
		model.Job{},
		model.JobCode{},
		model.JobPolicy{},
		model.JobReport{},
		model.KVAudit{},
		model.KVData{},
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
		model.SBOMComponent{},
		model.SBOMMinion{},
		model.SBOMProject{},
		model.SBOMVuln{},
		model.SIEMServer{},
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
