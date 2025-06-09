package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Minto312/passkey-practice/backend/ent"
	repo "github.com/Minto312/passkey-practice/backend/internal/infra/repository"
	usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	handler "github.com/Minto312/passkey-practice/backend/internal/web/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=passkey password=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		api.POST("/signup", handler.RegisterUserHandlerGin(userUC))
	}

	return r
}
