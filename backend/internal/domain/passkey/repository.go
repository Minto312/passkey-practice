package passkey

import (
	"context"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

// Repository はパスキーの永続化を担当するリポジトリです。
type Repository interface {
	// FindByUserID はユーザーIDに紐づくパスキーの一覧を取得します。
	FindByUserID(ctx context.Context, userID user.UserID) ([]*Passkey, error)
	// Save はパスキーを保存します。
	Create(ctx context.Context, passkey *Passkey) error
}
