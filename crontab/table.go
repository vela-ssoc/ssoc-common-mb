package crontab

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Table struct {
	crond *cron.Cron
	mutex sync.Mutex
	tasks map[string]any
}

func NewTable(opts ...cron.Option) *Table {
	return &Table{
		crond: cron.New(opts...),
		tasks: make(map[string]any, 16),
	}
}

func (tbl *Table) Start() {
	tbl.crond.Start()
}

func (tbl *Table) Stop() {
	tbl.crond.Stop()
}
