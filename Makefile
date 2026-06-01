build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
run:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
	chmod +x bin/hexlet-path-size
	./bin/hexlet-path-size
lint:
	golangci-lint run