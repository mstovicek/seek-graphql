package resolver

import (
	"github.com/mstovicek/seek-graphql/model"
	"golang.org/x/net/context"
)

func (r *Resolver) AddCard(ctx context.Context, args struct {
	Input model.CardInput
}) (*cardResolver, error) {
	cardLoader, _ := getCardWriter(ctx)

	return newCardResolverByInput(ctx, cardLoader, args.Input.Title)
}
