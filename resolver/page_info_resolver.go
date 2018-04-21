package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"context"
	"github.com/mstovicek/seek-graphql/service"
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
		startCursor: service.EncodeCursor(startID),
		endCursor:   service.EncodeCursor(endID),
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
