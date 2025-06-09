package repository

import (
	"context"
	"github.com/google/uuid"
	entpkg "passkey-practice/backend/ent"
	"passkey-practice/backend/internal/domain/user"
)

type EntUserRepository struct {
	client *entpkg.Client
}

func NewEntUserRepository(client *entpkg.Client) *EntUserRepository {
	return &EntUserRepository{client: client}
}

func (r *EntUserRepository) Save(u *user.User) error {
	ctx := context.Background()
	return r.client.User.Create().
		SetID(u.ID).
		SetName(u.Name).
		SetEmail(u.Email.String()).
		SetPasswordHash(u.Password.Hash()).
		Exec(ctx)
}

func (r *EntUserRepository) FindByEmail(email user.Email) (*user.User, error) {
	ctx := context.Background()
	entUser, err := r.client.User.Query().Where(entpkg.UserEmailEQ(email.String())).Only(ctx)
	if err != nil {
		return nil, err
	}
	return entToDomainUser(entUser)
}

func (r *EntUserRepository) FindByID(id user.ID) (*user.User, error) {
	ctx := context.Background()
	entUser, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entToDomainUser(entUser)
}

func entToDomainUser(entUser *entpkg.User) (*user.User, error) {
	email, err := user.NewEmail(entUser.Email)
	if err != nil {
		return nil, err
	}
	password := user.Password{hash: entUser.PasswordHash}
	return user.NewUser(entUser.ID, entUser.Name, email, password), nil
}
