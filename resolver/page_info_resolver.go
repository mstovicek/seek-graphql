package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
)

type pageInfoResolver struct {
	ctx         context.Context
	startCursor graphql.ID
	endCursor   graphql.ID
	hasNextPage bool
}

func newPageInfoResolver(
	ctx context.Context,
	startID *string,
	endID *string,
	hasNextPage bool,
) (*pageInfoResolver, error) {
	return &pageInfoResolver{
		ctx:         ctx,
		startCursor: encodeCursor(startID),
		endCursor:   encodeCursor(endID),
		hasNextPage: hasNextPage,
	}, nil
}

func (r *pageInfoResolver) StartCursor() *graphql.ID {
	return &r.startCursor
}

func (r *pageInfoResolver) EndCursor() *graphql.ID {
	return &r.endCursor
}

func (r *pageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}
