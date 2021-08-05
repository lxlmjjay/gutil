package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordCrypt(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("生成密码错误", err)
	}
	return string(bytes)
}

//密码验证
func PasswordCompare(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err == nil {
		return true
	}
	return false
}
