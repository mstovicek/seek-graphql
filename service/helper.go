package service

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"strings"
)

func EncodeCursor(i *string) graphql.ID {
	return graphql.ID(fmt.Sprintf("cursor:%s", *i))
	//return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%s", *i))))
}

func DecodeCursor(after *string) (*string, error) {
	i := strings.TrimPrefix(*after, "cursor:")
	return &i, nil

	//var decodedValue string
	//if after != nil {
	//	b, err := base64.StdEncoding.DecodeString(*after)
	//	if err != nil {
	//		return nil, err
	//	}
	//	i := strings.TrimPrefix(string(b), "cursor:")
	//	decodedValue = i
	//}
	//return &decodedValue, nil
}
