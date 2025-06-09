package user

import "github.com/google/uuid"

// ID型
type ID = uuid.UUID

// Userエンティティ
type User struct {
    ID       ID
    Name     string
    Email    Email
    Password Password
}

// Userのコンストラクタ
func NewUser(id ID, name string, email Email, password Password) *User {
    return &User{
        ID:       id,
        Name:     name,
        Email:    email,
        Password: password,
    }
} 