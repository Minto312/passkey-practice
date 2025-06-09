package main

import (
	"context"
	"log"

	"github.com/Minto312/passkey-practice/backend/internal/infra/repository/ent"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=db port=5432 user=postgres password=postgres dbname=passkey sslmode=disable"
	client, err := ent.Open("postgres", dsn, ent.Debug())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("ent: migration completed")
}
