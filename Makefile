build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
test:
	go test -v ./...
lint:
	golangci-lint run
lint-fix:
	golangci-lint run --fix