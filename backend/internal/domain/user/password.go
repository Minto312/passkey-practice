package user

import (
    "fmt"
)

type Password struct {
    hash string
}

// 本来はここでハッシュ化するが、今回はダミー実装
func NewPasswordFromPlain(plain string) (Password, error) {
    if len(plain) < 8 {
        return Password{}, fmt.Errorf("password must be at least 8 characters")
    }
    // TODO: 実際はハッシュ化する
    return Password{hash: "hashed:" + plain}, nil
}

func (p Password) Hash() string {
    return p.hash
} 