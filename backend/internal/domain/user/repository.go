package user

import "context"

// UserRepository はユーザーの永続化を担当するリポジトリです。
type UserRepository interface {
	// Create は新しいユーザーを永続化します。
	Create(ctx context.Context, user *User) (*User, error)
	// FindByEmail はメールアドレスでユーザーを検索します。
	FindByEmail(ctx context.Context, email Email) (*User, error)
	// FindByID はIDでユーザーを検索します。
	FindByID(ctx context.Context, id UserID) (*User, error)
}
