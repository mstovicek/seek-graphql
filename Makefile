build: clean
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/sls-api-gateway cmd/sls-api-gateway.go
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/test cmd/test.go

clean:
	rm -f bin/*
