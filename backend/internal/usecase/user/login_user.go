package user

import (
	"context"
	"errors"
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/auth"
	"github.com/Minto312/passkey-practice/backend/internal/domain/session"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

// LoginUserUseCase はユーザーログインのユースケースだよ。
type LoginUserUseCase interface {
	Execute(ctx context.Context, input LoginUserInput) (*LoginUserOutput, error)
}

// LoginUserInput はユーザーログインユースケースの入力だよ。
type LoginUserInput struct {
	Email    string
	Password string
}

// LoginUserOutput はユーザーログインユースケースの出力だよ。
type LoginUserOutput struct {
	AccessToken  string
	RefreshToken string
}

type loginUserInteractor struct {
	userRepo    user.UserRepository
	sessionRepo session.SessionRepository
}

// NewLoginUserInteractor は loginUserInteractor の新しいインスタンスを生成するよ。
func NewLoginUserInteractor(userRepo user.UserRepository, sessionRepo session.SessionRepository) LoginUserUseCase {
	return &loginUserInteractor{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

// Execute はユーザーログインユースケースを実行するよ。
func (i *loginUserInteractor) Execute(ctx context.Context, input LoginUserInput) (*LoginUserOutput, error) {
	// メールアドレスの形式を検証
	email, err := user.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	// メールアドレスでユーザーを検索
	foundUser, err := i.userRepo.FindByEmail(ctx, email)
	if err != nil {
		// FindByEmail でエラーが発生した場合 (DBエラーなど)
		return nil, err
	}
	if foundUser == nil {
		return nil, errors.New("user not found")
	}

	// パスワードを比較
	err = bcrypt.CompareHashAndPassword([]byte(string(foundUser.PasswordHash())), []byte(input.Password))
	if err != nil {
		// パスワードが一致しない場合
		return nil, errors.New("invalid password")
	}

	// トークンを生成
	tokens, err := auth.GenerateTokens(foundUser)
	if err != nil {
		return nil, err
	}

	// セッションを作成
	refreshToken, err := session.NewRefreshToken(tokens.RefreshToken)
	if err != nil {
		return nil, err
	}
	// TODO: IPアドレスとUserAgentをリクエストから取得する
	newSession := session.NewSession(
		foundUser.ID(),
		refreshToken,
		"dummy_ip",
		"dummy_user_agent",
		time.Now().Add(time.Hour*24*30), // 30日間有効
	)

	_, err = i.sessionRepo.Create(ctx, newSession)
	if err != nil {
		return nil, err
	}

	return &LoginUserOutput{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
