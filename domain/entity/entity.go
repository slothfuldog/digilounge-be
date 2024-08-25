package entity

type UserInf struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Sts        string `json:"sts"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Birthplace string `json:"birthplace"`
	Regdate    string `json:"regdate"`
	Regtime    string `json:"regtime"`
}
