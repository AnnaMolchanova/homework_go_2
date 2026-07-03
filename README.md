# Homework 2: переменные, типы и память

Домашка повторяет формат classroom2: одна тема — один `main`, один `internal` package и тесты.

## Структура

```text
homework_go_2
├── cmd
│   ├── 01_integer/main.go
│   ├── 02_bases_bytes/main.go
│   ├── 03_float/main.go
│   ├── 04_boolean/main.go
│   ├── 05_text/main.go
│   ├── 06_constants/main.go
│   ├── 07_conversion/main.go
│   ├── 08_pointers/main.go
│   ├── 09_calculator/main.go
│   ├── 10_profile/main.go
│   └── 11_order/main.go
├── internal
│   ├── integer
│   ├── basesbytes
│   ├── float
│   ├── boolean
│   ├── text
│   ├── constants
│   ├── conversion
│   ├── pointers
│   ├── calculator
│   ├── profile
│   └── order
├── test/integration
├── docs/task.md
├── Makefile
└── .github/workflows/ci.yml
```

## Правило проекта

`main.go` не содержит логику. Он только вызывает `Example()` из нужного `internal` package и печатает результат.

Логику нужно писать в функциях внутри `internal/<topic>`.

## Как работать

1. Откройте `docs/task.md`.
2. Выберите тему, например `integer`.
3. Откройте `internal/integer/integer.go`.
4. Реализуйте функции с `TODO`.
5. Запустите тесты конкретной темы.

```bash
make test-integer
make run-integer
```

Когда тема готова, переходите к следующей.

## Основные команды

```bash
make help
make compile
make test-unit
make test-integration
make coverage-check
make build
make ci
```

## Команды по отдельным темам

```bash
make run-integer
make test-integer

make run-bases
make test-bases

make run-float
make test-float

make run-boolean
make test-boolean

make run-text
make test-text

make run-constants
make test-constants

make run-conversion
make test-conversion

make run-pointers
make test-pointers

make run-calculator
make test-calculator

make run-profile
make test-profile

make run-order
make test-order
```

## Важно

В стартовой версии тесты падают специально. Это нормально: домашка считается готовой, когда `make ci` проходит локально и затем проходит в GitHub Actions.
