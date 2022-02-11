package services

import "golang.org/x/crypto/bcrypt"

func Hash(input string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(password), err
}

func Compare(hashed string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	return err != nil
}
