package db

import (
	"context"
	"hearxtest/model"

	"github.com/stretchr/testify/mock"
)

type MockDadJokeRepository struct {
	mock.Mock
}

func (m MockDadJokeRepository) Save(ctx context.Context, js *[]model.DadJoke) (int, error) {
	args := m.Called(ctx, js)
	return args.Get(0).(int), args.Error(1)
}

func (m MockDadJokeRepository) FetchPage(ctx context.Context, page, limit int) (*[]model.DadJoke, error) {
	args := m.Called(ctx, page)
	return args.Get(0).(*[]model.DadJoke), args.Error(1)
}

func (m MockDadJokeRepository) FetchJoke(ctx context.Context, id int) (model.DadJoke, error) {
	args := m.Called(ctx)
	return args.Get(0).(model.DadJoke), args.Error(1)
}

func (m MockDadJokeRepository) FetchAllIDs(ctx context.Context) (*[]int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]int), args.Error(1)
}
