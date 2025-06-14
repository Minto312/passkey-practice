package user

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidDisplayName = errors.New("display name is empty")
)

// UserID はユーザーIDを表す値オブジェクトです。
type UserID struct {
	uuid.UUID
}

// NewUserID は新しいユーザーIDを生成します。
func NewUserID() UserID {
	return UserID{uuid.New()}
}

// ParseUserID は文字列からユーザーIDをパースします。
func ParseUserID(s string) (UserID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UserID{}, err
	}
	return UserID{id}, nil
}

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
