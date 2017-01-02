package encryption

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4);
	if (err != nil) {
		return []byte(""), err
	}
	return hashedPassword, err
}

func Compare(hashedPassword []byte, password []byte) (error) {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}