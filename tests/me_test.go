package tests

import (
	"github.com/stretchr/testify/assert"
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
