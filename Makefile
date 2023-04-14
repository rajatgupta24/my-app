lint:
	golangci-lint run ./... --verbose --build-tags build

test:
	go test ./... -v

build:
	go build .
