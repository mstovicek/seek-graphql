package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/model"
)

type categoriesEdgeResolver struct {
	ctx        context.Context
	cardReader cardReaderInterface
	cursor     graphql.ID
	model      *model.Category
}

func (r *categoriesEdgeResolver) Cursor() (graphql.ID, error) {
	return r.cursor, nil
}

func (r *categoriesEdgeResolver) Node() (*categoryResolver, error) {
	return newCategoryResolverWithModel(r.ctx, r.cardReader, r.model)
}
