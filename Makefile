build:
	@go build -ldflags="-w -s" -o ./bin/api ./cmd/api/main.go

run:
	@air -c .air.toml

test:
	@./scripts/test.sh

install-deps:
	go install github.com/cosmtrek/air@latest
	go mod tidy

setup: install-deps

clean:
	rm -rf ./bin ./tmp ./coverage.txt
