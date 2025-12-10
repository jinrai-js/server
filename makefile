dev:
	@air

tests:
	@go test ./...


build:
	@go build -o ./dist/jinrai-dev-server cmd/jinrai-server/main.go