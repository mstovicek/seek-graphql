package resolver

import (
	"context"
	"github.com/mstovicek/seek-graphql/loader_dummy"
	"github.com/mstovicek/seek-graphql/model"
)

func getCardReader(ctx context.Context) (cardReaderInterface, error) {
	return loader_dummy.NewCardLoader(ctx)
}

type cardReaderInterface interface {
	LoadCard(ctx context.Context, ID string) (*model.Card, error)
	ListCardsByCategory(ctx context.Context, category *model.Category, first int, afterID *string) ([]*model.Card, error)
	CardsTotalCountByCategory(category *model.Category) (*int, error)
	CardsHasNextAfter(afterID *string) (bool, error)
}

func getCardWriter(ctx context.Context) (cardWriterInterface, error) {
	return loader_dummy.NewCardLoader(ctx)
}

type cardWriterInterface interface {
	InsertCard(ctx context.Context, title string) (*model.Card, error)
}

func getCategoryReader(ctx context.Context) (categoryReaderInterface, error) {
	return loader_dummy.NewCategoryLoader(ctx)
}

type categoryReaderInterface interface {
	LoadCategoryById(ctx context.Context, ID string) (*model.Category, error)
	ListCategories(ctx context.Context, first int, afterID *string) ([]*model.Category, error)
	CategoriesHasNextAfter(after *string) (bool, error)
	CategoriesTotalCount() (*int, error)
}
