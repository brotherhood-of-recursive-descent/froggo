game:
	go run cmd/froggo/main.go


build:
	go build -ldflags="-X main.VERSION=$(git rev-parse --short HEAD)" cmd/froggo/main.go

test:
	go test -cover -v ./...

wasm:
	env GOOS=js GOARCH=wasm go build -o dist/froggo.wasm cmd/froggo/main.go

serve:
	python3 -m http.server 9000 --directory dist