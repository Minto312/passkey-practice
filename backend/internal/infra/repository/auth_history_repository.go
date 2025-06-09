package repository

import (
	"context"
	"fmt"

	"github.com/Minto312/passkey-practice/backend/ent"
	ent_auth_history "github.com/Minto312/passkey-practice/backend/ent/authhistory"
	"github.com/Minto312/passkey-practice/backend/internal/domain/auth_history"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
)

type authHistoryRepository struct {
	client *ent.Client
}

// NewAuthHistoryRepository は authHistoryRepository の新しいインスタンスを生成します。
func NewAuthHistoryRepository(client *ent.Client) auth_history.AuthHistoryRepository {
	return &authHistoryRepository{
		client: client,
	}
}

func (r *authHistoryRepository) Create(ctx context.Context, h *auth_history.AuthHistory) (*auth_history.AuthHistory, error) {
	created, err := r.client.AuthHistory.
		Create().
		SetID(h.ID().UUID).
		SetUserID(h.UserID().UUID).
		SetMethod(string(h.Method())).
		SetIPAddress(string(h.IPAddress())).
		SetUserAgent(string(h.UserAgent())).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	saved, err := r.client.AuthHistory.Query().Where(ent_auth_history.IDEQ(created.ID)).WithUser().Only(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainAuthHistory(saved)
}

func toDomainAuthHistory(e *ent.AuthHistory) (*auth_history.AuthHistory, error) {
	id, err := auth_history.ParseAuthHistoryID(e.ID.String())
	if err != nil {
		return nil, err
	}
	if e.Edges.User == nil {
		return nil, fmt.Errorf("edge user not found for auth_history %s", e.ID.String())
	}
	userID, err := user.ParseUserID(e.Edges.User.ID.String())
	if err != nil {
		return nil, err
	}
	method, err := auth_history.NewAuthMethod(e.Method)
	if err != nil {
		return nil, err
	}

	return auth_history.FromRepository(
		id,
		userID,
		method,
		auth_history.IPAddress(e.IPAddress),
		auth_history.UserAgent(e.UserAgent),
		e.AuthenticatedAt,
	), nil
}
