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

const Equaled = "$eq"
const Unequal = "$ne"
const Greater = "$gt"
const Lesser = "$lt"
const GreaterEqual = "$gte"
const LesserEqual = "$lte"
const Innet = "$in"
const Anded = "$and"
const Ored = "$or"

type SelectCondition struct {
	Field string
	Key   string
	Value string
}
