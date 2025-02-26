package dyncond

var (
	Eq         = NewOperator("eq", "=")
	Neq        = NewOperator("neq", "≠")
	Gt         = NewOperator("gt", ">")
	Gte        = NewOperator("gte", "≥")
	Lt         = NewOperator("lt", "<")
	Lte        = NewOperator("lte", "≤")
	Like       = NewOperator("like", "LIKE")
	NotLike    = NewOperator("notlike", "NOT LIKE")
	Between    = NewOperator("between", "BETWEEN")
	NotBetween = NewOperator("notbetween", "NOT BETWEEN")
	In         = NewOperator("in", "IN")
	NotIn      = NewOperator("notin", "NOT IN")
	NotNull    = NewOperator("notnull", "NOT NULL")
	// Regex      = NewOperator("regex", "REGEX") // OpenGauss 不支持
	// NotRegex   = NewOperator("notregex", "NOT REGEX") // OpenGauss 不支持
)

var operators = map[string]Operator{
	Eq.key:         Eq,
	Neq.key:        Neq,
	Gt.key:         Gt,
	Gte.key:        Gte,
	Lt.key:         Lt,
	Lte.key:        Lte,
	Like.key:       Like,
	NotLike.key:    NotLike,
	Between.key:    Between,
	NotBetween.key: NotBetween,
	In.key:         In,
	NotIn.key:      NotIn,
	NotNull.key:    NotNull,
}

func NewOperator(key, desc string) Operator {
	return Operator{
		key:  key,
		desc: desc,
	}
}

type Operator struct {
	key  string
	desc string
}

func (o Operator) Info() (string, string) {
	return o.key, o.desc
}

type Operators []Operator

func LookupOperator(key string) (Operator, bool) {
	op, exist := operators[key]
	return op, exist
}
