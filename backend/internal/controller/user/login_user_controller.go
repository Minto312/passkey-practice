package user

import (
	"net/http"

	"github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	loginUserUseCase user.LoginUserUseCase
}

func NewLoginUserController(loginUserUseCase user.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{
		loginUserUseCase: loginUserUseCase,
	}
}

func (c *LoginUserController) LoginUser(ctx *gin.Context) {
	var reqBody struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := user.LoginUserInput{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}

	output, err := c.loginUserUseCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		// TODO: エラーハンドリングを詳細化する
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resBody := gin.H{
		"access_token":  output.AccessToken,
		"refresh_token": output.RefreshToken,
	}

	ctx.JSON(http.StatusOK, resBody)
}
