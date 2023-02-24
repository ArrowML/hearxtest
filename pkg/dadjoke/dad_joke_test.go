package dadjoke

import (
	"context"
	"hearxtest/db"
	"hearxtest/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	testCases := []struct {
		name   string
		jokes  *[]model.DadJoke
		expErr error
	}{
		{
			name:   "valid entry is saved",
			jokes:  &[]model.DadJoke{},
			expErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctx := context.TODO()
			r := db.MockDadJokeRepository{}

			err := Save(ctx, r, tc.jokes)

			assert.Equal(t, tc.expErr, err)

		})
	}
}

func TestGetPage(t *testing.T) {

}

func TestGetRandom(t *testing.T) {

}
