package auth_history

import (
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

// AuthHistory は認証履歴を表すエンティティです。
type AuthHistory struct {
	id              AuthHistoryID
	userID          user.UserID
	method          AuthMethod
	authenticatedAt time.Time
	ipAddress       IPAddress
	userAgent       UserAgent
}

// NewAuthHistory は新しい認証履歴を作成します。
func NewAuthHistory(
	userID user.UserID,
	method AuthMethod,
	ipAddress IPAddress,
	userAgent UserAgent,
) *AuthHistory {
	return &AuthHistory{
		userID:          userID,
		method:          method,
		authenticatedAt: time.Now(),
		ipAddress:       ipAddress,
		userAgent:       userAgent,
	}
}

// Reconstruct は永続化層から認証履歴を再構築します。
func Reconstruct(
	id AuthHistoryID,
	userID user.UserID,
	method AuthMethod,
	authenticatedAt time.Time,
	ipAddress IPAddress,
	userAgent UserAgent,
) *AuthHistory {
	return &AuthHistory{
		id:              id,
		userID:          userID,
		method:          method,
		authenticatedAt: authenticatedAt,
		ipAddress:       ipAddress,
		userAgent:       userAgent,
	}
}

// ID は履歴IDを返します。
func (a *AuthHistory) ID() AuthHistoryID {
	return a.id
}

// UserID はユーザーIDを返します。
func (a *AuthHistory) UserID() user.UserID {
	return a.userID
}

// Method は認証方式を返します。
func (a *AuthHistory) Method() AuthMethod {
	return a.method
}

// AuthenticatedAt は認証日時を返します。
func (a *AuthHistory) AuthenticatedAt() time.Time {
	return a.authenticatedAt
}

// IPAddress はIPアドレスを返します。
func (a *AuthHistory) IPAddress() IPAddress {
	return a.ipAddress
}

// UserAgent はユーザーエージェントを返します。
func (a *AuthHistory) UserAgent() UserAgent {
	return a.userAgent
}
