package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(plainPwd string) (string, error) {
	fmt.Println(plainPwd)
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
