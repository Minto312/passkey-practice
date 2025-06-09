package main

import (
	"context"
	"fmt"
	"github.com/Minto312/passkey-practice/backend/ent"
	repo "github.com/Minto312/passkey-practice/backend/internal/infra/repository"
	usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	handler "github.com/Minto312/passkey-practice/backend/internal/web/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server exited with error", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	client, err := setupDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			slog.Error("failed to close database client", slog.String("err", err.Error()))
		}
	}()

	r := setupRouter(client)

	slog.Info("Listening", slog.String("addr", ":8080"))
	return r.Run(":8080")
}

func setupDB(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:passkey.db?cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}
	return client, nil
}

func setupRouter(client *ent.Client) *gin.Engine {
	userRepo := repo.NewEntUserRepository(client)
	userUC := usecase.NewRegisterUserUseCase(userRepo)

	r := gin.Default()
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/api/signup", handler.RegisterUserHandlerGin(userUC))
	return r
}
