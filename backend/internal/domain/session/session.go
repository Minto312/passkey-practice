package session

import (
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

// Session はセッションを表すエンティティです。
type Session struct {
	id           SessionID
	userID       user.UserID
	createdAt    time.Time
	expiresAt    time.Time
	refreshToken RefreshToken
	ipAddress    IPAddress
	userAgent    UserAgent
}

// NewSession は新しいセッションを作成します。
func NewSession(
	id SessionID,
	userID user.UserID,
	expiresAt time.Time,
	refreshToken RefreshToken,
	ipAddress IPAddress,
	userAgent UserAgent,
) *Session {
	return &Session{
		id:           id,
		userID:       userID,
		createdAt:    time.Now(),
		expiresAt:    expiresAt,
		refreshToken: refreshToken,
		ipAddress:    ipAddress,
		userAgent:    userAgent,
	}
}

// Reconstruct は永続化層からセッションを再構築します。
func Reconstruct(
	id SessionID,
	userID user.UserID,
	createdAt time.Time,
	expiresAt time.Time,
	refreshToken RefreshToken,
	ipAddress IPAddress,
	userAgent UserAgent,
) *Session {
	return &Session{
		id:           id,
		userID:       userID,
		createdAt:    createdAt,
		expiresAt:    expiresAt,
		refreshToken: refreshToken,
		ipAddress:    ipAddress,
		userAgent:    userAgent,
	}
}

// ID はセッションIDを返します。
func (s *Session) ID() SessionID {
	return s.id
}

// UserID はユーザーIDを返します。
func (s *Session) UserID() user.UserID {
	return s.userID
}

// CreatedAt は作成日時を返します。
func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}

// ExpiresAt は有効期限を返します。
func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

// RefreshToken はリフレッシュトークンを返します。
func (s *Session) RefreshToken() RefreshToken {
	return s.refreshToken
}

// IPAddress はIPアドレスを返します。
func (s *Session) IPAddress() IPAddress {
	return s.ipAddress
}

// UserAgent はユーザーエージェントを返します。
func (s *Session) UserAgent() UserAgent {
	return s.userAgent
}
