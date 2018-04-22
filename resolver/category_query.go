package resolver

import (
	"golang.org/x/net/context"
)

func (r *Resolver) Category(ctx context.Context, args struct {
	ID string
}) (*categoryResolver, error) {
	categoryReader, _ := getCategoryReader(ctx)
	cardReader, _ := getCardReader(ctx)

	return newCategoryResolverById(ctx, categoryReader, cardReader, args.ID)
}
