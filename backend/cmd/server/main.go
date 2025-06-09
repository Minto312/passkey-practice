package main

import (
    "context"
    "log"
    "net/http"
    "passkey-practice/backend/ent"
    repo "passkey-practice/backend/internal/infra/repository"
    usecase "passkey-practice/backend/internal/usecase/user"
    handler "passkey-practice/backend/internal/web/handler"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    client, err := ent.Open("sqlite3", "file:passkey.db?cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    userRepo := repo.NewEntUserRepository(client)
    userUC := usecase.NewRegisterUserUseCase(userRepo)

    http.HandleFunc("/api/signup", handler.RegisterUserHandler(userUC))

    log.Println("Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
} 