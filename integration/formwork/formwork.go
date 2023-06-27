package formwork

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/vela-ssoc/vela-common-mb/logback"
)

type Render interface {
	Loaders() []TmplLoader
	LoginDong(ctx context.Context, title, body any) (string, string)
	RiskDong(ctx context.Context, title, body any) (string, string)
	EventDong(ctx context.Context, title, body any) (string, string)
}

func NewRend(slog logback.Logger) Render {
	return &tmplStore{
		slog:           slog,
		loginDongTitle: newTmplLoad("global.ding.title", false),
		loginDongBody:  newTmplLoad("global.ding.tmpl", false),
		riskDongTitle:  newTmplLoad("global.risk.dong.title", true),
		riskDongBody:   newTmplLoad("global.risk.dong.tmpl", true),
		eventDongTitle: newTmplLoad("global.event.dong.title", true),
		eventDongBody:  newTmplLoad("global.event.dong.tmpl", true),
	}
}

type tmplStore struct {
	slog           logback.Logger
	loginDongTitle TmplLoader
	loginDongBody  TmplLoader
	riskDongTitle  TmplLoader
	riskDongBody   TmplLoader
	eventDongTitle TmplLoader
	eventDongBody  TmplLoader
}

func (ts *tmplStore) Loaders() []TmplLoader {
	return []TmplLoader{
		ts.loginDongTitle,
		ts.loginDongBody,
		ts.riskDongTitle,
		ts.riskDongBody,
		ts.eventDongTitle,
		ts.eventDongBody,
	}
}

func (ts *tmplStore) LoginDong(ctx context.Context, title, body any) (string, string) {
	tbuf, bbuf := new(bytes.Buffer), new(bytes.Buffer)
	if err := ts.render(ctx, tbuf, ts.loginDongTitle, title); err != nil {
		msg := fmt.Sprintf("渲染咚咚登录验证码标题出错：%s", err)
		ts.slog.Warn(msg)
		tbuf.Reset()
		tbuf.WriteString(msg)
	}
	if err := ts.render(ctx, bbuf, ts.loginDongBody, body); err != nil {
		msg := fmt.Sprintf("渲染咚咚登录验证码内容出错：%s", err)
		ts.slog.Warn(msg)
		bbuf.Reset()
		bbuf.WriteString(msg)
	}
	return tbuf.String(), bbuf.String()
}

func (ts *tmplStore) RiskDong(ctx context.Context, title, body any) (string, string) {
	tbuf, bbuf := new(bytes.Buffer), new(bytes.Buffer)
	if err := ts.render(ctx, tbuf, ts.riskDongTitle, title); err != nil {
		msg := fmt.Sprintf("风险标题渲染出错：%s", err)
		ts.slog.Warn(msg)
		tbuf.Reset()
		tbuf.WriteString(msg)
	}
	if err := ts.render(ctx, bbuf, ts.riskDongBody, body); err != nil {
		msg := fmt.Sprintf("风险内容渲染出错：%s", err)
		ts.slog.Warn(msg)
		bbuf.Reset()
		bbuf.WriteString(msg)
	}
	return tbuf.String(), bbuf.String()
}

func (ts *tmplStore) EventDong(ctx context.Context, title, body any) (string, string) {
	tbuf, bbuf := new(bytes.Buffer), new(bytes.Buffer)
	if err := ts.render(ctx, tbuf, ts.eventDongTitle, title); err != nil {
		msg := fmt.Sprintf("事件通知标题渲染出错：%s", err)
		ts.slog.Warn(msg)
		tbuf.Reset()
		tbuf.WriteString(msg)
	}
	if err := ts.render(ctx, bbuf, ts.eventDongTitle, body); err != nil {
		msg := fmt.Sprintf("事件通知内容渲染出错：%s", err)
		ts.slog.Warn(msg)
		bbuf.Reset()
		bbuf.WriteString(msg)
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
