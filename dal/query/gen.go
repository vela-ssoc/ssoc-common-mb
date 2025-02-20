// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:               db,
		AlertServer:      newAlertServer(db, opts...),
		AuthTemp:         newAuthTemp(db, opts...),
		Broker:           newBroker(db, opts...),
		BrokerBin:        newBrokerBin(db, opts...),
		BrokerStat:       newBrokerStat(db, opts...),
		CASConfig:        newCASConfig(db, opts...),
		Certificate:      newCertificate(db, opts...),
		Cmdb:             newCmdb(db, opts...),
		Cmdb2:            newCmdb2(db, opts...),
		Domain:           newDomain(db, opts...),
		Dong:             newDong(db, opts...),
		EagleEyeData:     newEagleEyeData(db, opts...),
		Effect:           newEffect(db, opts...),
		Elastic:          newElastic(db, opts...),
		Email:            newEmail(db, opts...),
		Emc:              newEmc(db, opts...),
		Event:            newEvent(db, opts...),
		ExtensionMarket:  newExtensionMarket(db, opts...),
		ExtensionRecord:  newExtensionRecord(db, opts...),
		GridChunk:        newGridChunk(db, opts...),
		GridFile:         newGridFile(db, opts...),
		KVAudit:          newKVAudit(db, opts...),
		KVData:           newKVData(db, opts...),
		LoginLock:        newLoginLock(db, opts...),
		Minion:           newMinion(db, opts...),
		MinionAccount:    newMinionAccount(db, opts...),
		MinionBin:        newMinionBin(db, opts...),
		MinionCustomized: newMinionCustomized(db, opts...),
		MinionGroup:      newMinionGroup(db, opts...),
		MinionListen:     newMinionListen(db, opts...),
		MinionLogon:      newMinionLogon(db, opts...),
		MinionProcess:    newMinionProcess(db, opts...),
		MinionTag:        newMinionTag(db, opts...),
		MinionTask:       newMinionTask(db, opts...),
		Notifier:         newNotifier(db, opts...),
		Oplog:            newOplog(db, opts...),
		PassDNS:          newPassDNS(db, opts...),
		PassIP:           newPassIP(db, opts...),
		Plate:            newPlate(db, opts...),
		Purl:             newPurl(db, opts...),
		Recipient:        newRecipient(db, opts...),
		Resignation:      newResignation(db, opts...),
		Risk:             newRisk(db, opts...),
		RiskDNS:          newRiskDNS(db, opts...),
		RiskFile:         newRiskFile(db, opts...),
		RiskIP:           newRiskIP(db, opts...),
		SBOMComponent:    newSBOMComponent(db, opts...),
		SBOMMinion:       newSBOMMinion(db, opts...),
		SBOMProject:      newSBOMProject(db, opts...),
		SBOMVuln:         newSBOMVuln(db, opts...),
		SIEMServer:       newSIEMServer(db, opts...),
		Startup:          newStartup(db, opts...),
		Store:            newStore(db, opts...),
		Substance:        newSubstance(db, opts...),
		SubstanceTask:    newSubstanceTask(db, opts...),
		SysInfo:          newSysInfo(db, opts...),
		TaskExecute:      newTaskExecute(db, opts...),
		TaskExecuteItem:  newTaskExecuteItem(db, opts...),
		TaskExtension:    newTaskExtension(db, opts...),
		Third:            newThird(db, opts...),
		ThirdCustomized:  newThirdCustomized(db, opts...),
		User:             newUser(db, opts...),
		VIP:              newVIP(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AlertServer      alertServer
	AuthTemp         authTemp
	Broker           broker
	BrokerBin        brokerBin
	BrokerStat       brokerStat
	CASConfig        cASConfig
	Certificate      certificate
	Cmdb             cmdb
	Cmdb2            cmdb2
	Domain           domain
	Dong             dong
	EagleEyeData     eagleEyeData
	Effect           effect
	Elastic          elastic
	Email            email
	Emc              emc
	Event            event
	ExtensionMarket  extensionMarket
	ExtensionRecord  extensionRecord
	GridChunk        gridChunk
	GridFile         gridFile
	KVAudit          kVAudit
	KVData           kVData
	LoginLock        loginLock
	Minion           minion
	MinionAccount    minionAccount
	MinionBin        minionBin
	MinionCustomized minionCustomized
	MinionGroup      minionGroup
	MinionListen     minionListen
	MinionLogon      minionLogon
	MinionProcess    minionProcess
	MinionTag        minionTag
	MinionTask       minionTask
	Notifier         notifier
	Oplog            oplog
	PassDNS          passDNS
	PassIP           passIP
	Plate            plate
	Purl             purl
	Recipient        recipient
	Resignation      resignation
	Risk             risk
	RiskDNS          riskDNS
	RiskFile         riskFile
	RiskIP           riskIP
	SBOMComponent    sBOMComponent
	SBOMMinion       sBOMMinion
	SBOMProject      sBOMProject
	SBOMVuln         sBOMVuln
	SIEMServer       sIEMServer
	Startup          startup
	Store            store
	Substance        substance
	SubstanceTask    substanceTask
	SysInfo          sysInfo
	TaskExecute      taskExecute
	TaskExecuteItem  taskExecuteItem
	TaskExtension    taskExtension
	Third            third
	ThirdCustomized  thirdCustomized
	User             user
	VIP              vIP
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		AlertServer:      q.AlertServer.clone(db),
		AuthTemp:         q.AuthTemp.clone(db),
		Broker:           q.Broker.clone(db),
		BrokerBin:        q.BrokerBin.clone(db),
		BrokerStat:       q.BrokerStat.clone(db),
		CASConfig:        q.CASConfig.clone(db),
		Certificate:      q.Certificate.clone(db),
		Cmdb:             q.Cmdb.clone(db),
		Cmdb2:            q.Cmdb2.clone(db),
		Domain:           q.Domain.clone(db),
		Dong:             q.Dong.clone(db),
		EagleEyeData:     q.EagleEyeData.clone(db),
		Effect:           q.Effect.clone(db),
		Elastic:          q.Elastic.clone(db),
		Email:            q.Email.clone(db),
		Emc:              q.Emc.clone(db),
		Event:            q.Event.clone(db),
		ExtensionMarket:  q.ExtensionMarket.clone(db),
		ExtensionRecord:  q.ExtensionRecord.clone(db),
		GridChunk:        q.GridChunk.clone(db),
		GridFile:         q.GridFile.clone(db),
		KVAudit:          q.KVAudit.clone(db),
		KVData:           q.KVData.clone(db),
		LoginLock:        q.LoginLock.clone(db),
		Minion:           q.Minion.clone(db),
		MinionAccount:    q.MinionAccount.clone(db),
		MinionBin:        q.MinionBin.clone(db),
		MinionCustomized: q.MinionCustomized.clone(db),
		MinionGroup:      q.MinionGroup.clone(db),
		MinionListen:     q.MinionListen.clone(db),
		MinionLogon:      q.MinionLogon.clone(db),
		MinionProcess:    q.MinionProcess.clone(db),
		MinionTag:        q.MinionTag.clone(db),
		MinionTask:       q.MinionTask.clone(db),
		Notifier:         q.Notifier.clone(db),
		Oplog:            q.Oplog.clone(db),
		PassDNS:          q.PassDNS.clone(db),
		PassIP:           q.PassIP.clone(db),
		Plate:            q.Plate.clone(db),
		Purl:             q.Purl.clone(db),
		Recipient:        q.Recipient.clone(db),
		Resignation:      q.Resignation.clone(db),
		Risk:             q.Risk.clone(db),
		RiskDNS:          q.RiskDNS.clone(db),
		RiskFile:         q.RiskFile.clone(db),
		RiskIP:           q.RiskIP.clone(db),
		SBOMComponent:    q.SBOMComponent.clone(db),
		SBOMMinion:       q.SBOMMinion.clone(db),
		SBOMProject:      q.SBOMProject.clone(db),
		SBOMVuln:         q.SBOMVuln.clone(db),
		SIEMServer:       q.SIEMServer.clone(db),
		Startup:          q.Startup.clone(db),
		Store:            q.Store.clone(db),
		Substance:        q.Substance.clone(db),
		SubstanceTask:    q.SubstanceTask.clone(db),
		SysInfo:          q.SysInfo.clone(db),
		TaskExecute:      q.TaskExecute.clone(db),
		TaskExecuteItem:  q.TaskExecuteItem.clone(db),
		TaskExtension:    q.TaskExtension.clone(db),
		Third:            q.Third.clone(db),
		ThirdCustomized:  q.ThirdCustomized.clone(db),
		User:             q.User.clone(db),
		VIP:              q.VIP.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		AlertServer:      q.AlertServer.replaceDB(db),
		AuthTemp:         q.AuthTemp.replaceDB(db),
		Broker:           q.Broker.replaceDB(db),
		BrokerBin:        q.BrokerBin.replaceDB(db),
		BrokerStat:       q.BrokerStat.replaceDB(db),
		CASConfig:        q.CASConfig.replaceDB(db),
		Certificate:      q.Certificate.replaceDB(db),
		Cmdb:             q.Cmdb.replaceDB(db),
		Cmdb2:            q.Cmdb2.replaceDB(db),
		Domain:           q.Domain.replaceDB(db),
		Dong:             q.Dong.replaceDB(db),
		EagleEyeData:     q.EagleEyeData.replaceDB(db),
		Effect:           q.Effect.replaceDB(db),
		Elastic:          q.Elastic.replaceDB(db),
		Email:            q.Email.replaceDB(db),
		Emc:              q.Emc.replaceDB(db),
		Event:            q.Event.replaceDB(db),
		ExtensionMarket:  q.ExtensionMarket.replaceDB(db),
		ExtensionRecord:  q.ExtensionRecord.replaceDB(db),
		GridChunk:        q.GridChunk.replaceDB(db),
		GridFile:         q.GridFile.replaceDB(db),
		KVAudit:          q.KVAudit.replaceDB(db),
		KVData:           q.KVData.replaceDB(db),
		LoginLock:        q.LoginLock.replaceDB(db),
		Minion:           q.Minion.replaceDB(db),
		MinionAccount:    q.MinionAccount.replaceDB(db),
		MinionBin:        q.MinionBin.replaceDB(db),
		MinionCustomized: q.MinionCustomized.replaceDB(db),
		MinionGroup:      q.MinionGroup.replaceDB(db),
		MinionListen:     q.MinionListen.replaceDB(db),
		MinionLogon:      q.MinionLogon.replaceDB(db),
		MinionProcess:    q.MinionProcess.replaceDB(db),
		MinionTag:        q.MinionTag.replaceDB(db),
		MinionTask:       q.MinionTask.replaceDB(db),
		Notifier:         q.Notifier.replaceDB(db),
		Oplog:            q.Oplog.replaceDB(db),
		PassDNS:          q.PassDNS.replaceDB(db),
		PassIP:           q.PassIP.replaceDB(db),
		Plate:            q.Plate.replaceDB(db),
		Purl:             q.Purl.replaceDB(db),
		Recipient:        q.Recipient.replaceDB(db),
		Resignation:      q.Resignation.replaceDB(db),
		Risk:             q.Risk.replaceDB(db),
		RiskDNS:          q.RiskDNS.replaceDB(db),
		RiskFile:         q.RiskFile.replaceDB(db),
		RiskIP:           q.RiskIP.replaceDB(db),
		SBOMComponent:    q.SBOMComponent.replaceDB(db),
		SBOMMinion:       q.SBOMMinion.replaceDB(db),
		SBOMProject:      q.SBOMProject.replaceDB(db),
		SBOMVuln:         q.SBOMVuln.replaceDB(db),
		SIEMServer:       q.SIEMServer.replaceDB(db),
		Startup:          q.Startup.replaceDB(db),
		Store:            q.Store.replaceDB(db),
		Substance:        q.Substance.replaceDB(db),
		SubstanceTask:    q.SubstanceTask.replaceDB(db),
		SysInfo:          q.SysInfo.replaceDB(db),
		TaskExecute:      q.TaskExecute.replaceDB(db),
		TaskExecuteItem:  q.TaskExecuteItem.replaceDB(db),
		TaskExtension:    q.TaskExtension.replaceDB(db),
		Third:            q.Third.replaceDB(db),
		ThirdCustomized:  q.ThirdCustomized.replaceDB(db),
		User:             q.User.replaceDB(db),
		VIP:              q.VIP.replaceDB(db),
	}
}

type queryCtx struct {
	AlertServer      *alertServerDo
	AuthTemp         *authTempDo
	Broker           *brokerDo
	BrokerBin        *brokerBinDo
	BrokerStat       *brokerStatDo
	CASConfig        *cASConfigDo
	Certificate      *certificateDo
	Cmdb             *cmdbDo
	Cmdb2            *cmdb2Do
	Domain           *domainDo
	Dong             *dongDo
	EagleEyeData     *eagleEyeDataDo
	Effect           *effectDo
	Elastic          *elasticDo
	Email            *emailDo
	Emc              *emcDo
	Event            *eventDo
	ExtensionMarket  *extensionMarketDo
	ExtensionRecord  *extensionRecordDo
	GridChunk        *gridChunkDo
	GridFile         *gridFileDo
	KVAudit          *kVAuditDo
	KVData           *kVDataDo
	LoginLock        *loginLockDo
	Minion           *minionDo
	MinionAccount    *minionAccountDo
	MinionBin        *minionBinDo
	MinionCustomized *minionCustomizedDo
	MinionGroup      *minionGroupDo
	MinionListen     *minionListenDo
	MinionLogon      *minionLogonDo
	MinionProcess    *minionProcessDo
	MinionTag        *minionTagDo
	MinionTask       *minionTaskDo
	Notifier         *notifierDo
	Oplog            *oplogDo
	PassDNS          *passDNSDo
	PassIP           *passIPDo
	Plate            *plateDo
	Purl             *purlDo
	Recipient        *recipientDo
	Resignation      *resignationDo
	Risk             *riskDo
	RiskDNS          *riskDNSDo
	RiskFile         *riskFileDo
	RiskIP           *riskIPDo
	SBOMComponent    *sBOMComponentDo
	SBOMMinion       *sBOMMinionDo
	SBOMProject      *sBOMProjectDo
	SBOMVuln         *sBOMVulnDo
	SIEMServer       *sIEMServerDo
	Startup          *startupDo
	Store            *storeDo
	Substance        *substanceDo
	SubstanceTask    *substanceTaskDo
	SysInfo          *sysInfoDo
	TaskExecute      *taskExecuteDo
	TaskExecuteItem  *taskExecuteItemDo
	TaskExtension    *taskExtensionDo
	Third            *thirdDo
	ThirdCustomized  *thirdCustomizedDo
	User             *userDo
	VIP              *vIPDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AlertServer:      q.AlertServer.WithContext(ctx),
		AuthTemp:         q.AuthTemp.WithContext(ctx),
		Broker:           q.Broker.WithContext(ctx),
		BrokerBin:        q.BrokerBin.WithContext(ctx),
		BrokerStat:       q.BrokerStat.WithContext(ctx),
		CASConfig:        q.CASConfig.WithContext(ctx),
		Certificate:      q.Certificate.WithContext(ctx),
		Cmdb:             q.Cmdb.WithContext(ctx),
		Cmdb2:            q.Cmdb2.WithContext(ctx),
		Domain:           q.Domain.WithContext(ctx),
		Dong:             q.Dong.WithContext(ctx),
		EagleEyeData:     q.EagleEyeData.WithContext(ctx),
		Effect:           q.Effect.WithContext(ctx),
		Elastic:          q.Elastic.WithContext(ctx),
		Email:            q.Email.WithContext(ctx),
		Emc:              q.Emc.WithContext(ctx),
		Event:            q.Event.WithContext(ctx),
		ExtensionMarket:  q.ExtensionMarket.WithContext(ctx),
		ExtensionRecord:  q.ExtensionRecord.WithContext(ctx),
		GridChunk:        q.GridChunk.WithContext(ctx),
		GridFile:         q.GridFile.WithContext(ctx),
		KVAudit:          q.KVAudit.WithContext(ctx),
		KVData:           q.KVData.WithContext(ctx),
		LoginLock:        q.LoginLock.WithContext(ctx),
		Minion:           q.Minion.WithContext(ctx),
		MinionAccount:    q.MinionAccount.WithContext(ctx),
		MinionBin:        q.MinionBin.WithContext(ctx),
		MinionCustomized: q.MinionCustomized.WithContext(ctx),
		MinionGroup:      q.MinionGroup.WithContext(ctx),
		MinionListen:     q.MinionListen.WithContext(ctx),
		MinionLogon:      q.MinionLogon.WithContext(ctx),
		MinionProcess:    q.MinionProcess.WithContext(ctx),
		MinionTag:        q.MinionTag.WithContext(ctx),
		MinionTask:       q.MinionTask.WithContext(ctx),
		Notifier:         q.Notifier.WithContext(ctx),
		Oplog:            q.Oplog.WithContext(ctx),
		PassDNS:          q.PassDNS.WithContext(ctx),
		PassIP:           q.PassIP.WithContext(ctx),
		Plate:            q.Plate.WithContext(ctx),
		Purl:             q.Purl.WithContext(ctx),
		Recipient:        q.Recipient.WithContext(ctx),
		Resignation:      q.Resignation.WithContext(ctx),
		Risk:             q.Risk.WithContext(ctx),
		RiskDNS:          q.RiskDNS.WithContext(ctx),
		RiskFile:         q.RiskFile.WithContext(ctx),
		RiskIP:           q.RiskIP.WithContext(ctx),
		SBOMComponent:    q.SBOMComponent.WithContext(ctx),
		SBOMMinion:       q.SBOMMinion.WithContext(ctx),
		SBOMProject:      q.SBOMProject.WithContext(ctx),
		SBOMVuln:         q.SBOMVuln.WithContext(ctx),
		SIEMServer:       q.SIEMServer.WithContext(ctx),
		Startup:          q.Startup.WithContext(ctx),
		Store:            q.Store.WithContext(ctx),
		Substance:        q.Substance.WithContext(ctx),
		SubstanceTask:    q.SubstanceTask.WithContext(ctx),
		SysInfo:          q.SysInfo.WithContext(ctx),
		TaskExecute:      q.TaskExecute.WithContext(ctx),
		TaskExecuteItem:  q.TaskExecuteItem.WithContext(ctx),
		TaskExtension:    q.TaskExtension.WithContext(ctx),
		Third:            q.Third.WithContext(ctx),
		ThirdCustomized:  q.ThirdCustomized.WithContext(ctx),
		User:             q.User.WithContext(ctx),
		VIP:              q.VIP.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
