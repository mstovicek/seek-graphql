package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
