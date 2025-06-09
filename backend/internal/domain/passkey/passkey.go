package passkey

import (
	"context"
	"log"
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
	"github.com/go-webauthn/webauthn/webauthn"
)

// Passkey はパスキーを表すエンティティです。
type Passkey struct {
	id           PasskeyID
	userID       user.UserID
	credentialID CredentialID
	publicKey    PublicKey
	deviceName   DeviceName
	createdAt    time.Time
	lastUsedAt   time.Time
}

// NewPasskey は新しいパスキーを作成します。
func NewPasskey(
	userID user.UserID,
	credentialID CredentialID,
	publicKey PublicKey,
	deviceName DeviceName,
) *Passkey {
	now := time.Now()
	return &Passkey{
		userID:       userID,
		credentialID: credentialID,
		publicKey:    publicKey,
		deviceName:   deviceName,
		createdAt:    now,
		lastUsedAt:   now,
	}
}

// Reconstruct は永続化層からパスキーを再構築します。
func Reconstruct(
	id PasskeyID,
	userID user.UserID,
	credentialID CredentialID,
	publicKey PublicKey,
	deviceName DeviceName,
	createdAt time.Time,
	lastUsedAt time.Time,
) *Passkey {
	return &Passkey{
		id:           id,
		userID:       userID,
		credentialID: credentialID,
		publicKey:    publicKey,
		deviceName:   deviceName,
		createdAt:    createdAt,
		lastUsedAt:   lastUsedAt,
	}
}

func BeginAssertion(ctx context.Context) (*webauthn.WebAuthn, error) {
	
wconfig := &webauthn.Config{
	RPDisplayName: "Passkey Practice", // Display Name for your site
	RPID: "passkey-practice.local", // Generally the FQDN for your site
	RPOrigins: []string{"http://localhost:8080"}, // The origin URLs allowed for WebAuthn requests
}

webAuthn, err := webauthn.New(wconfig)
if err != nil {
	log.Fatalf("failed to initialize webauthn: %v", err)
	return nil, err
}

return webAuthn, nil
}

// ID はパスキーIDを返します。
func (p *Passkey) ID() PasskeyID {
	return p.id
}

// UserID はユーザーIDを返します。
func (p *Passkey) UserID() user.UserID {
	return p.userID
}

// CredentialID はクレデンシャルIDを返します。
func (p *Passkey) CredentialID() CredentialID {
	return p.credentialID
}

// PublicKey は公開鍵を返します。
func (p *Passkey) PublicKey() PublicKey {
	return p.publicKey
}

// DeviceName はデバイス名を返します。
func (p *Passkey) DeviceName() DeviceName {
	return p.deviceName
}

// CreatedAt は作成日時を返します。
func (p *Passkey) CreatedAt() time.Time {
	return p.createdAt
}

// LastUsedAt は最終使用日時を返します。
func (p *Passkey) LastUsedAt() time.Time {
	return p.lastUsedAt
}
