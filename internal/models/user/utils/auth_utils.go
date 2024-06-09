package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ValidatePassword(hashPass, inPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(inPass))
	if err != nil {
		return fmt.Errorf("access denied: %v", err.Error())
	}

	return nil
}
