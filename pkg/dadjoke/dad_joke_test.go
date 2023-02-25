package dadjoke

import (
	"context"
	"database/sql"
	"hearxtest/db"
	"hearxtest/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPage(t *testing.T) {

	ctx := context.TODO()

	testCases := []struct {
		name               string
		page               int
		limit              int
		fetchAllIdsMockRes []int
		fetchMockRes       *[]model.DadJoke
		expRes             *model.PaginatedDadJokes
		expErr             error
	}{
		{
			name:               "returns correct response of with default limit",
			page:               1,
			limit:              0,
			fetchAllIdsMockRes: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			fetchMockRes: &[]model.DadJoke{
				{Joke: "Joke1", Punchline: "Punchline1", Rating: 1},
				{Joke: "Joke2", Punchline: "Punchline2", Rating: 2},
				{Joke: "Joke3", Punchline: "Punchline3", Rating: 3},
				{Joke: "Joke4", Punchline: "Punchline4", Rating: 4},
				{Joke: "Joke5", Punchline: "Punchline5", Rating: 5},
				{Joke: "Joke6", Punchline: "Punchline6", Rating: 1},
				{Joke: "Joke7", Punchline: "Punchline7", Rating: 2},
				{Joke: "Joke8", Punchline: "Punchline8", Rating: 3},
				{Joke: "Joke9", Punchline: "Punchline9", Rating: 4},
				{Joke: "Joke10", Punchline: "Punchline10", Rating: 5},
			},
			expRes: &model.PaginatedDadJokes{
				Count: 12,
				Page:  1,
				Jokes: []model.DadJoke{
					{Joke: "Joke1", Punchline: "Punchline1", Rating: 1},
					{Joke: "Joke2", Punchline: "Punchline2", Rating: 2},
					{Joke: "Joke3", Punchline: "Punchline3", Rating: 3},
					{Joke: "Joke4", Punchline: "Punchline4", Rating: 4},
					{Joke: "Joke5", Punchline: "Punchline5", Rating: 5},
					{Joke: "Joke6", Punchline: "Punchline6", Rating: 1},
					{Joke: "Joke7", Punchline: "Punchline7", Rating: 2},
					{Joke: "Joke8", Punchline: "Punchline8", Rating: 3},
					{Joke: "Joke9", Punchline: "Punchline9", Rating: 4},
					{Joke: "Joke10", Punchline: "Punchline10", Rating: 5},
				},
			},
			expErr: nil,
		},
		{
			name:               "returns correct response with user defined limit",
			page:               1,
			limit:              5,
			fetchAllIdsMockRes: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			fetchMockRes: &[]model.DadJoke{
				{Joke: "Joke1", Punchline: "Punchline1", Rating: 1},
				{Joke: "Joke2", Punchline: "Punchline2", Rating: 2},
				{Joke: "Joke3", Punchline: "Punchline3", Rating: 3},
				{Joke: "Joke4", Punchline: "Punchline4", Rating: 4},
				{Joke: "Joke5", Punchline: "Punchline5", Rating: 5},
			},
			expRes: &model.PaginatedDadJokes{
				Count: 12,
				Page:  1,
				Jokes: []model.DadJoke{
					{Joke: "Joke1", Punchline: "Punchline1", Rating: 1},
					{Joke: "Joke2", Punchline: "Punchline2", Rating: 2},
					{Joke: "Joke3", Punchline: "Punchline3", Rating: 3},
					{Joke: "Joke4", Punchline: "Punchline4", Rating: 4},
					{Joke: "Joke5", Punchline: "Punchline5", Rating: 5},
				},
			},
			expErr: nil,
		},
		{
			name:               "returns error if page is larger than entries",
			page:               3,
			limit:              10,
			fetchAllIdsMockRes: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expRes:             nil,
			expErr:             ErrNoRecords,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := db.MockDadJokeRepository{}
			r.On("FetchAllIDs", ctx).Return(tc.fetchAllIdsMockRes, nil)
			r.On("FetchPage", ctx, tc.page, mock.Anything).Return(tc.fetchMockRes, nil)

			page, err := GetPage(ctx, r, tc.page, tc.limit)
			assert.Equal(t, tc.expRes, page)
			assert.Equal(t, tc.expErr, err)

		})
	}
}

func TestGetRandom(t *testing.T) {
	ctx := context.TODO()

	testCases := []struct {
		name               string
		page               int
		limit              int
		fetchAllIdsMockRes []int
		fetchMockRes       model.DadJoke
		expRes             *model.DadJoke
		expErr             error
	}{
		{
			name:               "returns correct response item",
			fetchAllIdsMockRes: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			fetchMockRes: model.DadJoke{
				Joke: "Joke", Punchline: "Punchline", Rating: 3,
			},
			expRes: &model.DadJoke{
				Joke: "Joke", Punchline: "Punchline", Rating: 3,
			},
			expErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := db.MockDadJokeRepository{}
			r.On("FetchAllIDs", ctx).Return(tc.fetchAllIdsMockRes, nil)
			r.On("FetchJoke", ctx, mock.Anything).Return(tc.fetchMockRes, nil)

			joke, err := GetRandom(ctx, r)
			assert.Equal(t, tc.expRes, joke)
			assert.Equal(t, tc.expErr, err)

		})
	}
}

func TestGetTotalRecords(t *testing.T) {
	ctx := context.TODO()

	testCases := []struct {
		name               string
		fetchAllIdsMockRes []int
		fetchAllIdsMockErr error
		expRes             int
		expErr             error
	}{
		{
			name:               "returns correct number of IDs",
			fetchAllIdsMockRes: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			fetchAllIdsMockErr: nil,
			expRes:             12,
			expErr:             nil,
		},
		{
			name:               "returns expected error if no records",
			fetchAllIdsMockRes: []int{},
			fetchAllIdsMockErr: sql.ErrNoRows,
			expRes:             -1,
			expErr:             ErrNoRecords,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := db.MockDadJokeRepository{}
			r.On("FetchAllIDs", ctx).Return(tc.fetchAllIdsMockRes, tc.fetchAllIdsMockErr)

			count, err := getTotalRecords(ctx, r)
			assert.Equal(t, tc.expRes, count)
			assert.Equal(t, tc.expErr, err)

		})
	}
}

func TestGetRandomID(t *testing.T) {
	ctx := context.TODO()

	testCases := []struct {
		name               string
		fetchAllIdsMockRes []int
		found              bool
		expErr             error
	}{
		{
			name:               "returns random number in ID list",
			fetchAllIdsMockRes: []int{143, 156, 158, 234, 235, 345},
			found:              true,
			expErr:             nil,
		},
		{
			name:               "return correct error if no records",
			fetchAllIdsMockRes: []int{},
			found:              false,
			expErr:             ErrNoRecords,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r := db.MockDadJokeRepository{}
			r.On("FetchAllIDs", ctx).Return(tc.fetchAllIdsMockRes, nil)

			rid, err := getRandomID(ctx, r)
			found := false
			for _, id := range tc.fetchAllIdsMockRes {
				if id == rid {
					found = true
					break
				}
			}
			assert.Equal(t, tc.found, found)
			assert.Equal(t, tc.expErr, err)

		})
	}
}
