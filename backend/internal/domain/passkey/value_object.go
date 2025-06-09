package passkey

import "github.com/google/uuid"

// PasskeyID はパスキーIDを表す値オブジェクトです。
type PasskeyID struct {
	uuid.UUID
}

// NewPasskeyID は新しいパスキーIDを生成します。
func NewPasskeyID() PasskeyID {
	return PasskeyID{uuid.New()}
}

// ParsePasskeyID は文字列からパスキーIDをパースします。
func ParsePasskeyID(s string) (PasskeyID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return PasskeyID{}, err
	}
	return PasskeyID{id}, nil
}

// CredentialID はクレデンシャルIDを表す値オブジェクトです。
type CredentialID string

// Value はクレデンシャルIDの文字列表現を返します。
func (c CredentialID) Value() string {
	return string(c)
}

// PublicKey は公開鍵を表す値オブジェクトです。
type PublicKey []byte

// Value は公開鍵のバイト配列を返します。
func (p PublicKey) Value() []byte {
	return p
}

// DeviceName はデバイス名を表す値オブジェクトです。
type DeviceName string

// Value はデバイス名の文字列表現を返します。
func (d DeviceName) Value() string {
	return string(d)
}
