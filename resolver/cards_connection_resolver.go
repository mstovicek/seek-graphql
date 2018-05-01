package resolver

import (
	"context"
	"github.com/mstovicek/seek-graphql/model"
	"log"
)

func newCardConnectionResolverByCategory(
	ctx context.Context,
	reader cardReaderInterface,
	category *model.Category,
	first int,
	afterCursor *string,
) (*cardsConnectionResolver, error) {
	afterID, err := decodeCursor(afterCursor)
	if err != nil {
		return nil, err
	}

	cards, err := reader.ListCardsByCategory(ctx, category, first, afterID)
	if err != nil {
		return nil, err
	}

	firstID := &(cards[0].ID)
	lastID := &(cards[len(cards)-1].ID)

	totalCount, err := reader.CardsTotalCountByCategory(category)
	if err != nil {
		return nil, err
	}

	hasNext, err := reader.CardsHasNextAfter(lastID)
	if err != nil {
		return nil, err
	}

	return &cardsConnectionResolver{
		ctx:        ctx,
		cards:      cards,
		totalCount: totalCount,
		from:       firstID,
		to:         lastID,
		hasNext:    hasNext,
	}, nil
}

type cardsConnectionResolver struct {
	ctx        context.Context
	cards      []*model.Card
	totalCount *int
	from       *string
	to         *string
	hasNext    bool
}

func (r *cardsConnectionResolver) TotalCount() (int32, error) {
	log.Printf("cardsConnectionResolver.TotalCount %v \n", r)

	return int32(*r.totalCount), nil
}

func (r *cardsConnectionResolver) Edges() (*[]*cardsEdgeResolver, error) {
	log.Printf("cardsConnectionResolver.Edges %v \n", r)

	l := make([]*cardsEdgeResolver, len(r.cards))
	for i := range l {
		l[i] = &cardsEdgeResolver{
			ctx:    r.ctx,
			cursor: encodeCursor(&(r.cards[i].ID)),
			model:  r.cards[i],
		}
	}
	return &l, nil
}

func (r *cardsConnectionResolver) PageInfo() (*pageInfoResolver, error) {
	log.Printf("cardsConnectionResolver.PageInfo %v \n", r)

	return newPageInfoResolver(
		r.ctx,
		r.from,
		r.to,
		r.hasNext,
	)
}
