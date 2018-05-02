package api

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func NewLoggerStdout() (*loggerStdout, error) {
	return &loggerStdout{}, nil
}

type loggerStdout struct {
}

func (l *loggerStdout) Info(ctx context.Context, message string, fields map[string]interface{}) {
	l.log("INFO", ctx, message, fields)
}

func (l *loggerStdout) Error(ctx context.Context, message string, fields map[string]interface{}) {
	l.log("ERROR", ctx, message, fields)
}

func (l *loggerStdout) log(prefix string, ctx context.Context, message string, fields map[string]interface{}) {
	log.Printf(
		"%s: %s, fields: %s",
		prefix,
		l.oneLine(message),
		l.fieldsAsString(fields),
	)
}

func (l *loggerStdout) oneLine(message string) string {
	return strings.Replace(message, "\n", "", -1)
}

func (l *loggerStdout) fieldsAsString(fields map[string]interface{}) string {
	var fieldsStr string

	for key, value := range fields {
		fieldsStr += l.oneLine(key + ": " + fmt.Sprintf("%s", value) + ", ")
	}

	return fieldsStr
}
