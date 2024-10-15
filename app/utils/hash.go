package utils

import "golang.org/x/crypto/bcrypt"

var hashCost = bcrypt.DefaultCost

func SetCost(cost int) {
	hashCost = cost
}

func MakeHash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), hashCost)

	return string(bytes), err
}

func CheckHash(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
