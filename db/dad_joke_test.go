package db

import (
	"context"
	"database/sql"
	"fmt"
	"hearxtest/model"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db *sql.DB

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {

	cwd, _ := os.Getwd()
	err = godotenv.Load(cwd + "/../" + "local.env")
	fmt.Print(err)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := setupForTesing()
	if err != nil {
		return -1, fmt.Errorf("could not connect to database: %w", err)
	}

	defer func() {
		dropTestTable(db)
		db.Close()
	}()
	return m.Run(), nil
}

func TestSave(t *testing.T) {

	pgRepo := PostgresDadJokeRepository{
		DB: db,
	}
	ctx := context.TODO()

	testCases := []struct {
		name     string
		jokes    *[]model.DadJoke
		expCount int
		expErr   error
	}{
		{
			name: "single entry is saved correctly",
			jokes: &[]model.DadJoke{
				{
					Joke:      "Test Joke",
					Punchline: "Test Punchline",
					Rating:    3,
				},
			},
			expCount: 1,
			expErr:   nil,
		},
		{
			name: "multiple entries are saved correctly",
			jokes: &[]model.DadJoke{
				{
					Joke:      "Test Joke 1",
					Punchline: "Test Punchline 1",
					Rating:    1,
				},
				{
					Joke:      "Test Joke 2",
					Punchline: "Test Punchline 2",
					Rating:    2,
				},
				{
					Joke:      "Test Joke 3",
					Punchline: "Test Punchline 3",
					Rating:    3,
				},
			},
			expCount: 3,
			expErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows, err := pgRepo.Save(ctx, tc.jokes)
			assert.Equal(t, tc.expCount, rows)
			require.Nil(t, err)
		})
	}

}

func TestFetchPage(t *testing.T) {

	pgRepo := PostgresDadJokeRepository{
		DB: db,
	}
	ctx := context.TODO()

	testCases := []struct {
		name     string
		jokes    *[]model.DadJoke
		expCount int
		expErr   error
	}{
		{
			name: "single entry is saved correctly",
			jokes: &[]model.DadJoke{
				{
					Joke:      "Test Joke",
					Punchline: "Test Punchline",
					Rating:    3,
				},
			},
			expCount: 1,
			expErr:   nil,
		},
		{
			name: "multiple entries are saved correctly",
			jokes: &[]model.DadJoke{
				{
					Joke:      "Test Joke 1",
					Punchline: "Test Punchline 1",
					Rating:    1,
				},
				{
					Joke:      "Test Joke 2",
					Punchline: "Test Punchline 2",
					Rating:    2,
				},
				{
					Joke:      "Test Joke 3",
					Punchline: "Test Punchline 3",
					Rating:    3,
				},
			},
			expCount: 3,
			expErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows, err := pgRepo.Save(ctx, tc.jokes)
			assert.Equal(t, tc.expCount, rows)
			require.Nil(t, err)
		})
	}

}

func TestFetchJoke(t *testing.T) {

	pgRepo := PostgresDadJokeRepository{
		DB: db,
	}
	ctx := context.TODO()

	testCases := []struct {
		name       string
		savedJokes *[]model.DadJoke
		expJoke    model.DadJoke
		expErr     error
	}{
		{
			name: "single entry is returned correctly",
			savedJokes: &[]model.DadJoke{
				{
					Joke:      "Test Joke 1",
					Punchline: "Test Punchline 1",
					Rating:    1,
				},
			},
			expJoke: model.DadJoke{
				Joke:      "Test Joke 1",
				Punchline: "Test Punchline 1",
				Rating:    1,
			},
			expErr: nil,
		},
		{
			name:       "no entry id returns correct error",
			savedJokes: nil,
			expJoke: model.DadJoke{
				Joke:      "Test Joke 1",
				Punchline: "Test Punchline 1",
				Rating:    1,
			},
			expErr: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.savedJokes != nil {
				_, err := pgRepo.Save(ctx, tc.savedJokes)
				require.Nil(t, err)
			}

			j, err := pgRepo.FetchJoke(ctx, 0)
			assert.Equal(t, tc.expJoke, j)
			assert.Equal(t, tc.expErr, err)
		})
	}

}

func TestFetchAllIDs(t *testing.T) {

	pgRepo := PostgresDadJokeRepository{
		DB: db,
	}
	ctx := context.TODO()

	testCases := []struct {
		name       string
		savedJokes *[]model.DadJoke
		expRes     *[]int
		expErr     error
	}{
		{
			name: "returns all ids correctly",
			savedJokes: &[]model.DadJoke{
				{
					Joke:      "Test Joke 1",
					Punchline: "Test Punchline 1",
					Rating:    1,
				},
				{
					Joke:      "Test Joke 1",
					Punchline: "Test Punchline 1",
					Rating:    1,
				},
				{
					Joke:      "Test Joke 1",
					Punchline: "Test Punchline 1",
					Rating:    1,
				},
			},
			expRes: &[]int{0, 1, 2},
			expErr: nil,
		},
		{
			name:       "no entries returns correct error",
			savedJokes: nil,
			expRes:     nil,
			expErr:     sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.savedJokes != nil {
				_, err := pgRepo.Save(ctx, tc.savedJokes)
				require.Nil(t, err)
			}

			ids, err := pgRepo.FetchAllIDs(ctx)
			assert.Equal(t, tc.expRes, ids)
			assert.Equal(t, tc.expErr, err)
		})
	}

}
