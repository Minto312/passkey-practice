package repository

import "passkey-practice/backend/internal/domain/user"

type UserRepository interface {
	Save(u *user.User) error
	FindByEmail(email user.Email) (*user.User, error)
	FindByID(id user.ID) (*user.User, error)
}
