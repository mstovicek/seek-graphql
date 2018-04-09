package resolver

import (
	"golang.org/x/net/context"
	"log"
	"strconv"
)

func (r *Resolver) Categories(ctx context.Context, args struct {
	First *int32
	After *string
}) (*categoriesConnectionResolver, error) {
	first := 0
	after := ""

	firstPointer := args.First
	if firstPointer != nil {
		first = int(*firstPointer)
		log.Println("first: " + strconv.Itoa(first))
	}

	afterPointer := args.After
	if afterPointer != nil {
		after = *afterPointer
		log.Println("after: " + after)
	}

	log.Printf("categories: %d, %s\n", first, after)

	return newCategoriesConnectionResolver(
		ctx,
		first,
		&after,
	)
}
