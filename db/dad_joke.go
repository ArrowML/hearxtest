package db

import (
	"context"
	"database/sql"
	"hearxtest/model"
)

type PostgresDadJokeRepository struct {
	DB *sql.DB
}

func (pg PostgresDadJokeRepository) Save(ctx context.Context, js *[]model.DadJoke) error {
	return nil
}

func (pg PostgresDadJokeRepository) FetchPage(ctx context.Context, page int) (*[]model.DadJoke, error) {
	return nil, nil
}

func (pg PostgresDadJokeRepository) FetchRandom(ctx context.Context) (model.DadJoke, error) {
	return model.DadJoke{}, nil
}
