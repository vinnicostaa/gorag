APP=gorag

.PHONY: build run test tidy fmt vet clean

build:
	go build -o bin/$(APP) ./cmd/gorag

run:
	go run ./cmd/gorag

test:
	go test ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf bin dist tmp .vaultgraph
