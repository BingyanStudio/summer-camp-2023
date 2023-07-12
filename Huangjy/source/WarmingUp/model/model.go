package model

type UserInfo struct {
	ID    uint32 `json:"id" bson:"id" form:"-"`
	Name  string `json:"name" bson:"name" form:"name"`
	Pwd   string `json:"-" bson:"pwd" form:"pwd"`
	Tel   string `json:"tel" bson:"tel" form:"tel"`
	Email string `json:"email" bson:"email" form:"email"`
	Role  uint8  `json:"-" bson:"role" form:"-"`
}

const (
	User uint8 = iota + 1
	Admin
)

type UserUpdate struct {
	Info  string `json:"info" bson:"info" form:"info"`
	Value string `json:"value" bson:"value" form:"value"`
}

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
