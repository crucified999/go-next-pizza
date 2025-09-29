package main

import (
	"context"
	"fmt"

	"github.com/go-next-pizza/internal/storage"
)

func main() {
	ctx := context.Background()

	storageConfig := &storage.Config{
		Host: "localhost",
		Port: "5433",
		Password: "12345",
		User: "super_postgres",
		Database: "postgres",
		SSLMode: "disable",
	}

	storage := storage.NewStorage(storageConfig)

	if err := storage.Start(); err != nil {
		panic(err)
	}

	// err := storage.Pool().QueryRow("SELECT id FROM users WHERE id IS NOT NULL")
	// var id []string
	// var email []string

	db := storage.Pool()

	query := `SELECT id, email FROM users`
	rows, err := db.Query(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id string
			email string
		)

		err := rows.Scan(&id, &email)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, email)
	}
	
}
