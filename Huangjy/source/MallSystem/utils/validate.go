package utils

import (
	"MallSystem/model"
	"errors"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidateRegisterInfo(u *model.UserInfo) error {
	var (
		passwordPattern string = "^[a-zA-Z0-9_-]{6,16}$"
		mobilePattern   string = "^(1[3-9][0-9])\\d{8}$"
		emailPattern    string = "^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$"
	)
	var (
		passwordReg = regexp.MustCompile(passwordPattern)
		mobileReg   = regexp.MustCompile(mobilePattern)
		emailReg    = regexp.MustCompile(emailPattern)
	)
	if ok := passwordReg.MatchString(u.Password); !ok {
		return errors.New("密码只能包含英文大小写字母和数字，不得少于6个或大于16个字符")
	}
	if ok := mobileReg.MatchString(u.Mobile); !ok {
		return errors.New("电话号码不正确")
	}
	if ok := emailReg.MatchString(u.Email); !ok {
		return errors.New("电子邮箱不正确")
	}
	return nil
}

func ComparePassword(l *model.Login, u *model.UserInfo) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
	if err != nil {
		return err
	}
	return nil
}

func EncryptUserPassword(s *string) error {
	hashed_pwd, err := bcrypt.GenerateFromPassword([]byte(*s), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	*s = string(hashed_pwd)
	return nil
}

func ValidateCommodityInfo(c *model.CommodityInfo) error {
	return nil
}
