package request

import "github.com/vela-ssoc/vela-common-mb/dynsql"

type Pages struct {
	Page int64 `query:"page" json:"page" form:"page" validate:"gte=0"`
	Size int64 `query:"size" json:"size" form:"size" validate:"gte=0,lte=1000"`

	// Current 当前页码，请优先使用 Page。
	//
	// Deprecated: 兼容前端请求。
	Current int64 `query:"current" json:"current" form:"current" validate:"gte=0"`
}

func (p Pages) PageSize() (int64, int64) {
	page, size := p.Page, p.Size
	if page <= 0 {
		page = p.Current
	}

	return page, size
}

type PageKeywords struct {
	Pages
	Keywords
}

type PageSQL struct {
	Pages
	dynsql.Input
}
