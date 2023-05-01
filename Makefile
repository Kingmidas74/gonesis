.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o ./assets/engine.wasm -tags js -tags wasm ./cmd/wasm/main.go

.PHONY: server
server:
	GOOS=js GOARCH=wasm go build -o ./assets/engine.wasm -tags js -tags wasm ./cmd/wasm/main.go
	go run ./cmd/server/main.go

.PHONY: test
test:
	go test -v ./...
