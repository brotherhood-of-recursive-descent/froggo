game:
	go run cmd/froggo/main.go


build:
	go build -ldflags="-X main.VERSION=$(git rev-parse --short HEAD)" cmd/froggo/main.go

test:
	go test -cover -v ./...
