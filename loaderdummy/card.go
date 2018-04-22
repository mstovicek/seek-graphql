package loaderdummy

import (
	"context"
	"github.com/mstovicek/seek-graphql/model"
	"log"
	"strconv"
)

func NewCardLoader(ctx context.Context) (*cardLoader, error) {
	return &cardLoader{
		ctx: ctx,
	}, nil
}

type cardLoader struct {
	ctx context.Context
}

func (loader *cardLoader) LoadCardById(ctx context.Context, ID string) (*model.Card, error) {
	log.Printf("loader.LocadCard (id: %s) !!\n", ID)

	return &model.Card{
		ID:    ID,
		Title: "card@" + ID,
	}, nil
}

func (loader *cardLoader) ListCardsByCategory(ctx context.Context, category *model.Category, first int, afterID *string) ([]*model.Card, error) {
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

func (loader *cardLoader) CardsTotalCountByCategory(category *model.Category) (*int, error) {
	log.Printf("loader.CardsTotalCountByCategory (category: %s)!!\n", category.Title)

	totalCount := 999

	return &totalCount, nil
}

func (loader *cardLoader) CardsHasNextAfter(afterID *string) (bool, error) {
	return true, nil
}

func (loader *cardLoader) InsertCard(ctx context.Context, title string) (*model.Card, error) {
	return &model.Card{
		ID:    "42",
		Title: "newCard:" + title,
	}, nil
}
