package loaderdummy

import (
	"context"
	"github.com/mstovicek/seek-graphql/model"
)

func NewMeLoader(ctx context.Context) (*meLoader, error) {
	return &meLoader{
		ctx: ctx,
	}, nil
}

type meLoader struct {
	ctx context.Context
}

func (loader *meLoader) LoadMeByCtx(ctx context.Context) (*model.Me, error) {
	return &model.Me{
		ID:    "42",
		Name:  "Milan",
		Email: "milan@me",
	}, nil
}
