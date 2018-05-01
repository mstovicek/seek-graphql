package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/model"
)

func newMeResolverByCtx(
	ctx context.Context,
	meReader meReaderInterface,
	categoryReader categoryReaderInterface,
	cardReader cardReaderInterface,
) (*meResolver, error) {
	me, err := meReader.LoadMeByCtx(ctx)
	if err != nil {
		return nil, err
	}

	return &meResolver{
		ctx: ctx,
		me:  me,
		categoryReader: categoryReader,
		cardReader: cardReader,
	}, nil
}

type meResolver struct {
	ctx context.Context
	me  *model.Me
	categoryReader categoryReaderInterface
	cardReader cardReaderInterface
}

func (r *meResolver) ID() (graphql.ID, error) {
	return graphql.ID(r.me.ID), nil
}

func (r *meResolver) Name() (*string, error) {
	return &r.me.Name, nil
}

func (r *meResolver) Email() (string, error) {
	return r.me.Email, nil
}

func (r *meResolver) Categories(ctx context.Context, args struct {
	First *int32
	After *string
}) (*categoriesConnectionResolver, error) {
	first := 0
	after := ""

	firstPointer := args.First
	if firstPointer != nil {
		first = int(*firstPointer)
	}

	afterPointer := args.After
	if afterPointer != nil {
		after = *afterPointer
	}

	return newCategoriesConnectionResolver(
		ctx,
		r.categoryReader,
		r.cardReader,
		first,
		&after,
	)
}