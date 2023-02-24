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
	db := connectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := createDadJokesTable(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func connectDB() *sql.DB {
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

func setupForTesing() (*sql.DB, error) {
	db := connectDB()
	err := createTestDadJokesTable(db)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func createTestDadJokesTable(db *sql.DB) error {
	ctx := context.TODO()

	query := `CREATE TABLE IF NOT EXISTS dad_jokes_test(
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

func dropTestTable(db *sql.DB) error {
	ctx := context.TODO()

	query := `DROP TABLE IF EXISTS dad_jokes_test CASCADE;`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
