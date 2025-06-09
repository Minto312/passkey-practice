package handler

import (
	"encoding/json"
	usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RegisterUserHandler(uc *usecase.RegisterUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var req RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		user, err := uc.Register(usecase.RegisterUserInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := RegisterUserResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email.String(),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func RegisterUserHandlerGin(uc *usecase.RegisterUserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		user, err := uc.Register(usecase.RegisterUserInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		res := RegisterUserResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email.String(),
		}
		c.JSON(200, res)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装
}
