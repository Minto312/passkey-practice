package auth_history

import (
	"net/http"
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/usecase/auth_history"
	"github.com/gin-gonic/gin"
)

type GetAuthHistoriesController struct {
	useCase auth_history.GetAuthHistoriesUseCase
}

func NewGetAuthHistoriesController(useCase auth_history.GetAuthHistoriesUseCase) *GetAuthHistoriesController {
	return &GetAuthHistoriesController{
		useCase: useCase,
	}
}

type AuthHistoryResponse struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Method          string    `json:"method"`
	IPAddress       string    `json:"ip_address"`
	UserAgent       string    `json:"user_agent"`
	AuthenticatedAt time.Time `json:"authenticated_at"`
}

func (c *GetAuthHistoriesController) GetAuthHistories(ctx *gin.Context) {
	output, err := c.useCase.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]AuthHistoryResponse, len(output.AuthHistories))
	for i, h := range output.AuthHistories {
		response[i] = AuthHistoryResponse{
			ID:              h.ID().UUID.String(),
			UserID:          h.UserID().UUID.String(),
			Method:          string(h.Method()),
			IPAddress:       string(h.IPAddress()),
			UserAgent:       string(h.UserAgent()),
			AuthenticatedAt: h.AuthenticatedAt(),
		}
	}

	ctx.JSON(http.StatusOK, response)
}
