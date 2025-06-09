package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
}

func NewPasswordFromPlain(plain string) (Password, error) {
	if len(plain) < 8 {
		return Password{}, fmt.Errorf("password must be at least 8 characters")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, fmt.Errorf("failed to hash password: %w", err)
	}
	return Password{hash: string(hashedPassword)}, nil
}

func NewPasswordFromHash(hash string) Password {
	return Password{hash: hash}
}

func (p Password) Hash() string {
	return p.hash
}

func (p Password) Verify(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plain))
	return err == nil
}
