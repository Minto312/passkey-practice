package passkey

import (
	"net/http"
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/controller/middleware"
	"github.com/Minto312/passkey-practice/backend/internal/domain/user"
	"github.com/Minto312/passkey-practice/backend/internal/usecase/passkey"
	"github.com/gin-gonic/gin"
)

type GetPasskeysController struct {
	useCase *passkey.GetPasskeysUseCase
}

func NewGetPasskeysController(useCase *passkey.GetPasskeysUseCase) *GetPasskeysController {
	return &GetPasskeysController{useCase: useCase}
}

type PasskeyResponse struct {
	ID           string    `json:"id"`
	CredentialID string    `json:"credential_id"`
	DeviceName   string    `json:"device_name"`
	CreatedAt    time.Time `json:"created_at"`
	LastUsedAt   time.Time `json:"last_used_at"`
}

func (c *GetPasskeysController) GetPasskeys(ctx *gin.Context) {
	userIDStr, exists := ctx.Get(middleware.UserIDKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	userID, err := user.ParseUserID(userIDStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid userID format"})
		return
	}

	passkeys, err := c.useCase.Execute(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]PasskeyResponse, len(passkeys))
	for i, p := range passkeys {
		response[i] = PasskeyResponse{
			ID:           p.ID().String(),
			CredentialID: p.CredentialID().Value(),
			DeviceName:   p.DeviceName().Value(),
			CreatedAt:    p.CreatedAt(),
			LastUsedAt:   p.LastUsedAt(),
		}
	}

	ctx.JSON(http.StatusOK, response)
}
