package auth_history

import (
	"errors"

	"github.com/google/uuid"
)

// AuthHistoryID は認証履歴IDを表す値オブジェクトだよ。
type AuthHistoryID struct {
	uuid.UUID
}

// NewAuthHistoryID は新しい認証履歴IDを生成するよ。
func NewAuthHistoryID() AuthHistoryID {
	return AuthHistoryID{uuid.New()}
}

// ParseAuthHistoryID は文字列から認証履歴IDをパースするよ。
func ParseAuthHistoryID(s string) (AuthHistoryID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return AuthHistoryID{}, err
	}
	return AuthHistoryID{id}, nil
}

// AuthMethod は認証方法を表す値オブジェクトだよ。
type AuthMethod string

const (
	PasswordAuth AuthMethod = "password"
	PasskeyAuth  AuthMethod = "passkey"
)

func NewAuthMethod(method string) (AuthMethod, error) {
	switch AuthMethod(method) {
	case PasswordAuth, PasskeyAuth:
		return AuthMethod(method), nil
	default:
		return "", errors.New("invalid auth method")
	}
}

// IPAddress はIPアドレスを表す値オブジェクトだよ。
type IPAddress string

// UserAgent はユーザーエージェントを表す値オブジェクトだよ。
type UserAgent string
