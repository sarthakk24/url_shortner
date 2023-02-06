package hashing

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
