package entity

type Article struct {
	Articleid string   `json:"articleid"`
	Type      []string `json:"type"`
	Title     string   `json:"title"`
	Userid    string   `json:"userid"`
	Username  string   `json:"username"`
	Time      string   `json:"time"`
	Readnum   int      `json:"readnum"`
	Content   string   `json:"content"`
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
	Field string `json:"field"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
