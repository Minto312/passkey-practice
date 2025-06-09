package repository

import (
	"context"

	entpkg "github.com/Minto312/passkey-practice/backend/ent"
	entuser "github.com/Minto312/passkey-practice/backend/ent/user"                // Entのuserパッケージ
	domainuser "github.com/Minto312/passkey-practice/backend/internal/domain/user" // ドメインのuserパッケージ
	"github.com/google/uuid"
)

type EntUserRepository struct {
	client *entpkg.Client
}

func NewEntUserRepository(client *entpkg.Client) *EntUserRepository {
	return &EntUserRepository{client: client}
}

func (r *EntUserRepository) Save(u *domainuser.User) error {
	ctx := context.Background()

	if u.ID == uuid.Nil { // 新規作成
		_, err := r.client.User.Create().
			SetName(u.Name).
			SetEmail(u.Email.String()).
			SetPasswordHash(u.Password.Hash()).
			Save(ctx)
		return err
	} else { // 更新
		_, err := r.client.User.UpdateOneID(u.ID).
			SetName(u.Name).
			SetEmail(u.Email.String()).
			SetPasswordHash(u.Password.Hash()).
			Save(ctx)
		return err
	}
}

func (r *EntUserRepository) FindByEmail(email domainuser.Email) (*domainuser.User, error) {
	ctx := context.Background()
	entUser, err := r.client.User.Query().Where(entuser.EmailEQ(email.String())).Only(ctx)
	if err != nil {
		return nil, err
	}
	return entToDomainUser(entUser)
}

func (r *EntUserRepository) FindByID(id domainuser.ID) (*domainuser.User, error) {
	ctx := context.Background()
	entUser, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entToDomainUser(entUser)
}

func entToDomainUser(entUser *entpkg.User) (*domainuser.User, error) {
	email, err := domainuser.NewEmail(entUser.Email)
	if err != nil {
		return nil, err
	}
	password := domainuser.NewPasswordFromHash(entUser.PasswordHash)
	return domainuser.NewUser(entUser.ID, entUser.Name, email, password), nil
}
