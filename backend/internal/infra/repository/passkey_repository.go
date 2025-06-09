package repository

import (
	"context"

	"github.com/Minto312/passkey-practice/backend/ent"
	entpasskey "github.com/Minto312/passkey-practice/backend/ent/passkey"
	entuser "github.com/Minto312/passkey-practice/backend/ent/user"
	domainpasskey "github.com/Minto312/passkey-practice/backend/internal/domain/passkey"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

type passkeyRepository struct {
	client *ent.Client
}

// NewPasskeyRepository は passkey.Repository の新しいインスタンスを生成します。
func NewPasskeyRepository(client *ent.Client) domainpasskey.Repository {
	return &passkeyRepository{client: client}
}

// FindByUserID はユーザーIDに紐づくパスキーの一覧を取得します。
func (r *passkeyRepository) FindByUserID(ctx context.Context, userID user.UserID) ([]*domainpasskey.Passkey, error) {
	passkeys, err := r.client.Passkey.
		Query().
		Where(entpasskey.HasUserWith(entuser.ID(userID.UUID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var domainPasskeys []*domainpasskey.Passkey
	for _, p := range passkeys {
		domainPasskeys = append(domainPasskeys, toDomainPasskey(p, userID))
	}

	return domainPasskeys, nil
}

// Save はパスキーを保存します。
func (r *passkeyRepository) Create(ctx context.Context, p *domainpasskey.Passkey) error {
	_, err := r.client.Passkey.
		Create().
		SetID(p.ID().UUID).
		SetUserID(p.UserID().UUID).
		SetCredentialID(p.CredentialID().Value()).
		SetPublicKey(p.PublicKey().Value()).
		SetDeviceName(p.DeviceName().Value()).
		SetCreatedAt(p.CreatedAt()).
		SetLastUsedAt(p.LastUsedAt()).
		Save(ctx)
	return err
}

func toDomainPasskey(p *ent.Passkey, userID user.UserID) *domainpasskey.Passkey {
	return domainpasskey.Reconstruct(
		domainpasskey.PasskeyID{UUID: p.ID},
		userID,
		domainpasskey.CredentialID(p.CredentialID),
		domainpasskey.PublicKey(p.PublicKey),
		domainpasskey.DeviceName(p.DeviceName),
		p.CreatedAt,
		p.LastUsedAt,
	)
}
