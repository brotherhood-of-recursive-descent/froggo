
build:
	go build -ldflags  -X 'main.VERSION=$(git rev-parse --short HEAD)'" -v ./...

test:
	go test -v ./...
	