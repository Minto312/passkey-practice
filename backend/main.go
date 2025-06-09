package main

import (
	"context"
	"log"

	"github.com/Minto312/passkey-practice/backend/ent/migrate"
	auth_history_controller "github.com/Minto312/passkey-practice/backend/internal/controller/auth_history"
	"github.com/Minto312/passkey-practice/backend/internal/controller/middleware"
	passkey_controller "github.com/Minto312/passkey-practice/backend/internal/controller/passkey"
	"github.com/Minto312/passkey-practice/backend/internal/controller/user"
	"github.com/Minto312/passkey-practice/backend/internal/infra/repository"
	infra_session "github.com/Minto312/passkey-practice/backend/internal/infra/session"
	auth_history_usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/auth_history"
	passkey_usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/passkey"
	user_usecase "github.com/Minto312/passkey-practice/backend/internal/usecase/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"

	"github.com/Minto312/passkey-practice/backend/ent"
	"github.com/Minto312/passkey-practice/backend/internal/auth"
	_ "github.com/lib/pq"
)

func main() {
	if err := auth.InitWebAuthn(); err != nil {
		log.Fatalf("failed to initialize webauthn: %v", err)
	}
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
	authHistoryRepo := repository.NewAuthHistoryRepository(client)
	passkeyRepo := repository.NewPasskeyRepository(client)
	webauthnSessionStore := infra_session.NewStore()

	registerUserUseCase := user_usecase.NewRegisterUserInteractor(userRepo)
	registerUserController := user.NewRegisterUserController(registerUserUseCase)

	loginUserUseCase := user_usecase.NewLoginUserInteractor(userRepo, sessionRepo, authHistoryRepo)
	loginUserController := user.NewLoginUserController(loginUserUseCase)

	getAuthHistoriesUseCase := auth_history_usecase.NewGetAuthHistoriesInteractor(authHistoryRepo)
	getAuthHistoriesController := auth_history_controller.NewGetAuthHistoriesController(getAuthHistoriesUseCase)

	getPasskeysUseCase := passkey_usecase.NewGetPasskeysUseCase(passkeyRepo)
	getPasskeysController := passkey_controller.NewGetPasskeysController(getPasskeysUseCase)

	registerPasskeyChallengeUseCase := passkey_usecase.NewRegisterPasskeyChallengeUseCase(userRepo, passkeyRepo, webauthnSessionStore)
	registerPasskeyChallengeController := passkey_controller.NewRegisterPasskeyChallengeController(registerPasskeyChallengeUseCase)

	registerPasskeyUseCase := passkey_usecase.NewRegisterPasskeyUseCase(userRepo, passkeyRepo, webauthnSessionStore)
	registerPasskeyController := passkey_controller.NewRegisterPasskeyController(registerPasskeyUseCase)

	

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

	authRouter := router.Group("/auth")
	authRouter.Use(middleware.AuthMiddleware())
	{
		authRouter.GET("/history", getAuthHistoriesController.GetAuthHistories)
		authRouter.GET("/assertion", assertionController.GetAssertion)
		authRouter.POST("/assertion", assertionController.ChallengeAssertion)
		authRouter.GET("/attestation", attestationController.RegisterPasskeyChallenge)
		authRouter.POST("/attestation", attestationController.RegisterPasskey)	
		authRouter.GET("/passkey", getPasskeysController.GetPasskeys)
	}

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
