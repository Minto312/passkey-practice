package session

import (
	"errors"

	"github.com/google/uuid"
)

// SessionID はセッションIDを表す値オブジェクトだよ。
type SessionID struct {
	uuid.UUID
}

// NewSessionID は新しいセッションIDを生成するよ。
func NewSessionID() SessionID {
	return SessionID{uuid.New()}
}

// ParseSessionID は文字列からセッションIDをパースするよ。
func ParseSessionID(s string) (SessionID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return SessionID{}, err
	}
	return SessionID{id}, nil
}

// RefreshToken はリフレッシュトークンを表す値オブジェクトだよ。
type RefreshToken string

func NewRefreshToken(token string) (RefreshToken, error) {
	// ここでは簡単なバリデーションのみ
	if token == "" {
		return "", errors.New("refresh token is empty")
	}
	return RefreshToken(token), nil
}

// IPAddress はIPアドレスを表す値オブジェクトだよ。
type IPAddress string

// UserAgent はユーザーエージェントを表す値オブジェクトだよ。
type UserAgent string
