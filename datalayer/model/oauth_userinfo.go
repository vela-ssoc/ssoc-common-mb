package model

type OAuthDongUserinfo struct {
	SUB                string   `json:"sub"`                // 工号（不带前缀，即证券用户为"17290"而非"Z17290"）
	CompanyPrefix      string   `json:"companyPrefix"`      // 公司：G-集团，Z-证券，T-天天基金 ，F-期货 ，D-东财基金
	Gender             string   `json:"gender"`             // 性别：F-女，M-男
	CompanyName        string   `json:"companyName"`        // 公司名称
	UniqueKey          string   `json:"uniqueKey"`          // 唯一ID
	ISS                string   `json:"iss"`                // 签发者
	Nonce              string   `json:"nonce"`              // 跳转登录系统前端时如果携带了 nonce，这里就会原样传回
	Picture            string   `json:"picture"`            // 头像链接
	AUD                []string `json:"aud"`                // client_id
	Name               string   `json:"name"`               // 姓名
	DepartmentFullName string   `json:"departmentFullName"` // 部门：运维中心/自动化运维部
	EXP                float64  `json:"exp"`                // 过期时间 1695698300.000000000
	IAT                float64  `json:"iat"`                // 签发时间 1695694700.000000000,
	Email              string   `json:"email"`              // 邮箱：暂不提供真实邮箱 a0669gcb5wem@em.com
	CompanyCode        string   `json:"companyCode"`        // 公司code: 集团-group ,证券-security ,天天基金-ttfund ,期货-futures , 东财基金-dcfund
}
