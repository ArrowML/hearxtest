package dadjoke

import (
	"context"
	"hearxtest/model"
)

func Save(ctx context.Context, r Repository, jokes *[]model.DadJoke) error {
	return nil
}

func GetPage(ctx context.Context, r Repository, page int) (*[]model.DadJoke, error) {
	return nil, nil
}

func GetRandom(ctx context.Context, r Repository) (*model.DadJoke, error) {
	return nil, nil
}
