package resolver

import (
	"golang.org/x/net/context"
)

func (r *Resolver) Category(ctx context.Context, args struct {
	ID string
}) (*categoryResolver, error) {
	return newCategoryResolverById(ctx, args.ID)
}
