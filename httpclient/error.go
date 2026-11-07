package httpclient

import (
	"net/http"
	"strconv"
)

type HTTPError struct {
	Code     int             `json:"code"`
	Body     []byte          `json:"body"`
	RFC9457  *ProblemDetails `json:"rfc9457"`
	Request  *http.Request   `json:"-"`
	Response *http.Response  `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.RFC9457 != nil {
		return e.RFC9457.String()
	}

	return "HTTP 服务响应错误 code=" + strconv.Itoa(e.Code) +
		", body=" + strconv.Quote(string(e.Body))
}

// ProblemDetails [RFC 9457] Problem Details for HTTP APIs.
//
// [RFC 9457]: https://www.rfc-editor.org/info/rfc9457/
type ProblemDetails struct {
	Type     string `json:"type"`
	Status   int    `json:"status"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func (d *ProblemDetails) String() string {
	return "HTTP 服务器响应错误 instance='" + d.Instance + "'" +
		", status=" + strconv.Itoa(d.Status) +
		", title='" + d.Title + "'" +
		", detail='" + d.Detail + "'"
}

func (d *ProblemDetails) Error() string {
	return d.String()
}
