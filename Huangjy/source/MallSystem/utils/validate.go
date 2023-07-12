package utils

import (
	"MallSystem/model"
	_ "errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ValidateRegisterInfo(u *model.UserInfo) error {
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
