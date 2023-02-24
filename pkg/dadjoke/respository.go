package dadjoke

import (
	"context"
	"hearxtest/model"
)

type Repository interface {
	Save(context.Context, *[]model.DadJoke) (int, error)
	FetchPage(context.Context, int, int) (*[]model.DadJoke, error)
	FetchJoke(context.Context, int) (model.DadJoke, error)
	FetchAllIds(context.Context) (*[]int, error)
}
