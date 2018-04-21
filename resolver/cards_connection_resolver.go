package resolver

import (
	"github.com/mstovicek/seek-graphql/model"
	"github.com/mstovicek/seek-graphql/service"
	"log"
	"context"
	"github.com/mstovicek/seek-graphql/loader"
)

type cardsConnectionResolver struct {
	ctx        context.Context
	cards      []*model.Card
	totalCount *int
	from       *string
	to         *string
}

func newCardConnectionResolverByCategory(
	ctx context.Context,
	category *model.Category,
	first int,
	afterCursor *string,
) (*cardsConnectionResolver, error) {
	afterID, _ := service.DecodeCursor(afterCursor)

	cards, _ := loader.ListCardsByCategory(ctx, category, first, afterID)
	totalCount, _ := loader.CardsTotalCountByCategory(category)

	return &cardsConnectionResolver{
		ctx:        ctx,
		cards:      cards,
		totalCount: totalCount,
		from:       &(cards[0].ID),
		to:         &(cards[len(cards)-1].ID),
	}, nil
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
			cursor: service.EncodeCursor(&(r.cards[i].ID)),
			model:  r.cards[i],
		}
	}
	return &l, nil
}

func (r *cardsConnectionResolver) PageInfo() (*pageInfoResolver, error) {
	log.Printf("cardsConnectionResolver.PageInfo %v \n", r)

	hasNext, _ := loader.CardsHasNextAfter(r.to)

	return newPageInfoResolver(
		r.ctx,
		r.from,
		r.to,
		hasNext,
	)
}
