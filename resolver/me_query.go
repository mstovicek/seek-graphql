package resolver

import (
	"golang.org/x/net/context"
)

func (r *Resolver) Me(ctx context.Context) (*meResolver, error) {
	meReader, err := getMeReader(ctx)
	if err != nil {
		return nil, err
	}

	categoryReader, err := getCategoryReader(ctx)
	if err != nil {
		return nil, err
	}

	cardReader, err := getCardReader(ctx)
	if err != nil {
		return nil, err
	}

	return newMeResolverByCtx(ctx, meReader, categoryReader, cardReader)
}
