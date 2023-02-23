package dadjoke

import (
	"context"
	model "hearxtest/model"
)

type Repository interface {
	Save(context.Context, *[]model.DadJoke) error
	FetchPage(context.Context, int) (*[]model.DadJoke, error)
	FetchRandom(context.Context) (model.DadJoke, error)
}
