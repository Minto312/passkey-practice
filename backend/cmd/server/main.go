package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"net/http"
	"os"
	"passkey-practice/backend/ent"
	repo "passkey-practice/backend/internal/infra/repository"
	usecase "passkey-practice/backend/internal/usecase/user"
	handler "passkey-practice/backend/internal/web/handler"
)

func main() {
	client, err := ent.Open("sqlite3", "file:passkey.db?cache=shared&_fk=1")
	if err != nil {
		slog.Error("failed opening connection to sqlite",
			slog.String("err", err.Error()),
		)
		os.Exit(1)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		slog.Error("failed creating schema resources",
			slog.String("err", err.Error()),
		)
		os.Exit(1)
	}

	userRepo := repo.NewEntUserRepository(client)
	userUC := usecase.NewRegisterUserUseCase(userRepo)

	r := gin.Default()
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/api/signup", handler.RegisterUserHandlerGin(userUC))

	slog.Info("Listening", slog.String("addr", ":8080"))
	if err := r.Run(":8080"); err != nil {
		slog.Error("server exited with error", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
