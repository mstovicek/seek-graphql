package schema

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/resolver"
)

func NewSchemaExecutor() (*schemaExecutor, error) {
	sch := schema
	s := graphql.MustParseSchema(sch, &resolver.Resolver{})

	return &schemaExecutor{
		schema: s,
	}, nil
}

type schemaExecutor struct {
	schema *graphql.Schema
}

func (s *schemaExecutor) Execute(ctx context.Context, query string) (interface{}, error) {
	return s.schema.Exec(ctx, query, "", make(map[string]interface{})), nil
}
