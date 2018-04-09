package schema

import (
	"io/ioutil"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-2/resolver"
	"context"
)

func Execute(query string) (*graphql.Response, error) {
	sch, _ := getSchema("./schema/schema.graphql")
	s := graphql.MustParseSchema(sch, &resolver.Resolver{})

	ctx := context.Background()

	return s.Exec(ctx, query, "", make(map[string]interface{})), nil
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
