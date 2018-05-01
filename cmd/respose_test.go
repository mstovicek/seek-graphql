package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mstovicek/seek-graphql/schema"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
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
	assertEqualResponse(t, expectedResponse, *actualResponse)
}

func TestMeWithCategory(t *testing.T) {
	query := `query {
	me {
		id,
		email,
		name,
		categories (first: 1) {
			totalCount,
			edges{
				cursor,
				node {
					id,
					title
				}
			},
			pageInfo{
				startCursor,
				endCursor,
				hasNextPage
			}
		}
	}
}`

	expectedResponse := `{
	"data": {
		"me": {
			"id": "42",
			"email": "milan@me",
			"name": "Milan",
			"categories": {
				"totalCount": 999,
				"edges": [
					{
						"cursor": "cursor:0",
						"node": {
							"id": "0",
							"title": "category@0"
						}
					}
				],
				"pageInfo": {
					"startCursor": "cursor:0",
					"endCursor": "cursor:0",
					"hasNextPage": true
				}
			}
		}
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assertEqualResponse(t, expectedResponse, *actualResponse)
}

func TestMeWithCategoryAndCard(t *testing.T) {
	query := `query {
	me {
		id,
		email,
		name,
		categories (first: 1) {
			totalCount,
			edges{
				cursor,
				node {
					id,
					title,
					cards (first: 1) {
						totalCount,
						edges {
							cursor,
							node {
								id,
								title
							}
						}
						pageInfo{
							startCursor,
							endCursor,
							hasNextPage
						},
					}
				}
			},
			pageInfo{
				startCursor,
				endCursor,
				hasNextPage
			}
		}
	}
}`

	expectedResponse := `{
	"data": {
		"me": {
			"id": "42",
			"email": "milan@me",
			"name": "Milan",
			"categories": {
				"totalCount": 999,
				"edges": [
					{
						"cursor": "cursor:0",
						"node": {
							"id": "0",
							"title": "category@0",
							"cards": {
								"totalCount": 999,
								"edges": [
									{
										"cursor": "cursor:0",
										"node": {
											"id": "0",
											"title": "category@0:card@0"
										}
									}
								],
								"pageInfo": {
									"startCursor": "cursor:0",
									"endCursor": "cursor:0",
									"hasNextPage": true
								}
							}
						}
					}
				],
				"pageInfo": {
					"startCursor": "cursor:0",
					"endCursor": "cursor:0",
					"hasNextPage": true
				}
			}
		}
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assertEqualResponse(t, expectedResponse, *actualResponse)
}

func TestCard(t *testing.T) {
	query := `query {
	card(id: "3") {
		id,
		title
	}
}`

	expectedResponse := `{
	"data": {
		"card": {
			"id": "3",
			"title": "card@3"
		}
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assertEqualResponse(t, expectedResponse, *actualResponse)
}

func TestAddCard(t *testing.T) {
	query := `mutation {
	addCard(input: {title: "NewTitle"}) {
		id,
		title
	}
}`

	expectedResponse := `{
	"data": {
		"addCard": {
			"id": "42",
			"title": "newCard:NewTitle"
		}
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assertEqualResponse(t, expectedResponse, *actualResponse)
}

func TestCategory(t *testing.T) {
	query := `query {
	category(id: "4") {
		id,
		title,
		cards (first: 1) {
			totalCount,
			edges {
				cursor,
				node {
					id
				}
			}
			pageInfo{
				startCursor,
				endCursor,
				hasNextPage
			},
		}
	}
}`

	expectedResponse := `{
	"data": {
		"category": {
			"id": "4",
			"title": "category@4",
			"cards": {
				"totalCount": 999,
				"edges": [
					{
						"cursor": "cursor:0",
						"node": {
							"id": "0"
						}
					}
				],
				"pageInfo": {
					"startCursor": "cursor:0",
					"endCursor": "cursor:0",
					"hasNextPage": true
				}
			}
		}
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assertEqualResponse(t, expectedResponse, *actualResponse)
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

func assertEqualResponse(t assert.TestingT, expected string, actual string) bool {
	a := strings.Split(expected, "\n")
	b := strings.Split(actual, "\n")

	return assert.Equal(t, a, b)
}

func aaa() {
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
}
