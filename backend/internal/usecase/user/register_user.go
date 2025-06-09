package user

import (
	"context"
	"errors"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUserUseCase はユーザー登録のユースケースです。
type RegisterUserUseCase interface {
	Execute(ctx context.Context, input RegisterUserInput) (*RegisterUserOutput, error)
}

// RegisterUserInput はユーザー登録ユースケースの入力です。
type RegisterUserInput struct {
	Email       string
	Password    string
	DisplayName string
}

// RegisterUserOutput はユーザー登録ユースケースの出力です。
type RegisterUserOutput struct {
	User *user.User
}

type registerUserInteractor struct {
	userRepo user.UserRepository
}

// NewRegisterUserInteractor は registerUserInteractor の新しいインスタンスを生成します。
func NewRegisterUserInteractor(userRepo user.UserRepository) RegisterUserUseCase {
	return &registerUserInteractor{
		userRepo: userRepo,
	}
}

// Execute はユーザー登録ユースケースを実行します。
func (i *registerUserInteractor) Execute(ctx context.Context, input RegisterUserInput) (*RegisterUserOutput, error) {
	// メールアドレスの形式を検証
	email, err := user.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	// 表示名の形式を検証
	displayName, err := user.NewDisplayName(input.DisplayName)
	if err != nil {
		return nil, err
	}

	// メールアドレスの重複チェック
	existingUser, err := i.userRepo.FindByEmail(ctx, email)
	if err != nil {
		// FindByEmail でエラーが発生した場合 (DBエラーなど)
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	passwordHash, err := user.NewPasswordHash(string(hashedPassword))
	if err != nil {
		return nil, err
	}

	// 新しいユーザーを作成
	newUser := user.NewUser(email, passwordHash, displayName)

	// ユーザーを永続化
	createdUser, err := i.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &RegisterUserOutput{User: createdUser}, nil
}
