package util

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(receivedPassword string, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(receivedPassword))
	return err == nil
}

func EncryptPassword(rawPassword string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err == nil {
		return string(encryptedPassword), nil
	} else {
		return "", nil
	}
}
