.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u -v github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

.PHONE: lint
lint:
	go get -u -v github.com/golang/lint/golint
	for file in $(shell find . -name '*.go' -not -path './vendor/*'); do golint $${file}; done

.PHONY: vet
vet:
	go vet ./...

build: clean deps
	env GOOS=linux go build -ldflags="-s -w" -o bin/sls-api-gateway cmd/sls-api-gateway.go
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/test cmd/test.go

clean:
	rm -f bin/*
