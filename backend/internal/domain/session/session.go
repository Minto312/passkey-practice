package session

import (
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

// Session はセッションを表すエンティティだよ。
type Session struct {
	id           SessionID
	userID       user.UserID
	refreshToken RefreshToken
	ipAddress    IPAddress
	userAgent    UserAgent
	expiresAt    time.Time
	createdAt    time.Time
}

// NewSession は新しいセッションエンティティを生成するよ。
func NewSession(
	userID user.UserID,
	refreshToken RefreshToken,
	ipAddress IPAddress,
	userAgent UserAgent,
	expiresAt time.Time,
) *Session {
	return &Session{
		id:           NewSessionID(),
		userID:       userID,
		refreshToken: refreshToken,
		ipAddress:    ipAddress,
		userAgent:    userAgent,
		expiresAt:    expiresAt,
		createdAt:    time.Now(),
	}
}

// FromRepository はリポジトリから取得した情報をもとにセッションエンティティを再構築するよ。
func FromRepository(
	id SessionID,
	userID user.UserID,
	refreshToken RefreshToken,
	ipAddress IPAddress,
	userAgent UserAgent,
	expiresAt time.Time,
	createdAt time.Time,
) *Session {
	return &Session{
		id:           id,
		userID:       userID,
		refreshToken: refreshToken,
		ipAddress:    ipAddress,
		userAgent:    userAgent,
		expiresAt:    expiresAt,
		createdAt:    createdAt,
	}
}

func (s *Session) ID() SessionID {
	return s.id
}

func (s *Session) UserID() user.UserID {
	return s.userID
}

func (s *Session) RefreshToken() RefreshToken {
	return s.refreshToken
}

func (s *Session) IPAddress() IPAddress {
	return s.ipAddress
}

func (s *Session) UserAgent() UserAgent {
	return s.userAgent
}

func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}
