package hash

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) []byte {
	password := []byte(pass)
	hashpassword, _ := bcrypt.GenerateFromPassword(password, 10)
	return hashpassword
}

func CompareHashValue(pass string, hashPass string) error {
	password := []byte(pass)
	hashPassword := []byte(hashPass)
	error := bcrypt.CompareHashAndPassword(hashPassword, password)
	return error

}
