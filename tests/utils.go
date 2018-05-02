package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mstovicek/seek-graphql/schema"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func getResponse(query string) (*string, error) {
	executor, err := schema.NewSchemaExecutor()
	if err != nil {
		return nil, err
	}

	response, err := executor.Execute(context.Background(), query, "", nil)
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
