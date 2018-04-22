package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env"
	"github.com/mstovicek/seek-graphql/schema"
	"log"
)

type config struct {
	CorsAllowOrigin  string `env:"CORS_ALLOW_ORIGIN" envDefault:"*"`
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
		Body: string(rJSON),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  cfg.CorsAllowOrigin,
			"Access-Control-Allow-Methods": cfg.CorsAllowMethods,
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
