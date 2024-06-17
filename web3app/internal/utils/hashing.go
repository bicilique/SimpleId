package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashAndEncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}

func CheckPasswordHash(password, encodedHash string) (bool, error) {
	hash, err := base64.StdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil, err
}
