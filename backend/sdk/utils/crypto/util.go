package crypto

import "golang.org/x/crypto/bcrypt"

type Crypto struct {
	Cost int
}

func (c *Crypto) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c.Cost)
	return string(bytes), err
}

func (c *Crypto) CompareAndHash(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
