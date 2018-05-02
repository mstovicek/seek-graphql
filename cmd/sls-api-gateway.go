package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mstovicek/seek-graphql/api"
	"github.com/mstovicek/seek-graphql/schema"
	"log"
)

type corsConfigInterface interface {
	GetAllowOrigin() string
	GetAllowMethods() string
}

type decoderInterface interface {
	DecodeParams(body string) (query string, operationName string, variables map[string]interface{})
}

type executorInterface interface {
	Execute(ctx context.Context, query string, operationName string, variables map[string]interface{}) (interface{}, error)
}

type loggerInterface interface {
	Info(ctx context.Context, message string, fields map[string]interface{})
	Error(ctx context.Context, message string, fields map[string]interface{})
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger, err := getLogger()
	if err != nil {
		log.Printf("ERROR: cannot get logger, error: %s", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	corsConfig, err := getCorsConfig()
	if err != nil {
		log.Printf("cannot get CORS config, error: %s", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	logger.Info(ctx, "request received", map[string]interface{}{
		"context": ctx,
		"request": request,
		"body":    request.Body,
	})

	executor, err := getSchemaExecutor()
	if err != nil {
		logger.Error(ctx, "cannot get schema executor", map[string]interface{}{
			"error": err,
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	decoder, err := getApiDecoder()
	if err != nil {
		logger.Error(ctx, "cannot get API decoder", map[string]interface{}{
			"error": err,
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	query, operationName, variables := decoder.DecodeParams(request.Body)
	response, err := executor.Execute(ctx, query, operationName, variables)
	if err != nil {
		logger.Error(ctx, "cannot execute query", map[string]interface{}{
			"error": err,
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	rJSON, err := json.Marshal(response)
	if err != nil {
		logger.Error(ctx, "cannot marshal response", map[string]interface{}{
			"error": err,
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	logger.Info(ctx, "query executed", map[string]interface{}{
		"response": string(rJSON),
	})

	return events.APIGatewayProxyResponse{
		Body: string(rJSON),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  corsConfig.GetAllowOrigin(),
			"Access-Control-Allow-Methods": corsConfig.GetAllowMethods(),
			"Content-type":                 "application/json",
		},
		StatusCode: 200,
	}, nil
}

func getSchemaExecutor() (executorInterface, error) {
	return schema.NewSchemaExecutor()
}

func getApiDecoder() (decoderInterface, error) {
	return api.NewParamsDecoder()
}

func getLogger() (loggerInterface, error) {
	return api.NewLoggerStdout()
}

func getCorsConfig() (corsConfigInterface, error) {
	return api.NewCorsConfig()
}
