package main

import (
	"github.com/stretchr/testify/assert"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mstovicek/seek-graphql/schema"
	"testing"
	"log"
)

func TestMe(t *testing.T) {
	query := `query {
	me {
		id,
		email,
		name
	}
}`

	expectedResponse := `{
	"data": {
		"me": {
			"id": "42",
			"email": "milan@me",
			"name": "Milan"
		}
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, *actualResponse)
}

func getResponse(query string) (*string, error) {
	response, err := schema.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}

	rJSON, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return nil, err
	}


	r := fmt.Sprintf("%s", rJSON)

	log.Printf("query: %s \n response: %s \n\n", query, r)

	return &r, nil
}

func aaa() {
	//	query := `
	//{
	//	card(id: "milan") {id,title}
	//}
	//`

	//	query := `
	//	{
	//		category(id: "milan") {
	//			id,
	//			title,
	//			cards (first: 3, after: "cursor:5") {
	//				totalCount,
	//				edges {
	//					cursor,
	//					node {
	//						id,
	//						title
	//					}
	//				}
	//				pageInfo{
	//					startCursor,
	//					endCursor,
	//					hasNextPage
	//				},
	//			}
	//		}
	//	}
	//`

	//	query := `
	//{
	//	categories (first: 2, after: "cursor:3"){
	//		totalCount,
	//		edges{
	//			cursor,
	//			node{
	//				id,
	//				title
	//				cards {
	//					totalCount
	//				}
	//			}
	//		},
	//		pageInfo{
	//			startCursor,
	//			endCursor,
	//			hasNextPage
	//		}
	//	}
	//}
	//`

	//	query := `
	//	mutation {
	//		addCard(input: {title: "NewCard"}) {
	//			id,
	//			title
	//		}
	//	}
	//`
}