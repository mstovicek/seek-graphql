package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/mstovicek/seek-2/schema"
	"encoding/json"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("request: %v, body: %s, context: %v", request, request.Body)
	response, _ := schema.Execute(ctx, request.Body)

	rJSON, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(rJSON),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
