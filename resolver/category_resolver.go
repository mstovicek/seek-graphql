package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/model"
	"log"
	"strconv"
)

func newCategoryResolverById(
	ctx context.Context,
	categoryReader categoryReaderInterface,
	cardReader cardReaderInterface,
	ID string,
) (*categoryResolver, error) {
	category, _ := categoryReader.LoadCategoryById(ctx, ID)

	return newCategoryResolverWithModel(ctx, cardReader, category)
}

func newCategoryResolverWithModel(
	ctx context.Context,
	cardReader cardReaderInterface,
	category *model.Category,
) (*categoryResolver, error) {
	return &categoryResolver{
		ctx:        ctx,
		cardReader: cardReader,
		category:   category,
	}, nil
}

type categoryResolver struct {
	ctx        context.Context
	cardReader cardReaderInterface
	category   *model.Category
}

func (r *categoryResolver) ID() (graphql.ID, error) {
	return graphql.ID(r.category.ID), nil
}

func (r *categoryResolver) Title() (*string, error) {
	return &r.category.Title, nil
}

func (r *categoryResolver) Cards(args struct {
	First *int32
	After *string
}) (*cardsConnectionResolver, error) {
	first := 10

	firstPointer := args.First
	if firstPointer != nil {
		first = int(*firstPointer)
		log.Println("first: " + strconv.Itoa(first))
	}

	after := ""

	afterPointer := args.After
	if afterPointer != nil {
		after = *afterPointer
		log.Println("after: " + after)
	}

	return newCardConnectionResolverByCategory(
		r.ctx,
		r.cardReader,
		r.category,
		first,
		&after,
	)
}
