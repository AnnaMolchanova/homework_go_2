GO ?= go
COVERAGE_FILE ?= coverage.out
COVERAGE_MIN ?= 80
PACKAGE_NAME ?= homework-go-2-linux-amd64

.PHONY: help deps-check mod-check fmt fmt-check vet compile test test-unit test-integration test-race coverage coverage-check build package clean run-all ci \
run-integer run-bases run-float run-boolean run-text run-constants run-conversion run-pointers run-calculator run-profile run-order \
test-integer test-bases test-float test-boolean test-text test-constants test-conversion test-pointers test-calculator test-profile test-order

help: ## Показать список команд
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_-]+:.*##/ {printf "%-24s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps-check: ## Скачать и проверить зависимости Go-модулей
	$(GO) mod download
	$(GO) mod verify

mod-check: ## Проверить, что go.mod/go.sum не меняются после go mod tidy
	$(GO) mod tidy
	@if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then \
		if [ -n "$$(git status --short -- go.mod go.sum)" ]; then \
			git status --short -- go.mod go.sum; \
			echo "go.mod/go.sum changed after go mod tidy"; \
			exit 1; \
		fi; \
	else \
		echo "not a git repository, skip go.mod/go.sum diff check"; \
	fi

fmt: ## Отформатировать Go-файлы
	$(GO) fmt ./...

fmt-check: ## Проверить gofmt без изменения файлов
	@test -z "$$(gofmt -l . | grep -v '^bin/')" || (echo "Run gofmt for files above" && gofmt -l . && exit 1)

vet: ## Запустить go vet
	$(GO) vet ./...

compile: ## Проверить, что все пакеты компилируются без запуска тестов
	$(GO) test -run '^$$' ./...

test: test-unit test-integration ## Запустить все тесты

test-unit: ## Запустить unit-тесты по internal-пакетам
	$(GO) test ./internal/...

test-integration: ## Запустить integration-тесты по cmd/main
	$(GO) test ./test/integration

test-race: ## Запустить тесты с race detector
	$(GO) test -race ./...

coverage: ## Посчитать coverage по internal-пакетам
	$(GO) test ./internal/... -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FILE)

coverage-check: coverage ## Проверить минимальный coverage
	@total=$$($(GO) tool cover -func=$(COVERAGE_FILE) | awk '/total:/ {gsub("%", "", $$3); print $$3}'); \
	awk -v total="$$total" -v min="$(COVERAGE_MIN)" 'BEGIN { if (total + 0 < min + 0) { printf("coverage %.1f%% is below %.1f%%\n", total, min); exit 1 } else { printf("coverage %.1f%% is OK\n", total) } }'

build: ## Собрать все main в bin/
	mkdir -p bin
	$(GO) build -o bin/01_integer ./cmd/01_integer
	$(GO) build -o bin/02_bases_bytes ./cmd/02_bases_bytes
	$(GO) build -o bin/03_float ./cmd/03_float
	$(GO) build -o bin/04_boolean ./cmd/04_boolean
	$(GO) build -o bin/05_text ./cmd/05_text
	$(GO) build -o bin/06_constants ./cmd/06_constants
	$(GO) build -o bin/07_conversion ./cmd/07_conversion
	$(GO) build -o bin/08_pointers ./cmd/08_pointers
	$(GO) build -o bin/09_calculator ./cmd/09_calculator
	$(GO) build -o bin/10_profile ./cmd/10_profile
	$(GO) build -o bin/11_order ./cmd/11_order

package: build ## Упаковать собранные бинарники в tar.gz
	tar -czf bin/$(PACKAGE_NAME).tar.gz \
		-C bin \
		01_integer \
		02_bases_bytes \
		03_float \
		04_boolean \
		05_text \
		06_constants \
		07_conversion \
		08_pointers \
		09_calculator \
		10_profile \
		11_order

clean: ## Удалить временные файлы
	rm -rf bin $(COVERAGE_FILE)

run-all: ## Запустить все main по очереди
	$(GO) run ./cmd/01_integer
	$(GO) run ./cmd/02_bases_bytes
	$(GO) run ./cmd/03_float
	$(GO) run ./cmd/04_boolean
	$(GO) run ./cmd/05_text
	$(GO) run ./cmd/06_constants
	$(GO) run ./cmd/07_conversion
	$(GO) run ./cmd/08_pointers
	$(GO) run ./cmd/09_calculator
	$(GO) run ./cmd/10_profile
	$(GO) run ./cmd/11_order

run-integer: ## Запустить 01_integer
	$(GO) run ./cmd/01_integer

run-bases: ## Запустить 02_bases_bytes
	$(GO) run ./cmd/02_bases_bytes

run-float: ## Запустить 03_float
	$(GO) run ./cmd/03_float

run-boolean: ## Запустить 04_boolean
	$(GO) run ./cmd/04_boolean

run-text: ## Запустить 05_text
	$(GO) run ./cmd/05_text

run-constants: ## Запустить 06_constants
	$(GO) run ./cmd/06_constants

run-conversion: ## Запустить 07_conversion
	$(GO) run ./cmd/07_conversion

run-pointers: ## Запустить 08_pointers
	$(GO) run ./cmd/08_pointers

run-calculator: ## Запустить 09_calculator
	$(GO) run ./cmd/09_calculator

run-profile: ## Запустить 10_profile
	$(GO) run ./cmd/10_profile

run-order: ## Запустить 11_order
	$(GO) run ./cmd/11_order

test-integer: ## Unit-тесты пакета internal/integer
	$(GO) test ./internal/integer

test-bases: ## Unit-тесты пакета internal/basesbytes
	$(GO) test ./internal/basesbytes

test-float: ## Unit-тесты пакета internal/float
	$(GO) test ./internal/float

test-boolean: ## Unit-тесты пакета internal/boolean
	$(GO) test ./internal/boolean

test-text: ## Unit-тесты пакета internal/text
	$(GO) test ./internal/text

test-constants: ## Unit-тесты пакета internal/constants
	$(GO) test ./internal/constants

test-conversion: ## Unit-тесты пакета internal/conversion
	$(GO) test ./internal/conversion

test-pointers: ## Unit-тесты пакета internal/pointers
	$(GO) test ./internal/pointers

test-calculator: ## Unit-тесты пакета internal/calculator
	$(GO) test ./internal/calculator

test-profile: ## Unit-тесты пакета internal/profile
	$(GO) test ./internal/profile

test-order: ## Unit-тесты пакета internal/order
	$(GO) test ./internal/order

ci: deps-check mod-check fmt-check vet test-unit test-integration test-race coverage-check build package ## Полная локальная проверка как в CI