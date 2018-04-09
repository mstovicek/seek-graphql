package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-2/model"
	"context"
)

type categoriesEdgeResolver struct {
	ctx context.Context
	cursor graphql.ID
	model  *model.Category
}

func (r *categoriesEdgeResolver) Cursor() (graphql.ID, error) {
	return r.cursor, nil
}

func (r *categoriesEdgeResolver) Node() (*categoryResolver, error) {
	return newCategoryResolverWithModel(r.ctx, r.model)
}
