package formwork

import (
	"context"
	"html/template"
	"sync"

	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

type TmplLoader interface {
	ID() string
	Shared() bool
	Reset()
	Validate([]byte) error
	Load(ctx context.Context) (*template.Template, error)
}

func newTmplLoad(id string, share bool) TmplLoader {
	return &tmplLoad{id: id, shared: share}
}

type tmplLoad struct {
	id     string
	shared bool
	mutex  sync.RWMutex
	done   bool
	err    error
	tmpl   *template.Template
}

func (t *tmplLoad) ID() string {
	return t.id
}

func (t *tmplLoad) Shared() bool {
	return t.shared
}

func (t *tmplLoad) Reset() {
	t.mutex.Lock()
	t.done, t.err, t.done, t.tmpl = false, nil, true, nil
	t.mutex.Unlock()
}

func (t *tmplLoad) Validate(dat []byte) error {
	_, err := template.New(t.id).Parse(string(dat))
	return err
}

func (t *tmplLoad) Load(ctx context.Context) (*template.Template, error) {
	t.mutex.RLock()
	done, err, tmpl := t.done, t.err, t.tmpl
	t.mutex.RUnlock()
	if done {
		return tmpl, err
	}

	return t.slowLoad(ctx)
}

func (t *tmplLoad) slowLoad(ctx context.Context) (*template.Template, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.done {
		return t.tmpl, t.err
	}

	tbl := query.Store
	dat, err := tbl.WithContext(ctx).Where(tbl.ID.Eq(t.id)).First()
	if err != nil {
		t.err = err
		t.done = true
		return nil, err
	}
	tmpl, err := template.New(t.id).Parse(string(dat.Value))
	t.err, t.tmpl = err, tmpl
	t.done = true

	return tmpl, err
}
