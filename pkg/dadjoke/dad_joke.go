package dadjoke

import (
	"context"
	"database/sql"
	"errors"
	"hearxtest/model"
	"math/rand"
	"time"
)

func Save(ctx context.Context, r Repository, jokes *[]model.DadJoke) error {
	_, err := r.Save(ctx, jokes)
	if err != nil {
		return ErrSavingJokes
	}
	return nil
}

func GetPage(ctx context.Context, r Repository, page, records int) (*model.PaginatedDadJokes, error) {

	limit := 10
	if records != 0 {
		limit = records
	}
	total, err := getTotalRecords(ctx, r)
	if err != nil {
		return nil, ErrFetchingJokes
	}
	if total == 0 {
		return nil, ErrNoRecords
	}
	if (limit*page)-total > limit {
		return nil, ErrNoRecords
	}

	jokes, err := r.FetchPage(ctx, page, limit)
	if err != nil {
		return nil, ErrFetchingJokes
	}

	res := model.PaginatedDadJokes{
		Count: total,
		Page:  page,
		Jokes: *jokes,
	}

	return &res, nil
}

func GetRandom(ctx context.Context, r Repository) (*model.DadJoke, error) {

	id, err := getRandomID(ctx, r)
	if err != nil {
		return nil, ErrFetchingJokes
	}

	joke, err := r.FetchJoke(ctx, id)
	if err != nil {
		return nil, ErrFetchingJokes
	}

	return &joke, nil
}

func getTotalRecords(ctx context.Context, r Repository) (int, error) {
	ids, err := r.FetchAllIDs(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, ErrNoRecords
		}
		return -1, err
	}
	return len(ids), nil
}

func getRandomID(ctx context.Context, r Repository) (int, error) {
	ids, err := r.FetchAllIDs(ctx)
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, ErrNoRecords
	}
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(ids) - 1
	i := rand.Intn(max-min) + min
	return ids[i], nil
}
