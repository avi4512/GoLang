package hashing

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	byte_data, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(byte_data), err

}
