package user

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidDisplayName = errors.New("display name is empty")
)

// UserID はユーザーIDを表す値オブジェクトです。
type UserID int

// Email はメールアドレスを表す値オブジェクトです。
type Email string

func NewEmail(email string) (Email, error) {
	// 簡単なメールアドレスのバリデーション
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return "", ErrInvalidEmail
	}
	return Email(email), nil
}

// PasswordHash はパスワードハッシュを表す値オブジェクトです。
type PasswordHash string

func NewPasswordHash(hash string) (PasswordHash, error) {
	if hash == "" {
		return "", errors.New("password hash is empty")
	}
	return PasswordHash(hash), nil
}

// DisplayName は表示名を表す値オブジェクトです。
type DisplayName string

func NewDisplayName(name string) (DisplayName, error) {
	if name == "" {
		return "", ErrInvalidDisplayName
	}
	return DisplayName(name), nil
}
