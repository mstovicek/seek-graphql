package resolver

import (
	"golang.org/x/net/context"
)

func (r *Resolver) Card(ctx context.Context, args struct {
	ID string
}) (*cardResolver, error) {
	cardLoader, err := getCardReader(ctx)
	if err != nil {
		return nil, err
	}

	return newCardResolverByID(ctx, cardLoader, args.ID)
}
