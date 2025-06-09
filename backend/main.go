package main

import (
	"context"
	"log"

	"github.com/Minto312/passkey-practice/backend/ent/migrate"
	"github.com/Minto312/passkey-practice/backend/internal/controller/user"
	"github.com/Minto312/passkey-practice/backend/internal/infra/repository"
	user_usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Minto312/passkey-practice/backend/ent"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=devcontainer-db-1 port=5432 user=postgres password=postgres dbname=passkey sslmode=disable"
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("ent: migration completed")

	// DI
	userRepo := repository.NewUserRepository(client)
	sessionRepo := repository.NewSessionRepository(client)

	registerUserUseCase := user_usecase.NewRegisterUserInteractor(userRepo)
	registerUserController := user.NewRegisterUserController(registerUserUseCase)

	loginUserUseCase := user_usecase.NewLoginUserInteractor(userRepo, sessionRepo)
	loginUserController := user.NewLoginUserController(loginUserUseCase)

	// Setup router
	router := gin.Default()

	// CORSミドルウェアの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.POST("/signup", registerUserController.RegisterUser)
	router.POST("/login", loginUserController.LoginUser)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
