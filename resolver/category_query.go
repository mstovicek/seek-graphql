package resolver

import (
	"golang.org/x/net/context"
)

func (r *Resolver) Category(ctx context.Context, args struct {
	ID string
}) (*categoryResolver, error) {
	categoryReader, err := getCategoryReader(ctx)
	if err != nil {
		return nil, err
	}

	cardReader, err := getCardReader(ctx)
	if err != nil {
		return nil, err
	}

	return newCategoryResolverByID(ctx, categoryReader, cardReader, args.ID)
}
