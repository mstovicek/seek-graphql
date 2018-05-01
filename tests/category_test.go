package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategory(t *testing.T) {
	query := `query {
	me {
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
	}
}`

	expectedResponse := `{
	"data": {
		"me": {
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
	}
}`

	actualResponse, err := getResponse(query)

	assert.Nil(t, err)
	assertEqualResponse(t, expectedResponse, *actualResponse)
}
