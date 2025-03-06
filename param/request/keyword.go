package request

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

type Keywords struct {
	Keyword string `json:"keyword" query:"keyword" form:"keyword" validate:"max=255"`
}

func (k Keywords) Likes(fields ...field.String) []field.Expr {
	size := len(fields)
	if size == 0 {
		return nil
	}
	kw := k.Format()
	if kw == "" {
		return nil
	}

	likes := make([]field.Expr, 0, len(fields))
	for _, f := range fields {
		likes = append(likes, f.Like(kw))
	}

	return likes
}

func (k Keywords) LikeScopes(or bool, fields ...field.String) func(gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		likes := k.Likes(fields...)
		if len(likes) == 0 {
			return dao
		}
		if or {
			return dao.Where(field.Or(likes...))
		}

		return dao.Where(field.And(likes...))
	}
}

func (k Keywords) Format() string {
	if kw := strings.TrimSpace(k.Keyword); kw != "" {
		return "%" + kw + "%"
	}

	return ""
}
