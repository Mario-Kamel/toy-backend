build:
	@go build -o bin/app.exe cmd/main.go
test:
	@go test -v ./...

run: build
	@./bin/app.exe