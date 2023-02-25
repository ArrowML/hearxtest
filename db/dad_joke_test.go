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

var dbc *sql.DB

func setup() {
	cwd, _ := os.Getwd()
	err := godotenv.Load(cwd + "/../" + "local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbc, err = setupForTesing()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func teardown() {
	dropTestTable(dbc)
	dbc.Close()
}

func TestSave(t *testing.T) {
	setup()

	pgRepo := PostgresDadJokeRepository{
		DB:        dbc,
		TableName: "dad_jokes_test",
	}
	ctx := context.TODO()

	testCases := []struct {
		name     string
		jokes    *[]model.DadJoke
		expCount int64
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

	teardown()
}

func TestFetchPage(t *testing.T) {

	setup()

	pgRepo := PostgresDadJokeRepository{
		DB:        dbc,
		TableName: "dad_jokes_test",
	}
	ctx := context.TODO()

	jokes := &[]model.DadJoke{
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
		{
			Joke:      "Test Joke 4",
			Punchline: "Test Punchline 4",
			Rating:    1,
		},
		{
			Joke:      "Test Joke 5",
			Punchline: "Test Punchline 5",
			Rating:    2,
		},
		{
			Joke:      "Test Joke 6",
			Punchline: "Test Punchline 6",
			Rating:    3,
		},
		{
			Joke:      "Test Joke 7",
			Punchline: "Test Punchline 7",
			Rating:    1,
		},
		{
			Joke:      "Test Joke 8",
			Punchline: "Test Punchline 8",
			Rating:    2,
		},
		{
			Joke:      "Test Joke 9",
			Punchline: "Test Punchline 9",
			Rating:    3,
		},
	}
	_, err := pgRepo.Save(ctx, jokes)
	require.Nil(t, err)

	testCases := []struct {
		name     string
		page     int
		limit    int
		expCount int
		expErr   error
	}{
		{
			name:     "all entries returned for page",
			page:     1,
			limit:    5,
			expCount: 5,
			expErr:   nil,
		},
		{
			name:     "page returns correct number limit",
			page:     2,
			limit:    5,
			expCount: 4,
			expErr:   nil,
		},
		{
			name:     "page exceeding entries returns correct error",
			page:     3,
			limit:    5,
			expCount: 0,
			expErr:   sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rows, err := pgRepo.FetchPage(ctx, tc.page, tc.limit)
			fmt.Print(rows, err)
			if rows != nil {
				assert.Equal(t, tc.expCount, len(*rows))
			}
			assert.Equal(t, tc.expErr, err)
		})
	}

	teardown()
}

func TestFetchJoke(t *testing.T) {

	setup()

	pgRepo := PostgresDadJokeRepository{
		DB:        dbc,
		TableName: "dad_jokes_test",
	}
	ctx := context.TODO()

	testCases := []struct {
		name       string
		savedJokes *[]model.DadJoke
		expJoke    model.DadJoke
		expErr     error
	}{
		{
			name:       "no entry id returns correct error",
			savedJokes: nil,
			expJoke:    model.DadJoke{},
			expErr:     sql.ErrNoRows,
		},
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.savedJokes != nil {
				_, err := pgRepo.Save(ctx, tc.savedJokes)
				require.Nil(t, err)
			}

			j, err := pgRepo.FetchJoke(ctx, 1)
			assert.Equal(t, tc.expJoke, j)
			assert.Equal(t, tc.expErr, err)
		})
	}

	teardown()
}

func TestFetchAllIDs(t *testing.T) {

	setup()

	pgRepo := PostgresDadJokeRepository{
		DB:        dbc,
		TableName: "dad_jokes_test",
	}
	ctx := context.TODO()

	testCases := []struct {
		name       string
		savedJokes *[]model.DadJoke
		expRes     []int
		expErr     error
	}{
		{
			name:       "no entries returns correct error",
			savedJokes: nil,
			expRes:     nil,
			expErr:     sql.ErrNoRows,
		},
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
			expRes: []int{1, 2, 3},
			expErr: nil,
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

	teardown()

}
