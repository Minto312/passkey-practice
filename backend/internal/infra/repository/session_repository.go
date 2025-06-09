package repository

import (
	"context"
	"fmt"

	"github.com/Minto312/passkey-practice/backend/ent"
	ent_session "github.com/Minto312/passkey-practice/backend/ent/session"
	"github.com/Minto312/passkey-practice/backend/internal/domain/session"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

type sessionRepository struct {
	client *ent.Client
}

// NewSessionRepository は sessionRepository の新しいインスタンスを生成します。
func NewSessionRepository(client *ent.Client) session.SessionRepository {
	return &sessionRepository{
		client: client,
	}
}

func (r *sessionRepository) Create(ctx context.Context, s *session.Session) (*session.Session, error) {
	created, err := r.client.Session.
		Create().
		SetID(s.ID().UUID).
		SetUserID(s.UserID().UUID).
		SetRefreshToken(string(s.RefreshToken())).
		SetIPAddress(string(s.IPAddress())).
		SetUserAgent(string(s.UserAgent())).
		SetExpiresAt(s.ExpiresAt()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// 再度取得してEdgeを読み込む
	savedSession, err := r.client.Session.Query().Where(ent_session.IDEQ(created.ID)).WithUser().Only(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainSession(savedSession)
}

func toDomainSession(e *ent.Session) (*session.Session, error) {
	id, err := session.ParseSessionID(e.ID.String())
	if err != nil {
		return nil, err
	}
	if e.Edges.User == nil {
		return nil, fmt.Errorf("edge user not found for session %s", e.ID.String())
	}
	userID, err := user.ParseUserID(e.Edges.User.ID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := session.NewRefreshToken(e.RefreshToken)
	if err != nil {
		return nil, err
	}

	return session.FromRepository(
		id,
		userID,
		refreshToken,
		session.IPAddress(e.IPAddress),
		session.UserAgent(e.UserAgent),
		e.ExpiresAt,
		e.CreatedAt,
	), nil
}
