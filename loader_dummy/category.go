package loader_dummy

import (
	"context"
	"github.com/mstovicek/seek-graphql/model"
	"log"
	"strconv"
)

func NewCategoryLoader(ctx context.Context) (*categoryLoader, error) {
	return &categoryLoader{
		ctx: ctx,
	}, nil
}

type categoryLoader struct {
	ctx context.Context
}

func (loader *categoryLoader) LoadCategoryById(ctx context.Context, ID string) (*model.Category, error) {
	log.Printf("loader.LoadCategoryById (id: %s) !!\n", ID)

	return &model.Category{
		ID:    ID,
		Title: "category@" + ID,
	}, nil
}

func (loader *categoryLoader) ListCategories(ctx context.Context, first int, afterID *string) ([]*model.Category, error) {
	log.Printf("loader.ListCategories(first: %d, after: %s) !!\n", first, *afterID)

	afterInt, _ := strconv.Atoi(*afterID)
	slice := make([]*model.Category, first)

	for i := 0; i < first; i++ {
		id := afterInt + i
		slice[i] = &model.Category{
			ID:    strconv.Itoa(id),
			Title: "category@" + strconv.Itoa(id),
		}
	}

	return slice, nil
}

func (loader *categoryLoader) CategoriesHasNextAfter(after *string) (bool, error) {
	return true, nil
}

func (loader *categoryLoader) CategoriesTotalCount() (*int, error) {
	log.Printf("loader.CategoriesTotalCount!!\n")

	totalCount := 999

	return &totalCount, nil
}
