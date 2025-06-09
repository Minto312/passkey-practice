package auth

import (
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TODO: 環境変数から読み込むようにする
var jwtSecret = []byte("super-secret-key")

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// GenerateTokens は新しいアクセストークンとリフレッシュトークンを生成するよ。
func GenerateTokens(u *user.User) (*Tokens, error) {
	// アクセストークンの生成
	accessTokenClaims := jwt.MapClaims{
		"sub":  u.ID().String(),
		"name": string(u.DisplayName()),
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // 1時間有効
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// リフレッシュトークンの生成 (ここでは単純なUUIDを使用)
	refreshTokenString := uuid.New().String()

	return &Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
