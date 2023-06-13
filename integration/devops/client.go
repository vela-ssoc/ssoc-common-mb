package devops

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mba/netutil"
)

// {
//    "origin_name": "test001",
//    "alert_type": "System",
//    "alert_object": "cpu",
//    "alert_attribute": "load",
//    "severity": "middle",
//    "subject": "这是一个测试",
//    "tags": "hostname=localhost,a=b,c=d",
//    "notifier": "[{\"name\":\"张三\",\"mobile\":\"15927458888\",\"email\":\"zhangsan@eastmoney.com\",\"serial_number\":\"160001\",\"notify_methods\":\"weixin,sms\"}]"
//}

// def alarm():
//    data = dict(alert_type="honeypot",
//                alert_object="honeypot found scanning",
//                alert_attribute="scan",
//                origin_name="security-alert",
//                body="蜜罐发现扫描行为，详情查看邮件内容",
//                severity="disaster",
//                subject="蜜罐发现扫描行为,详情请查看告警邮件")
//
//    sunke = dict(name="孙科",
//                mobile="18521351052",
//                email="sunke@eastmoney.com",
//                serial_number="161101",
//                notify_methods="call")
//    feiliang = dict(
//                name = "费亮",
//                mobile = "13818852180",
//                email = "feiliang@eastmoney.com",
//                serial_number = "130653",
//                notify_methods = "call"
//            )
//    data["notifier"] = json.dumps([sunke, feiliang])
//    url = "http://alert.eastmoney.com/api/v0.1/alerts"
//    try:
//        r = requests.post(url, data=data, timeout=30)
//        resp = r.json()
//        print(resp)
//    except Exception as e:
//        print(str(e))

type Client interface {
	Send(ctx context.Context, title, body string, users []*model.Devops) error
}

func NewClient(client netutil.HTTPClient) Client {
	return &devopsClient{
		client: client,
	}
}

type devopsClient struct {
	client netutil.HTTPClient
}

func (dc *devopsClient) Send(ctx context.Context, title, body string, users []*model.Devops) error {
	dat, _ := json.Marshal(users)
	req := &request{
		OriginName:     "security-alert",
		AlertType:      "security-alert",
		AlertObject:    title,
		AlertAttribute: "ssoc",
		Severity:       "disaster",
		Subject:        title,
		Body:           body,
		Notifier:       string(dat),
	}

	addr := "http://alert.eastmoney.com/api/v0.1/alerts"
	return dc.client.JSON(ctx, http.MethodPost, addr, req, &struct{}{}, nil)
}

type request struct {
	OriginName     string `json:"origin_name"`
	AlertType      string `json:"alert_type"`
	AlertObject    string `json:"alert_object"`
	AlertAttribute string `json:"alert_attribute"`
	Severity       string `json:"severity"`
	Subject        string `json:"subject"`
	Body           string `json:"body"`
	Notifier       string `json:"notifier"`
}
