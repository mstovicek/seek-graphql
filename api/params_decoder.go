package api

import (
	"encoding/json"
	"strings"
)

type params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func NewParamsDecoder() (*paramsDecoder, error) {
	return &paramsDecoder{}, nil
}

type paramsDecoder struct {
}

func (p *paramsDecoder) DecodeParams(body string) (query string, operationName string, variables map[string]interface{}) {
	var params params

	decoder := json.NewDecoder(strings.NewReader(body))
	err := decoder.Decode(&params)
	if err != nil {
		params.Query = body
		params.OperationName = ""
		params.Variables = make(map[string]interface{})
	}

	return params.Query, params.OperationName, params.Variables
}
