package repository

import (
	"context"
	"fmt"

	"github.com/Minto312/passkey-practice/backend/ent"
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

func (r *authHistoryRepository) Save(ctx context.Context, authHistory *auth_history.AuthHistory) error {
	_, err := r.client.AuthHistory.
		Create().
		SetID(authHistory.ID().UUID).
		SetUserID(authHistory.UserID().UUID).
		SetMethod(string(authHistory.Method())).
		SetIPAddress(string(authHistory.IPAddress())).
		SetUserAgent(string(authHistory.UserAgent())).
		SetAuthenticatedAt(authHistory.AuthenticatedAt()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating auth history: %w", err)
	}
	return nil
}

func (r *authHistoryRepository) FindAll(ctx context.Context) ([]*auth_history.AuthHistory, error) {
	histories, err := r.client.AuthHistory.Query().WithUser().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth histories: %w", err)
	}

	domainHistories := make([]*auth_history.AuthHistory, len(histories))
	for i, h := range histories {
		domainHistory, err := toDomainAuthHistory(h)
		if err != nil {
			return nil, fmt.Errorf("failed to convert auth history: %w", err)
		}
		domainHistories[i] = domainHistory
	}

	return domainHistories, nil
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
