package infra

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			slog.Error("Failed to load .env file")
			panic(err)
		}
	}
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping the database: %w", err)
	}
	return db, nil
}

func Migrate(db *sql.DB) error {
	slog.Info("Migrating the database...")
	createTableDDL := []string{
		`
    CREATE TABLE IF NOT EXISTS hosts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        port INTEGER NOT NULL
    );
    `,
		`
	CREATE TABLE IF NOT EXISTS spans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		host_id INTEGER NOT NULL,
		duration_time INTEGER NOT NULL,
	);
	`,
	}

	for _, ddl := range createTableDDL {
		_, err := db.Exec(ddl)
		if err != nil {
			return fmt.Errorf("Failed to create hosts table: %w", err)
		}
	}

	return nil
}
