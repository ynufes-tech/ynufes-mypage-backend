package cipher

import "golang.org/x/crypto/bcrypt"

func GenHash(target string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(target), bcrypt.DefaultCost)
}

func Compare(target string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(target))
}
