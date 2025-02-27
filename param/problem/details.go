package problem

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Details is Problem Details for HTTP APIs.
// RFC7807 https://www.rfc-editor.org/rfc/rfc7807
type Details struct {
	Host string `json:"host" xml:"host"`
	// Type A URI reference [RFC3986] that identifies the problem type. This specification
	// encourages that, when dereferenced, it provide human-readable documentation for
	// the problem type (e.g., using HTML [W3C.REC-html5-20141028]).
	// When this member is not present, its value is assumed to be "about:blank".
	Type string `json:"type" xml:"type"`

	// Title A short, human-readable summary of the problem type.  It SHOULD NOT change from
	// occurrence to occurrence of the problem, except for purposes of localization (e.g., using
	// proactive content negotiation; see [RFC7231], Section 3.4).
	Title string `json:"title" xml:"title"`

	// Status The HTTP status code ([RFC7231], Section 6) generated by the origin server for
	// this occurrence of the problem.
	Status int `json:"status" xml:"status"`

	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail" xml:"detail"`

	// Instance A URI reference that identifies the specific occurrence of the problem.
	// It may or may not yield further information if dereferenced.
	Instance string `json:"instance" xml:"instance"`

	Method string `json:"method" xml:"method"`

	Datetime time.Time `json:"datetime" xml:"datetime"`
}

func (d Details) JSON(w http.ResponseWriter) error {
	code := d.Status
	if code < http.StatusBadRequest || code >= 600 {
		code = http.StatusBadRequest
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(d); err == nil {
		size := int64(buf.Len())
		w.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	}
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("Content-Language", "zh")
	w.WriteHeader(code)
	_, err := buf.WriteTo(w)

	return err
}

func (d Details) String() string {
	return "problem detail, type='" + d.Type + "'" +
		", title='" + d.Title + "'" +
		", status=" + strconv.FormatInt(int64(d.Status), 10) +
		", detail='" + d.Detail + "'" +
		", instance='" + d.Instance + "'"
}
