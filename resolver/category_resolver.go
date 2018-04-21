package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/model"
	"log"
	"strconv"
	"github.com/mstovicek/seek-graphql/loader"
	"context"
)

type categoryResolver struct {
	ctx      context.Context
	category *model.Category
}

func newCategoryResolverById(ctx context.Context, ID string) (*categoryResolver, error) {
	category, _ := loader.LoadCategoryById(ctx, ID)

	return &categoryResolver{
		ctx:      ctx,
		category: category,
	}, nil
}

func newCategoryResolverWithModel(ctx context.Context, category *model.Category) (*categoryResolver, error) {
	return &categoryResolver{
		ctx:      ctx,
		category: category,
	}, nil
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
		r.category,
		first,
		&after,
	)
}

//func (r *categoryResolver) Cards(args struct {
//	First *int32
//	After *string
//}) *cardsConnectionResolver {
//	first := 10
//
//	firstPointer := args.First
//	if firstPointer != nil {
//		first = int(*firstPointer)
//		log.Println("first: " + strconv.Itoa(first))
//	}
//
//	after := ""
//
//	afterPointer := args.After
//	if afterPointer != nil {
//		after = *afterPointer
//		log.Println("after: " + after)
//	}
//
//	cards, _ := loader.ListCardsByCategory(r.ctx, first, after)
//
//	return newCardConnectionResolverWithModels(
//		r.ctx,
//		r.category,
//		cards,
//	)
//}
