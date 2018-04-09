package resolver

import (
	"github.com/mstovicek/seek-2/model"
	"github.com/mstovicek/seek-2/service"
	"context"
	"github.com/mstovicek/seek-2/loader"
)

type categoriesConnectionResolver struct {
	ctx        context.Context
	categories []*model.Category
	totalCount *int
	from       *string
	to         *string
}

func newCategoriesConnectionResolver(
	ctx context.Context,
	first int,
	afterCursor *string,
) (*categoriesConnectionResolver, error) {
	afterID, _ := service.DecodeCursor(afterCursor)

	categories, err := loader.ListCategories(ctx, first, afterID)
	if err != nil {
		return nil, err
	}

	totalCount, _ := loader.CategoriesTotalCount()

	return &categoriesConnectionResolver{
		ctx:        ctx,
		categories: categories,
		totalCount: totalCount,
		from:       &(categories[0].ID),
		to:         &(categories[len(categories)-1].ID),
	}, nil
}

func (r *categoriesConnectionResolver) TotalCount() (int32, error) {
	return int32(*r.totalCount), nil
}

func (r *categoriesConnectionResolver) Edges() (*[]*categoriesEdgeResolver, error) {
	l := make([]*categoriesEdgeResolver, len(r.categories))
	for i := range l {
		l[i] = &categoriesEdgeResolver{
			ctx:    r.ctx,
			cursor: service.EncodeCursor(&(r.categories[i].ID)),
			model:  r.categories[i],
		}
	}
	return &l, nil
}

func (r *categoriesConnectionResolver) PageInfo() (*pageInfoResolver, error) {
	hasNext, _ := loader.CategoriesHasNextAfter(r.to)

	return newPageInfoResolver(
		r.ctx,
		r.from,
		r.to,
		hasNext,
	)
}
