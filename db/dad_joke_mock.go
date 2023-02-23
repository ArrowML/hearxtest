package db

import (
	"context"
	"hearxtest/model"

	"github.com/stretchr/testify/mock"
)

type MockDadJokeRepository struct {
	mock.Mock
}

func (m MockDadJokeRepository) Save(ctx context.Context, js *[]model.DadJoke) error {
	return nil
}

func (m MockDadJokeRepository) FetchPage(ctx context.Context, page int) (*[]model.DadJoke, error) {
	return nil, nil
}

func (pm MockDadJokeRepository) FetchRandom(ctx context.Context) (model.DadJoke, error) {
	return model.DadJoke{}, nil
}
