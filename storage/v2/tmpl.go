package storage

import (
	"context"
	htm "html/template"
	"io"
	"sync"
	txt "text/template"
)

type templateExecutor interface {
	Execute(io.Writer, any) error
}

type tmplValuer interface {
	valuer
	rend(ctx context.Context, w io.Writer, v any) error
}

type valueTmpl struct {
	under  valuer
	escape bool
	mutex  sync.RWMutex
	loaded bool
	tmpl   templateExecutor
	err    error
}

func (v *valueTmpl) id() string {
	return v.under.id()
}

func (v *valueTmpl) load(ctx context.Context) ([]byte, error) {
	return v.under.load(ctx)
}

func (v *valueTmpl) forget() bool {
	v.mutex.Lock()
	v.loaded, v.tmpl = false, nil
	shared := v.under.forget()
	v.mutex.Unlock()

	return shared
}

func (v *valueTmpl) validate(data []byte) error {
	return v.under.validate(data)
}

func (v *valueTmpl) rend(ctx context.Context, w io.Writer, data any) error {
	exec, err := v.loadExecutor(ctx)
	if err == nil {
		err = exec.Execute(w, data)
	}
	return err
}

func (v *valueTmpl) loadExecutor(ctx context.Context) (templateExecutor, error) {
	v.mutex.RLock()
	loaded, tmpl, err := v.loaded, v.tmpl, v.err
	v.mutex.RUnlock()
	if loaded {
		return tmpl, err
	}

	return v.slowLoadExecutor(ctx)
}

func (v *valueTmpl) slowLoadExecutor(ctx context.Context) (templateExecutor, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if v.loaded {
		return v.tmpl, v.err
	}

	data, err := v.under.load(ctx)
	if err != nil {
		v.err = err
		v.loaded = true
		return nil, err
	}

	uid := v.under.id()
	tmpl, err := v.newExecutor(uid, data)

	v.tmpl = tmpl
	v.err = err
	v.loaded = true

	return tmpl, err
}

func (v *valueTmpl) newExecutor(id string, data []byte) (templateExecutor, error) {
	if v.escape {
		return htm.New(id).Parse(string(data))
	}
	return txt.New(id).Parse(string(data))
}
