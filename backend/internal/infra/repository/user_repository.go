package repository

import "passkey-practice/backend/internal/domain/user"

type UserRepository interface {
    Save(u *user.User) error
    FindByEmail(email user.Email) (*user.User, error)
    FindByID(id user.ID) (*user.User, error)
}

func (r *UserRepository) SaveUser(id, name string) error {
    // TODO: 実装
    return nil
} 