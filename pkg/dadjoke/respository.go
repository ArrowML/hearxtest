package dadjoke

import (
	"context"
	"hearxtest/model"
)

type Repository interface {
	Save(context.Context, *[]model.DadJoke) (int64, error)
	FetchPage(context.Context, int, int) (*[]model.DadJoke, error)
	FetchJoke(context.Context, int) (model.DadJoke, error)
	FetchAllIDs(context.Context) ([]int, error)
}
