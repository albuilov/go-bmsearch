.PHONY: help build run test bench fmt vet lint tidy clean search-naive search-horspool search-bm

BINARY := bmsearch
BIN_DIR := bin
CMD_DIR := ./cmd/bmsearch

help: ## Показать список доступных команд
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Собрать бинарник в ./bin/bmsearch
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) $(CMD_DIR)

run: build ## Собрать и запустить с --help
	./$(BIN_DIR)/$(BINARY) --help

test: ## Запустить все тесты
	go test -v -race ./...

bench: ## Запустить все бенчмарки с замером аллокаций
	go test -bench=. -benchmem ./...

fmt: ## Отформатировать код
	go fmt ./...

vet: ## Проверить код через go vet
	go vet ./...

lint: ## Запустить golangci-lint (требует установки)
	golangci-lint run ./...

tidy: ## Привести go.mod / go.sum в порядок
	go mod tidy

clean: ## Удалить артефакты сборки
	rm -rf $(BIN_DIR)

search-naive: build ## Пример: поиск наивным алгоритмом
	./$(BIN_DIR)/$(BINARY) search -p "abc" -t "abcXabcYabc" -a naive

search-horspool: build ## Пример: поиск алгоритмом Horspool
	./$(BIN_DIR)/$(BINARY) search -p "abc" -t "abcXabcYabc" -a horspool

search-bm: build ## Пример: поиск полным Boyer-Moore
	./$(BIN_DIR)/$(BINARY) search -p "abc" -t "abcXabcYabc" -a bm