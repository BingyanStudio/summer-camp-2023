package utils

import (
	"errors"
	"fmt"
	"myserver/model"
	"reflect"
)

// 检查注册信息完整和合法性
func CheckRegisterValid(u *model.User) error {
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
