package auth_history

import (
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

// AuthHistory は認証履歴を表すエンティティだよ。
type AuthHistory struct {
	id              AuthHistoryID
	userID          user.UserID
	method          AuthMethod
	ipAddress       IPAddress
	userAgent       UserAgent
	authenticatedAt time.Time
}

// NewAuthHistory は新しい認証履歴エンティティを生成するよ。
func NewAuthHistory(
	userID user.UserID,
	method AuthMethod,
	ipAddress IPAddress,
	userAgent UserAgent,
) *AuthHistory {
	return &AuthHistory{
		id:              NewAuthHistoryID(),
		userID:          userID,
		method:          method,
		ipAddress:       ipAddress,
		userAgent:       userAgent,
		authenticatedAt: time.Now(),
	}
}

// FromRepository はリポジトリから取得した情報をもとに認証履歴エンティティを再構築するよ。
func FromRepository(
	id AuthHistoryID,
	userID user.UserID,
	method AuthMethod,
	ipAddress IPAddress,
	userAgent UserAgent,
	authenticatedAt time.Time,
) *AuthHistory {
	return &AuthHistory{
		id:              id,
		userID:          userID,
		method:          method,
		ipAddress:       ipAddress,
		userAgent:       userAgent,
		authenticatedAt: authenticatedAt,
	}
}

func (h *AuthHistory) ID() AuthHistoryID {
	return h.id
}

func (h *AuthHistory) UserID() user.UserID {
	return h.userID
}

func (h *AuthHistory) Method() AuthMethod {
	return h.method
}

func (h *AuthHistory) IPAddress() IPAddress {
	return h.ipAddress
}

func (h *AuthHistory) UserAgent() UserAgent {
	return h.userAgent
}

func (h *AuthHistory) AuthenticatedAt() time.Time {
	return h.authenticatedAt
}
