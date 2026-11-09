package httpship

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/vela-ssoc/ssoc-common-mb/httpclient"
	"github.com/vela-ssoc/ssoc-common-mb/validation"
	"github.com/xgfone/ship/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

func ErrorHandle(w http.ResponseWriter, r *http.Request, e error) {
	v := &httpclient.ProblemDetails{
		Type:     "about:blank",
		Status:   http.StatusBadRequest,
		Title:    http.StatusText(http.StatusBadRequest),
		Detail:   e.Error(),
		Instance: r.URL.Path,
	}
	switch err := e.(type) {
	case *httpclient.ProblemDetails:
		v = err
	case ship.HTTPServerError:
		v.Status = err.Code
	case *ship.HTTPServerError:
		v.Status = err.Code
	case *validation.ValidError:
		v.Title = "参数校验错误"
	case *time.ParseError:
		v.Title = "参数格式错误"
		v.Detail = "时间格式错误，正确格式：" + err.Layout
	case *net.ParseError:
		v.Title = "参数格式错误"
		v.Detail = err.Text + " 不是有效的 " + err.Type
	case base64.CorruptInputError:
		v.Title = "参数格式错误"
		v.Detail = "base64 编码错误：" + err.Error()
	case *json.SyntaxError:
		v.Title = "报文格式错误"
		v.Detail = "请求报错必须是 JSON 格式"
	case *json.UnmarshalTypeError:
		v.Title = "数据类型错误"
		v.Detail = err.Field + " 收到无效的数据类型"
	case *strconv.NumError:
		v.Title = "数据类型错误"
		var msg string
		if sn := strings.SplitN(err.Func, "Parse", 2); len(sn) == 2 {
			msg = err.Num + " 不是 " + strings.ToLower(sn[1]) + " 类型"
		} else {
			msg = "类型错误：" + err.Num
		}
		v.Detail = msg
	case *mysql.MySQLError:
		switch err.Number {
		case 1062:
			v.Detail = "数据已存在"
		default:
			v.Status = http.StatusInternalServerError
			v.Detail = "内部错误"
		}
	default:
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			v.Detail = "数据不存在"
		case errors.Is(err, mongo.ErrNilDocument):
			v.Detail = "数据不存在"
		}
	}

	_ = writeJSON(w, r, v)
}

func writeJSON(w http.ResponseWriter, r *http.Request, v *httpclient.ProblemDetails) error {
	code := v.Status
	if code < http.StatusBadRequest || code >= 600 {
		code = http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(code)

	if r.Method == http.MethodHead {
		return nil
	}

	return json.NewEncoder(w).Encode(v)
}
