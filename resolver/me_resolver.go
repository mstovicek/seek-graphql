package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/model"
)

func newMeResolverByCtx(
	ctx context.Context,
	reader meReaderInterface,
) (*meResolver, error) {
	me, err := reader.LoadMeByCtx(ctx)
	if err != nil {
		return nil, err
	}

	return &meResolver{
		ctx: ctx,
		me:  me,
	}, nil
}

type meResolver struct {
	ctx context.Context
	me  *model.Me
}

func (r *meResolver) ID() (graphql.ID, error) {
	return graphql.ID(r.me.ID), nil
}

func (r *meResolver) Name() (*string, error) {
	return &r.me.Name, nil
}

func (r *meResolver) Email() (string, error) {
	return r.me.Email, nil
}
