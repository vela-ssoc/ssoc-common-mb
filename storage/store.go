package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strings"
	"text/template"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
)

type Storer interface {
	Reset(id string)
	Shared(id string) bool
	Invalid(id string, val []byte) bool
	LocalAddr(ctx context.Context) (string, error)
	CmdbURL(ctx context.Context) (string, error)
	SsoURL(ctx context.Context) (string, error)
	Startup(ctx context.Context) (*model.Startup, error)
	Deploy(ctx context.Context, goos string, v any) *bytes.Buffer
	LoginDong(ctx context.Context, v any) (title, body string)
	RiskDong(ctx context.Context, v any) (title, body string)
	RiskHTML(ctx context.Context, v any) *bytes.Buffer
	EventDong(ctx context.Context, v any) (title, body string)
	EventHTML(ctx context.Context, v any) *bytes.Buffer
}

func NewStore() Storer {
	ret := &storeDB{values: make(map[string]Valuer, 16)}

	{
		val := ret.newValue("global.local.addr", false, validIP)
		ret.values[val.ID()] = val
		ret.localAddr = val
	}
	{
		val := ret.newHTTP("global.cmdb.url", true)
		ret.values[val.ID()] = val
		ret.cmdbURL = val
	}
	{
		val := ret.newHTTP("global.sso.url", false)
		ret.values[val.ID()] = val
		ret.ssoURL = val
	}
	{
		val := ret.newValue("global.startup.param", true, validJSON)
		ret.values[val.ID()] = val
		ret.startupParam = val
	}

	// ============ 部署脚本
	{
		val := ret.newTmpl("global.deploy.linux.tmpl", true)
		ret.values[val.ID()] = val
		ret.deployLinuxTmpl = val
	}
	{
		val := ret.newTmpl("global.deploy.windows.tmpl", true)
		ret.values[val.ID()] = val
		ret.deployWindowsTmpl = val
	}
	{
		val := ret.newTmpl("global.deploy.darwin.tmpl", true)
		ret.values[val.ID()] = val
		ret.deployDarwinTmpl = val
	}
	{
		val := ret.newTmpl("global.deploy.android.tmpl", true)
		ret.values[val.ID()] = val
		ret.deployAndroidTmpl = val
	}

	// 咚咚登录验证码
	{
		val := ret.newTmpl("global.ding.title", false)
		ret.values[val.ID()] = val
		ret.loginDongTitle = val
	}
	{
		val := ret.newTmpl("global.ding.tmpl", false)
		ret.values[val.ID()] = val
		ret.loginDongBody = val
	}

	//  ============ event/risk
	{
		val := ret.newTmpl("global.risk.dong.title", true)
		ret.values[val.ID()] = val
		ret.riskDongTitle = val
	}
	{
		val := ret.newTmpl("global.risk.dong.tmpl", true)
		ret.values[val.ID()] = val
		ret.riskDongBody = val
	}
	{
		val := ret.newTmpl("global.risk.html.tmpl", false)
		ret.values[val.ID()] = val
		ret.riskHTML = val
	}
	{
		val := ret.newTmpl("global.event.dong.title", true)
		ret.values[val.ID()] = val
		ret.eventDongTitle = val
	}
	{
		val := ret.newTmpl("global.event.dong.tmpl", true)
		ret.values[val.ID()] = val
		ret.eventDongBody = val
	}
	{
		val := ret.newTmpl("global.event.html.tmpl", false)
		ret.values[val.ID()] = val
		ret.eventHTML = val
	}

	return ret
}

type storeDB struct {
	values map[string]Valuer

	localAddr    Valuer
	cmdbURL      httpValuer
	ssoURL       httpValuer
	startupParam Valuer

	deployLinuxTmpl   storeRender
	deployWindowsTmpl storeRender
	deployDarwinTmpl  storeRender
	deployAndroidTmpl storeRender

	loginDongTitle storeRender
	loginDongBody  storeRender

	riskDongTitle  storeRender
	riskDongBody   storeRender
	riskHTML       storeRender
	eventDongTitle storeRender
	eventDongBody  storeRender
	eventHTML      storeRender
}

func (sdb *storeDB) Reset(id string) {
	if val := sdb.values[id]; val != nil {
		val.Reset()
	}
}

func (sdb *storeDB) Shared(id string) bool {
	if val := sdb.values[id]; val != nil {
		return val.Shared()
	}
	return false
}

func (sdb *storeDB) Invalid(id string, dat []byte) bool {
	if val := sdb.values[id]; val != nil {
		return val.Invalid(dat)
	}
	return false
}

func (sdb *storeDB) LocalAddr(ctx context.Context) (string, error) {
	val, err := sdb.localAddr.Value(ctx)
	return string(val), err
}

func (sdb *storeDB) CmdbURL(ctx context.Context) (string, error) {
	return sdb.cmdbURL.Addr(ctx)
}

func (sdb *storeDB) SsoURL(ctx context.Context) (string, error) {
	return sdb.ssoURL.Addr(ctx)
}

func (sdb *storeDB) Startup(ctx context.Context) (*model.Startup, error) {
	val, err := sdb.startupParam.Value(ctx)
	if err != nil {
		return nil, err
	}
	ret := new(model.Startup)
	err = json.Unmarshal(val, ret)
	return ret, err
}

func (sdb *storeDB) LoginDong(ctx context.Context, v any) (string, string) {
	title := sdb.loginDongTitle.Rend(ctx, v)
	body := sdb.loginDongBody.Rend(ctx, v)
	return title.String(), body.String()
}

func (sdb *storeDB) RiskDong(ctx context.Context, v any) (string, string) {
	title := sdb.riskDongTitle.Rend(ctx, v)
	body := sdb.riskDongBody.Rend(ctx, v)
	return title.String(), body.String()
}

func (sdb *storeDB) RiskHTML(ctx context.Context, v any) *bytes.Buffer {
	return sdb.riskHTML.Rend(ctx, v)
}

func (sdb *storeDB) EventDong(ctx context.Context, v any) (string, string) {
	title := sdb.eventDongTitle.Rend(ctx, v)
	body := sdb.eventDongBody.Rend(ctx, v)
	return title.String(), body.String()
}

func (sdb *storeDB) EventHTML(ctx context.Context, v any) *bytes.Buffer {
	return sdb.eventHTML.Rend(ctx, v)
}

func (sdb *storeDB) Deploy(ctx context.Context, goos string, v any) *bytes.Buffer {
	var rend storeRender
	switch strings.ToLower(goos) {
	case "linux":
		rend = sdb.deployLinuxTmpl
	case "windows":
		rend = sdb.deployWindowsTmpl
	case "darwin":
		rend = sdb.deployDarwinTmpl
	case "android":
		rend = sdb.deployAndroidTmpl
	default:
		msg := fmt.Sprintf("不支持的 %s 操作系统部署脚本", goos)
		return bytes.NewBufferString(msg)
	}

	return rend.Rend(ctx, v)
}

func (sdb *storeDB) newValue(id string, shared bool, valid func([]byte) bool) Valuer {
	return &storeValue{
		id:     id,
		shared: shared,
		valid:  valid,
	}
}

func (sdb *storeDB) newTmpl(id string, shared bool) storeRender {
	val := &storeValue{
		id:     id,
		shared: shared,
		valid:  validTmpl,
	}
	return &storeTemplate{
		value: val,
	}
}

func (sdb *storeDB) newHTTP(id string, shared bool) httpValuer {
	val := &storeValue{
		id:     id,
		shared: shared,
		valid:  validHTTP,
	}
	return &httpValue{
		value: val,
	}
}

func validIP(dat []byte) bool {
	ip := net.ParseIP(string(dat))
	return ip.IsUnspecified() || ip.IsLoopback()
}

func validJSON(dat []byte) bool {
	return !json.Valid(dat)
}

func validTmpl(dat []byte) bool {
	_, err := template.New("valid").Parse(string(dat))
	return err != nil
}

func validHTTP(dat []byte) bool {
	u, err := url.Parse(string(dat))
	if err != nil {
		return true
	}
	scheme := u.Scheme
	return scheme != "http" && scheme != "https"
}
