package model

import (
	"maps"
	"net/http"
	"strings"
)

type HTTPHeader map[string]string

func (h HTTPHeader) Canonical() HTTPHeader {
	ret := make(HTTPHeader, len(h))
	for k, v := range h {
		key := http.CanonicalHeaderKey(k)
		ret[key] = v
	}

	return ret
}

func (h HTTPHeader) Header() http.Header {
	ret := make(http.Header, len(h))
	for k, v := range h {
		ret.Set(k, v)
	}

	return ret
}

func (h HTTPHeader) Host() (string, bool) {
	for k, v := range h {
		if strings.EqualFold(k, "host") {
			return v, true
		}
	}

	return "", false
}

func (h HTTPHeader) Array() []string {
	arr := make([]string, 0, len(h))
	for k, v := range h {
		arr = append(arr, k+": "+v)
	}

	return arr
}

func (h HTTPHeader) Clone() HTTPHeader {
	hm := maps.Clone(h)
	if hm == nil {
		hm = HTTPHeader{}
	}

	return hm
}
