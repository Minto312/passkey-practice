package user

import (
	"net/http"

	"github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

// RegisterUserController はユーザー登録コントローラーです。
type RegisterUserController struct {
	useCase user.RegisterUserUseCase
}

// NewRegisterUserController は RegisterUserController の新しいインスタンスを生成します。
func NewRegisterUserController(useCase user.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{
		useCase: useCase,
	}
}

// RegisterUser はユーザーを登録するための Gin ハンドラ関数です。
func (c *RegisterUserController) RegisterUser(ctx *gin.Context) {
	var reqBody struct {
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=8"`
		DisplayName string `json:"display_name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := user.RegisterUserInput{
		Email:       reqBody.Email,
		Password:    reqBody.Password,
		DisplayName: reqBody.DisplayName,
	}

	output, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		// TODO: エラーハンドリングを詳細化する
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resBody := gin.H{
		"id":           output.User.ID().String(),
		"email":        string(output.User.Email()),
		"display_name": string(output.User.DisplayName()),
	}

	ctx.JSON(http.StatusCreated, resBody)
}
