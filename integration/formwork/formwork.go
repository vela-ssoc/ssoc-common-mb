package formwork

import (
	"bytes"
	"context"
	"io"

	"github.com/vela-ssoc/vela-common-mb/logback"
)

type Loader interface {
	Loaders() []TmplLoader
	LoginDong(ctx context.Context, title, body any) (string, string)
}

func NewStore(slog logback.Logger) Loader {
	return &tmplStore{
		slog:           slog,
		loginDongTitle: newTmplLoad("global.ding.title"),
		loginDongBody:  newTmplLoad("global.ding.tmpl"),
	}
}

type tmplStore struct {
	slog           logback.Logger
	loginDongTitle TmplLoader
	loginDongBody  TmplLoader
}

func (ts *tmplStore) Loaders() []TmplLoader {
	return []TmplLoader{
		ts.loginDongTitle,
		ts.loginDongBody,
	}
}

func (ts *tmplStore) LoginDong(ctx context.Context, title, body any) (string, string) {
	tbuf, bbuf := new(bytes.Buffer), new(bytes.Buffer)
	if err := ts.render(ctx, tbuf, ts.loginDongTitle, title); err != nil {
		ts.slog.Warnf("渲染咚咚登录验证码标题出错：%s", err)
		tbuf.Reset()
		tbuf.WriteString("咚咚登录验证码")
	}
	if err := ts.render(ctx, bbuf, ts.loginDongBody, body); err != nil {
		ts.slog.Warnf("渲染咚咚登录验证码内容出错：%s", err)
		bbuf.Reset()
		bbuf.WriteString("咚咚登录模板未配置")
	}
	return tbuf.String(), bbuf.String()
}

func (ts *tmplStore) render(ctx context.Context, w io.Writer, tl TmplLoader, data any) error {
	tpl, err := tl.Load(ctx)
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}
