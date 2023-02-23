package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitPostgresDB() *sql.DB {

	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASSWORD")
	dbHost := os.Getenv("PG_HOST")
	dbName := os.Getenv("PG_DB")
	dbSSL := os.Getenv("PG_SSL")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbName, dbSSL)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = createDadJokesTable(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createDadJokesTable(ctx context.Context, db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS dad_jokes(
		id SERIAL PRIMARY KEY ,
		joke TEXT NOT NULL,
		punchline TEXT NOT NULL,
        rating INT
		);`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
