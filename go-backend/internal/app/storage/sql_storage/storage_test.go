package sql_storage_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=go_next_pizza sslmode=disable"
		os.Setenv("DATABASE_URL", databaseURL)
	}

	os.Exit(m.Run())
}