package session

import (
	"sync"

	"github.com/go-webauthn/webauthn/webauthn"
)

// Store はWebAuthnのセッションデータをインメモリで管理します。
// 本番環境ではRedisなどを使うべきです。
type Store struct {
	sync.RWMutex
	sessions map[string]*webauthn.SessionData
}

// NewStore は新しいStoreを生成します。
func NewStore() *Store {
	return &Store{
		sessions: make(map[string]*webauthn.SessionData),
	}
}

// Save はセッションデータを保存します。
func (s *Store) Save(session *webauthn.SessionData) {
	s.Lock()
	defer s.Unlock()
	// 簡単のため、チャレンジをキーにする
	s.sessions[string(session.Challenge)] = session
}

// Get はチャレンジに対応するセッションデータを取得します。
func (s *Store) Get(challenge string) (*webauthn.SessionData, bool) {
	s.RLock()
	defer s.RUnlock()
	session, ok := s.sessions[challenge]
	return session, ok
}
