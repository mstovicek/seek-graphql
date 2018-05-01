package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/model"
)

func newCardResolverByID(
	ctx context.Context,
	reader cardReaderInterface,
	ID string,
) (*cardResolver, error) {
	card, err := reader.LoadCardByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &cardResolver{
		ctx:  ctx,
		card: card,
	}, nil
}

func newCardResolverByInput(
	ctx context.Context,
	writer cardWriterInterface,
	title string,
) (*cardResolver, error) {
	card, err := writer.InsertCard(ctx, title)
	if err != nil {
		return nil, err
	}

	return &cardResolver{
		ctx:  ctx,
		card: card,
	}, nil
}

func newCardResolvedWithModel(
	ctx context.Context,
	card *model.Card,
) (*cardResolver, error) {
	return &cardResolver{
		ctx:  ctx,
		card: card,
	}, nil
}

type cardResolver struct {
	ctx  context.Context
	card *model.Card
}

func (r *cardResolver) ID() graphql.ID {
	return graphql.ID(r.card.ID)
}

func (r *cardResolver) Title() *string {
	return &r.card.Title
}
