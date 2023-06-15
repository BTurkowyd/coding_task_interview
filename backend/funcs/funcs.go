package funcs

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return hashedPassword
}

func ValidatePassword(password string, hashedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}
