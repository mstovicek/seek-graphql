package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-2/model"
	"context"
)

type cardsEdgeResolver struct {
	ctx context.Context
	cursor graphql.ID
	model  *model.Card
}

func (r *cardsEdgeResolver) Cursor() (graphql.ID, error) {
	return r.cursor, nil
}

func (r *cardsEdgeResolver) Node() (*cardResolver, error) {
	return newCardResolvedWithModel(r.ctx, r.model)
}
