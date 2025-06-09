package auth_history

import "context"

// AuthHistoryRepository は認証履歴の永続化を担当するリポジトリだよ。
type AuthHistoryRepository interface {
	Save(ctx context.Context, authHistory *AuthHistory) error
	FindAll(ctx context.Context) ([]*AuthHistory, error)
}
