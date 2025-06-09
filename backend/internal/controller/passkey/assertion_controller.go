package passkey

import (
	"net/http"

	passkey_domain "github.com/Minto312/passkey-practice/backend/internal/domain/passkey"
	passkey_usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/passkey"
	"github.com/gin-gonic/gin"
)

type AssertionController struct {
	useCase *passkey_usecase.AssertionUseCase
}

func NewAssertionController(useCase *passkey_usecase.AssertionUseCase) *AssertionController {
	return &AssertionController{useCase: useCase}
}

func (c *AssertionController) GetAssertion(ctx *gin.Context) {
	webAuthn, err := passkey_domain.BeginAssertion(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	

	
}