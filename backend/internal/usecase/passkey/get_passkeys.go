package passkey

import (
	"context"

	"github.com/Minto312/passkey-practice/backend/internal/domain/passkey"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

// GetPasskeysUseCase はパスキー一覧取得のユースケースです。
type GetPasskeysUseCase struct {
	passkeyRepo passkey.Repository
}

// NewGetPasskeysUseCase は GetPasskeysUseCase の新しいインスタンスを生成します。
func NewGetPasskeysUseCase(passkeyRepo passkey.Repository) *GetPasskeysUseCase {
	return &GetPasskeysUseCase{passkeyRepo: passkeyRepo}
}

// Execute はパスキー一覧取得を実行します。
func (uc *GetPasskeysUseCase) Execute(ctx context.Context, userID user.UserID) ([]*passkey.Passkey, error) {
	return uc.passkeyRepo.FindByUserID(ctx, userID)
}
