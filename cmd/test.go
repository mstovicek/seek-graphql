package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mstovicek/seek-graphql/schema"
)

func main() {
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

	query := `
	query {
		me {
			id,
			email,
			name
		}
	}
`

	response, _ := schema.Execute(context.Background(), query)

	rJSON, _ := json.MarshalIndent(response, "", "\t")
	fmt.Printf("%s \n", query)
	fmt.Printf("%s \n", rJSON)
}
