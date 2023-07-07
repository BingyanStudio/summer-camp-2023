package utils

import (
	"errors"
	"fmt"
	"log"
	db "myserver/database"
	"myserver/model"
	"reflect"

	jwt "github.com/golang-jwt/jwt/v5"
)

// 检查注册信息完整和合法性
func CheckRegisterValid(u *model.UserInfo) error {
	if u == nil {
		return errors.New("nil pointer")
	}
	// 验证信息是否会空
	rVal := reflect.ValueOf(*u)
	rType := reflect.TypeOf(*u)
	for i, n := 0, rType.NumField(); i < n; i++ {
		if rVal.Field(i).Kind() == reflect.String && rVal.Field(i).String() == "" {
			return errors.New(fmt.Sprintf("%s could not be empty", rType.Field(i).Name))
		}
	}
	// 然后验证邮箱和电话号码的合法。。。
	return nil
}

func GenerateJWTToken(user_id uint32, role uint8) string {
	jwtKey := db.QueryOtherInfo("jwtKey")
	if jwtKey == nil {
		return ""
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user_id,
		"role":   role,
	})
	tokenString, err := token.SignedString([]byte(jwtKey.(string)))
	if err != nil {
		log.Println(err)
		return ""
	}
	return tokenString
}
