package resolver

import (
	"golang.org/x/net/context"
	"fmt"
	"github.com/mstovicek/seek-2/model"
)

func (r *Resolver) AddCard(ctx context.Context, args struct {
	Input model.CardInput
}) (*cardResolver, error) {
	fmt.Printf("input %v", args)
	return newCardResolverByInput(ctx, args.Input.Title)
}
