package auth_history

import (
	"context"

	"github.com/Minto312/passkey-practice/backend/internal/domain/auth_history"
)

type GetAuthHistoriesUseCase interface {
	Execute(ctx context.Context) (*GetAuthHistoriesOutput, error)
}

type GetAuthHistoriesOutput struct {
	AuthHistories []*auth_history.AuthHistory
}

type getAuthHistoriesInteractor struct {
	authHistoryRepo auth_history.AuthHistoryRepository
}

func NewGetAuthHistoriesInteractor(
	authHistoryRepo auth_history.AuthHistoryRepository,
) GetAuthHistoriesUseCase {
	return &getAuthHistoriesInteractor{
		authHistoryRepo: authHistoryRepo,
	}
}

func (i *getAuthHistoriesInteractor) Execute(ctx context.Context) (*GetAuthHistoriesOutput, error) {
	histories, err := i.authHistoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return &GetAuthHistoriesOutput{
		AuthHistories: histories,
	}, nil
}
