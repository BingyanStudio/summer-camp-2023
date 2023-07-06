package model

type User struct {
	ID    uint32 `json:"id" bson:"id" form:"-"`
	Name  string `json:"name" bson:"name" form:"name"`
	Pwd   string `json:"-" bson:"pwd" form:"pwd"`
	Tel   string `json:"tel" bson:"tel" form:"tel"`
	Email string `json:"email" bson:"email" form:"email"`
	Role  uint8  `json:"role" bson:"role" form:"-"`
}

type UserUpdate struct {
	Info  string      `json:"info" bson:"info"`
	Value interface{} `json:"value" bson:"value"`
}
