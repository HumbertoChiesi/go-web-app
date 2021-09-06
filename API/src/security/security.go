package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashPassword, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(passwordString))
}
