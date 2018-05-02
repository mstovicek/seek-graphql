package schema

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/resolver"
)

func NewSchemaExecutor() (*schemaExecutor, error) {
	sch := schema
	s, err := graphql.ParseSchema(sch, &resolver.Resolver{})
	if err != nil {
		return nil, err
	}

	return &schemaExecutor{
		schema: s,
	}, nil
}

type schemaExecutor struct {
	schema *graphql.Schema
}

func (s *schemaExecutor) Execute(
	ctx context.Context,
	query string,
	operationName string,
	variables map[string]interface{},
) (interface{}, error) {
	return s.schema.Exec(ctx, query, operationName, variables), nil
}
