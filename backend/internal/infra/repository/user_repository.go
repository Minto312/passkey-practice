package repository

import (
	"context"

	"github.com/Minto312/passkey-practice/backend/ent"
	ent_user "github.com/Minto312/passkey-practice/backend/ent/user"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

type userRepository struct {
	client *ent.Client
}

// NewUserRepository は user.UserRepository の新しいインスタンスを生成します。
func NewUserRepository(client *ent.Client) user.UserRepository {
	return &userRepository{client: client}
}

// Create は新しいユーザーを永続化します。
func (r *userRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	// u.ID() はリポジトリ層で採番されるため、ドメイン層で生成されたIDは無視する
	created, err := r.client.User.
		Create().
		SetEmail(string(u.Email())).
		SetPassword(string(u.PasswordHash())).
		SetDisplayName(string(u.DisplayName())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainUser(created), nil
}

// FindByEmail はメールアドレスでユーザーを検索します。
func (r *userRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	found, err := r.client.User.
		Query().
		Where(ent_user.EmailEQ(string(email))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainUser(found), nil
}

// FindByID はIDでユーザーを検索します。
func (r *userRepository) FindByID(ctx context.Context, id user.UserID) (*user.User, error) {
	found, err := r.client.User.
		Query().
		Where(ent_user.IDEQ(id.UUID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainUser(found), nil
}

func toDomainUser(e *ent.User) *user.User {
	userID, _ := user.ParseUserID(e.ID.String())
	email, _ := user.NewEmail(e.Email)
	passwordHash, _ := user.NewPasswordHash(e.Password)
	displayName, _ := user.NewDisplayName(e.DisplayName)

	return user.Reconstruct(
		userID,
		email,
		passwordHash,
		displayName,
		e.CreatedAt,
		e.UpdatedAt,
	)
}
