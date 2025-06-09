package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Minto312/passkey-practice/backend/internal/controller/user"
	domain "github.com/Minto312/passkey-practice/backend/internal/domain/user"
	user_usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MockRegisterUserUseCase は RegisterUserUseCase のモックです。
type MockRegisterUserUseCase struct {
	ExecuteFunc func(ctx context.Context, input user_usecase.RegisterUserInput) (*user_usecase.RegisterUserOutput, error)
}

func (m *MockRegisterUserUseCase) Execute(ctx context.Context, input user_usecase.RegisterUserInput) (*user_usecase.RegisterUserOutput, error) {
	return m.ExecuteFunc(ctx, input)
}

func TestRegisterUserController_RegisterUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("正常系: ユーザー登録が成功する", func(t *testing.T) {
		// 準備
		mockUseCase := &MockRegisterUserUseCase{
			ExecuteFunc: func(ctx context.Context, input user_usecase.RegisterUserInput) (*user_usecase.RegisterUserOutput, error) {
				userID, _ := domain.ParseUserID(uuid.NewString())
				email, _ := domain.NewEmail(input.Email)
				displayName, _ := domain.NewDisplayName(input.DisplayName)
				passwordHash, _ := domain.NewPasswordHash("hashed_password")

				return &user_usecase.RegisterUserOutput{
					User: domain.Reconstruct(userID, email, passwordHash, displayName, time.Now(), time.Now()),
				}, nil
			},
		}
		controller := user.NewRegisterUserController(mockUseCase)

		router := gin.New()
		router.POST("/users", controller.RegisterUser)

		reqBody := map[string]string{
			"email":        "test@example.com",
			"password":     "password123",
			"display_name": "Test User",
		}
		jsonBody, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		// 実行
		router.ServeHTTP(rr, req)

		// 検証
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var resBody map[string]string
		if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if _, err := uuid.Parse(resBody["id"]); err != nil {
			t.Errorf("response body 'id' is not a valid UUID: %v", resBody["id"])
		}
		if resBody["email"] != "test@example.com" {
			t.Errorf("handler returned unexpected body: got %v want %v", resBody["email"], "test@example.com")
		}
		if resBody["display_name"] != "Test User" {
			t.Errorf("handler returned unexpected body: got %v want %v", resBody["display_name"], "Test User")
		}
	})
}
