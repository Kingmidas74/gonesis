.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o ./web/scripts/application/engine/engine.wasm -tags js -tags wasm ./cmd/wasm/main.go
	#GOOS=js GOARCH=wasm go build -o ./web/scripts/application/engine/agent.wasm -tags js -tags wasm ./cmd/wasm/game/main.go

.PHONY: run
run:
	GOOS=js GOARCH=wasm go build -o ./web/scripts/application/engine/engine.wasm -tags js -tags wasm ./cmd/wasm/main.go
	#GOOS=js GOARCH=wasm go build -o ./web/scripts/application/engine/agent.wasm -tags js -tags wasm ./cmd/wasm/agent/main.go
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
