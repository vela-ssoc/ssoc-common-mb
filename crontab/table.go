package crontab

import (
	"log/slog"
	"sync"

	"github.com/robfig/cron/v3"
)

type Table struct {
	log   *slog.Logger
	crond *cron.Cron
	mutex sync.RWMutex
	jobs  map[string]JobInfo
}

func NewTable(log *slog.Logger, opts ...cron.Option) *Table {
	return &Table{
		log:   log,
		crond: cron.New(opts...),
		jobs:  make(map[string]JobInfo, 16),
	}
}

func (tbl *Table) Start() {
	tbl.crond.Start()
}

func (tbl *Table) Stop() {
	tbl.crond.Stop()
}

func (tbl *Table) AddJob(job Jober) error {
	last, curr, exists, err := tbl.addJob(job)
	if err != nil {
		tbl.log.Error("任务新增出错", "current", curr, "err", err)
	} else if exists {
		tbl.log.Warn("任务替换成功", "current", curr, "last", last)
	} else {
		tbl.log.Info("任务添加成功", "current", curr)
	}

	return err
}

func (tbl *Table) AddJobs(jobs []Jober) error {
	for _, job := range jobs {
		if err := tbl.AddJob(job); err != nil {
			return err
		}
	}

	return nil
}

func (tbl *Table) All() []JobInfo {
	tbl.mutex.RLock()
	defer tbl.mutex.RUnlock()

	index := tbl.entries()
	infos := make([]JobInfo, 0, len(index))

	for _, info := range tbl.jobs {
		ent := index[info.EntryID]
		info.Next = ent.Next
		infos = append(infos, info)
	}

	return infos
}

func (tbl *Table) Remove(id string) (JobInfo, bool) {
	tbl.mutex.Lock()
	defer tbl.mutex.Unlock()

	return tbl.remove(id)
}

func (tbl *Table) Sweep() []JobInfo {
	tbl.mutex.Lock()
	defer tbl.mutex.Unlock()

	index := tbl.entries()
	infos := make([]JobInfo, 0, len(index))
	for _, inf := range tbl.jobs {
		entryID := inf.EntryID
		ent := index[entryID]
		if !ent.Next.IsZero() {
			continue
		}

		tbl.crond.Remove(entryID)
		delete(tbl.jobs, inf.ID)
		infos = append(infos, inf)
	}

	return infos
}

func (tbl *Table) addJob(job Jober) (last, curr JobInfo, lastExists bool, err error) {
	sjob := tbl.newSafeJob(job)
	if sjob.opt.ID == "" {
		sjob.opt.ID = QualifiedID(job)
	}
	opt := sjob.opt
	id := opt.ID
	curr = JobInfo{JobOptions: opt}

	tbl.mutex.Lock()
	defer tbl.mutex.Unlock()

	last, lastExists = tbl.jobs[id]
	if sjob.opt.Sched != nil {
		curr.EntryID = tbl.crond.Schedule(opt.Sched, sjob)
	} else {
		curr.EntryID, err = tbl.crond.AddJob(opt.Expr, sjob)
	}
	if err != nil {
		return
	}

	tbl.jobs[id] = curr
	if lastExists {
		tbl.crond.Remove(last.EntryID)
	}
	if opt.Immediate {
		go sjob.Run()
	}

	return
}

func (tbl *Table) entries() map[cron.EntryID]cron.Entry {
	entries := tbl.crond.Entries()
	index := make(map[cron.EntryID]cron.Entry, len(entries))
	for _, e := range entries {
		index[e.ID] = e
	}

	return index
}

func (tbl *Table) remove(id string) (JobInfo, bool) {
	inf, exists := tbl.jobs[id]
	if !exists {
		return inf, false
	}

	entID := inf.EntryID
	ent := tbl.crond.Entry(entID)
	inf.Next = ent.Next

	tbl.crond.Remove(entID)
	delete(tbl.jobs, id)

	return inf, true
}

func (tbl *Table) newSafeJob(job Jober) *safeJob {
	opt := job.Options()
	return &safeJob{
		job: job,
		opt: opt,
		log: tbl.log,
	}
}
