package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/mstovicek/seek-2/schema"
	"encoding/json"
	"github.com/caarlos0/env"
)

type config struct {
	CorsAllowOrigin string `env:"CORS_ALLOW_ORIGIN" envDefault:"*"`
	CorsAllowMethods string `env:"CORS_ALLOW_METHODS" envDefault:"*"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

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
		Headers: map[string]string {
			"Access-Control-Allow-Origin": cfg.CorsAllowOrigin,
			"Access-Control-Allow-Methods": cfg.CorsAllowMethods,
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
