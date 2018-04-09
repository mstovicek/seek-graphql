package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-2/model"
	"context"
	"github.com/mstovicek/seek-2/loader"
)

type cardResolver struct {
	ctx  context.Context
	card *model.Card
}

func newCardResolverById(ctx context.Context, ID string) (*cardResolver, error) {
	card, _ := loader.LoadCard(ctx, ID)

	return &cardResolver{
		ctx:  ctx,
		card: card,
	}, nil
}

func newCardResolvedWithModel(ctx context.Context, card *model.Card) (*cardResolver, error) {
	return &cardResolver{
		ctx:  ctx,
		card: card,
	}, nil
}

func (r *cardResolver) ID() graphql.ID {
	return graphql.ID(r.card.ID)
}

func (r *cardResolver) Title() *string {
	return &r.card.Title
}
