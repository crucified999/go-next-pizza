package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-next-pizza/internal/app/storage/sql_storage"
)

func Start(config *Config) error {
	db, err := newDB(config.DataBaseUrl)

	if err != nil {
		return err
	}

	defer db.Close()
	

    storage := sql_storage.NewSQLStorage(db)
    s := newServer(
        storage,
        []byte(config.JWTSecret),
        config.JWTTTLSeconds,
        config.RefreshTTLSeconds,
    )

	fmt.Println("Server is starting...")
	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}