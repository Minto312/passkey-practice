package session

import "context"

// SessionRepository はセッションの永続化を担当するリポジトリだよ。
type SessionRepository interface {
	Create(ctx context.Context, session *Session) (*Session, error)
}
