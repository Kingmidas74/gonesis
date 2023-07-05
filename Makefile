.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o ./web/scripts/services/data-access/wasm/engine.wasm -tags js -tags wasm ./cmd/wasm/main.go

.PHONY: run
run:
	GOOS=js GOARCH=wasm go build -o ./web/scripts/services/data-access/wasm/engine.wasm -tags js -tags wasm ./cmd/wasm/main.go
	go run ./cmd/server/main.go


.PHONY: server
server:
	docker-compose up --build --remove-orphans -d

.PHONE: clean
clean:
	docker-compose rm --stop --force

.PHONY: test
test:
	go test -v ./...
