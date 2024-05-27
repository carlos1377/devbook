package authentication

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(strPassword, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(strPassword))
}
