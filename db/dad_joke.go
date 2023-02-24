package db

import (
	"context"
	"database/sql"
	"hearxtest/model"
)

type PostgresDadJokeRepository struct {
	DB *sql.DB
}

func (pg PostgresDadJokeRepository) Save(ctx context.Context, js *[]model.DadJoke) (int, error) {
	return -1, nil
}

func (pg PostgresDadJokeRepository) FetchPage(ctx context.Context, page, limit int) (*[]model.DadJoke, error) {
	return nil, nil
}

func (pg PostgresDadJokeRepository) FetchJoke(ctx context.Context, id int) (model.DadJoke, error) {
	return model.DadJoke{}, nil
}

func (pg PostgresDadJokeRepository) FetchAllIDs(ctx context.Context) (*[]int, error) {
	return nil, nil
}
