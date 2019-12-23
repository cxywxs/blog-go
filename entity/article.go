package entity

type Article struct {
	Articleid string
	Type1     string
	Type2     string
	Title     string
	Userid    string
	Username  string
	Time      string
	Readnum   int
	Content   string
}

const (
	Equaled      = "$eq"
	Unequal      = "$ne"
	Greater      = "$gt"
	Lesser       = "$lt"
	GreaterEqual = "$gte"
	LesserEqual  = "$lte"
	Innet        = "$in"
	Anded        = "$and"
	Ored         = "$or"
)

type SelectCondition struct {
	Field string
	Key   string
	Value string
}
