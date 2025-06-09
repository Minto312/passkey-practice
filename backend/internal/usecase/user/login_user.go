package user

import (
	"context"
	"errors"
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/auth"
	"github.com/Minto312/passkey-practice/backend/internal/domain/auth_history"
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
	userRepo        user.UserRepository
	sessionRepo     session.SessionRepository
	authHistoryRepo auth_history.AuthHistoryRepository
}

// NewLoginUserInteractor は loginUserInteractor の新しいインスタンスを生成するよ。
func NewLoginUserInteractor(
	userRepo user.UserRepository,
	sessionRepo session.SessionRepository,
	authHistoryRepo auth_history.AuthHistoryRepository,
) LoginUserUseCase {
	return &loginUserInteractor{
		userRepo:        userRepo,
		sessionRepo:     sessionRepo,
		authHistoryRepo: authHistoryRepo,
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
	ipAddress := "dummy_ip"
	userAgent := "dummy_user_agent"

	// セッションを作成
	newSession := session.NewSession(
		foundUser.ID(),
		refreshToken,
		session.IPAddress(ipAddress),
		session.UserAgent(userAgent),
		time.Now().Add(time.Hour*24*30), // 30日間有効
	)
	_, err = i.sessionRepo.Create(ctx, newSession)
	if err != nil {
		return nil, err
	}

	// 認証履歴を記録
	authMethod, _ := auth_history.NewAuthMethod("password") // エラーは発生しないはず
	newAuthHistory := auth_history.NewAuthHistory(
		foundUser.ID(),
		authMethod,
		auth_history.IPAddress(ipAddress),
		auth_history.UserAgent(userAgent),
	)
	_, err = i.authHistoryRepo.Create(ctx, newAuthHistory)
	if err != nil {
		// 認証履歴の記録失敗はログイン自体を失敗させない（ロギング推奨）
		// return nil, err
	}

	return &LoginUserOutput{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
