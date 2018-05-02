package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env"
	"github.com/mstovicek/seek-graphql/schema"
	"log"
	"strings"
)

type config struct {
	CorsAllowOrigin  string `env:"CORS_ALLOW_ORIGIN" envDefault:"*"`
	CorsAllowMethods string `env:"CORS_ALLOW_METHODS" envDefault:"*"`
}

type executorInterface interface {
	Execute(ctx context.Context, query string) (interface{}, error)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	logContextRequest(ctx, request)

	executor, err := getSchemaExecutor()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	response, err := executor.Execute(ctx, request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

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

func getSchemaExecutor() (executorInterface, error) {
	return schema.NewSchemaExecutor()
}

func logContextRequest(ctx context.Context, request events.APIGatewayProxyRequest) {
	request.Body = strings.Replace(request.Body, "\n", "", -1)

	log.Printf("context: %v", ctx)
	log.Printf("request: %v", request)
	log.Printf("query: %s", request.Body)
}
