package crontab

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Table struct {
	crond *cron.Cron
	mutex sync.Mutex
	tasks map[string]cron.EntryID
}

func NewTable(opts ...cron.Option) *Table {
	return &Table{
		crond: cron.New(opts...),
		tasks: make(map[string]cron.EntryID, 16),
	}
}

func (tbl *Table) Start() {
	tbl.crond.Start()
}

func (tbl *Table) Stop() {
	tbl.crond.Stop()
}

func (tbl *Table) res() {

}
