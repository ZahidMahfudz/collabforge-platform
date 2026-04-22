package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	dsn := "postgres://postgres:zahid1@localhost:5432/colabforge_db?sslmode=disable"
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	migration, err := os.ReadFile("migrations/000001_create_users_table.up.sql")
	if err != nil {
		log.Fatalf("Unable to read migration file: %v\n", err)
	}

	_, err = conn.Exec(ctx, string(migration))
	if err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}

	fmt.Println("Migration successful!")
}
