package util

import (
	"fmt"
	//"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Phone    string `json:"phone"`
	Password string
	jwt.StandardClaims
}

func GenJWT(phone string, password string) string {
	mySigningKey := []byte("keepongoing")
	c := MyClaims{
		Phone:    phone,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*2,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(t)
	s, e := t.SignedString(mySigningKey)
	if e != nil {
		fmt.Println(e)

	}
	fmt.Println(s)
	return s
}

// 解密
func ParseJWT(s string) jwt.Claims {
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte("keepongoing"), nil })
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("结果", token.Claims)
	//断言
	fmt.Println("断言", token.Claims.(*MyClaims))

	res := token.Claims
	return res
}
