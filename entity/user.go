package entity

type IrisUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Tokenstring string `json:"tokenstring"`
}
