package resolver

import (
	"context"
	"github.com/mstovicek/seek-graphql/model"
	"github.com/mstovicek/seek-graphql/service"
)

func newCategoriesConnectionResolver(
	ctx context.Context,
	categoryReader categoryReaderInterface,
	cardReader cardReaderInterface,
	first int,
	afterCursor *string,
) (*categoriesConnectionResolver, error) {
	afterID, _ := service.DecodeCursor(afterCursor)

	categories, err := categoryReader.ListCategories(ctx, first, afterID)
	if err != nil {
		return nil, err
	}

	firstID := &(categories[0].ID)
	lastID := &(categories[len(categories)-1].ID)

	totalCount, _ := categoryReader.CategoriesTotalCount()
	hasNext, _ := categoryReader.CategoriesHasNextAfter(lastID)

	return &categoriesConnectionResolver{
		ctx:        ctx,
		cardReader: cardReader,
		categories: categories,
		totalCount: totalCount,
		from:       firstID,
		to:         lastID,
		hasNext:    hasNext,
	}, nil
}

type categoriesConnectionResolver struct {
	ctx        context.Context
	cardReader cardReaderInterface
	categories []*model.Category
	totalCount *int
	from       *string
	to         *string
	hasNext    bool
}

func (r *categoriesConnectionResolver) TotalCount() (int32, error) {
	return int32(*r.totalCount), nil
}

func (r *categoriesConnectionResolver) Edges() (*[]*categoriesEdgeResolver, error) {
	l := make([]*categoriesEdgeResolver, len(r.categories))
	for i := range l {
		l[i] = &categoriesEdgeResolver{
			ctx:        r.ctx,
			cardReader: r.cardReader,
			cursor:     service.EncodeCursor(&(r.categories[i].ID)),
			model:      r.categories[i],
		}
	}
	return &l, nil
}

func (r *categoriesConnectionResolver) PageInfo() (*pageInfoResolver, error) {
	return newPageInfoResolver(
		r.ctx,
		r.from,
		r.to,
		r.hasNext,
	)
}
