package auth_history

import "errors"

var (
	ErrInvalidAuthMethod = errors.New("invalid auth method")
)

// AuthHistoryID は履歴IDを表す値オブジェクトです。
type AuthHistoryID int

// AuthMethod は認証方式を表す値オブジェクトです。
type AuthMethod string

const (
	AuthMethodPassword AuthMethod = "password"
	AuthMethodPasskey  AuthMethod = "passkey"
)

func NewAuthMethod(method string) (AuthMethod, error) {
	switch AuthMethod(method) {
	case AuthMethodPassword, AuthMethodPasskey:
		return AuthMethod(method), nil
	default:
		return "", ErrInvalidAuthMethod
	}
}

// IPAddress はIPアドレスを表す値オブジェクトです。
type IPAddress string

// UserAgent はユーザーエージェントを表す値オブジェクトです。
type UserAgent string
