package loader

import (
	"context"
	"github.com/mstovicek/seek-2/model"
	"log"
	"strconv"
)

func LoadCard(ctx context.Context, ID string) (*model.Card, error) {
	log.Printf("loader.LocadCard (id: %s) !!\n", ID)

	return &model.Card{
		ID:    ID,
		Title: "card@" + ID,
	}, nil
}

func ListCardsByCategory(ctx context.Context, category *model.Category, first int, afterID *string) ([]*model.Card, error) {
	log.Printf("loader.ListCardsByCategory(first: %d, after: %s) !!\n", first, *afterID)

	afterInt, _ := strconv.Atoi(*afterID)
	slice := make([]*model.Card, first)

	for i := 0; i < first; i++ {
		id := afterInt + i
		slice[i] = &model.Card{
			ID:    strconv.Itoa(id),
			Title: category.Title + ":card@" + strconv.Itoa(id),
		}
	}

	return slice, nil
}

func CardsTotalCountByCategory(category *model.Category) (*int, error) {
	log.Printf("loader.CardsTotalCountByCategory (category: %s)!!\n", category.Title)

	totalCount := 999

	return &totalCount, nil
}

func CardsHasNextAfter(afterID *string) (bool, error) {
	return true, nil
}

func InsertCard(ctx context.Context, title string) (*model.Card, error) {
	return &model.Card{
		ID: "42",
		Title: "newCard:" + title,
	}, nil
}
