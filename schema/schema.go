package schema

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/mstovicek/seek-graphql/resolver"
	"io/ioutil"
)

func Execute(ctx context.Context, query string) (*graphql.Response, error) {
	//sch, _ := getSchemaFromFile("./schema/schema.graphql")
	sch := getSchema()
	s := graphql.MustParseSchema(sch, &resolver.Resolver{})

	return s.Exec(ctx, query, "", make(map[string]interface{})), nil
}

func getSchemaFromFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func getSchema() string {
	return `
schema {
    query: Query
	mutation: Mutation
}

type Query {
    me: Me
}

type PageInfo {
    startCursor: ID
    endCursor: ID
    hasNextPage: Boolean!
}

type Me {
    id: ID!
    email: String!
    name: String
    categories(first: Int,  after: String): CategoriesConnection!
}

type Category {
    id: ID!
    title: String
    cards(first: Int, after: String): CardsConnection
}

type CategoriesEdge {
    cursor: ID!
    node: Category
}

type CategoriesConnection {
    totalCount: Int!
    edges: [CategoriesEdge]
    pageInfo: PageInfo!
}

type Card {
    id: ID!
    title: String
}

type CardsEdge {
    cursor: ID!
    node: Card
}

type CardsConnection {
    totalCount: Int!
    edges: [CardsEdge]
    pageInfo: PageInfo!
}

type Mutation {
	addCard(input: CardInput!): Card
}

input CardInput {
	title: String!
}
`
}
