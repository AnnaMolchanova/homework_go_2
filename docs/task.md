# Задание 2: переменные, типы, память

Нужно реализовать функции с `TODO` внутри `internal/*`.

Для каждой темы есть:

- `cmd/<topic>/main.go` — только вызывает `Example()`;
- `internal/<topic>/example.go` — собирает демонстрационный вывод;
- `internal/<topic>/<topic>.go` — функции, которые нужно реализовать;
- `internal/<topic>/<topic>_test.go` — unit-тесты;
- `test/integration/homework_test.go` — проверка вывода программ.

## 01. Integer

Файл: `internal/integer/integer.go`

Реализовать минимум 15 задач:

1. `Add` — сложение двух `int`.
2. `Subtract` — вычитание.
3. `Multiply` — умножение.
4. `Divide` — целочисленное деление.
5. `Remainder` — остаток от деления.
6. `IsEven` — проверка чётности.
7. `LastDigit` — последняя цифра числа, включая отрицательные числа.
8. `Max` — большее из двух чисел.
9. `Min` — меньшее из двух чисел.
10. `Clamp` — ограничить значение диапазоном.
11. `SumThree` — сумма трёх чисел.
12. `Average` — среднее через целочисленное деление.
13. `IntToInt64` — явное преобразование `int` в `int64`.
14. `NonNegativeToUint` — отрицательные числа превращать в `0`, остальные — в `uint`.
15. `CountPages` — количество страниц для `totalItems` и `pageSize`.

Проверка:

```bash
make test-integer
make run-integer
```

## 02. Bases / Byte

Файл: `internal/basesbytes/basesbytes.go`

1. `Decimal42`
2. `Hex2A`
3. `Binary101010`
4. `Octal52`
5. `SameNumber`
6. `MaxByte`
7. `FormatDecimal`
8. `FormatBinary`
9. `FormatOctal`
10. `FormatHex`
11. `FormatAllBases`
12. `IsASCII`
13. `ToUpperASCII`
14. `ToLowerASCII`
15. `PackTwoBytes`

Проверка:

```bash
make test-bases
make run-bases
```

## 03. Float

Файл: `internal/float/float.go`

1. `Add`
2. `Subtract`
3. `Multiply`
4. `Divide`
5. `DiscountPrice`
6. `AddTax`
7. `CelsiusToFahrenheit`
8. `FahrenheitToCelsius`
9. `Average`
10. `Round2`
11. `FormatPrice`
12. `PercentOf`
13. `GrowthPercent`
14. `IsPositive`
15. `FloatToInt`

Проверка:

```bash
make test-float
make run-float
```

## 04. Boolean

Файл: `internal/boolean/boolean.go`

1. `CanEnter`
2. `IsAdult`
3. `CanBuyAlcohol`
4. `CanRest`
5. `IsWorkingDay`
6. `HasAccess`
7. `CanApplyDiscount`
8. `ShouldNotify`
9. `IsValidScore`
10. `IsInRange`
11. `IsLeapYear`
12. `CanWithdraw`
13. `LoginAllowed`
14. `IsEmpty`
15. `Not`

Проверка:

```bash
make test-boolean
make run-boolean
```

## 05. Text

Файл: `internal/text/text.go`

1. `ByteLen`
2. `RuneLen`
3. `FirstRune`
4. `LastRune`
5. `Trim`
6. `ToLower`
7. `ToUpper`
8. `NormalizeEmail`
9. `ContainsWord`
10. `ReplaceFirstRune`
11. `ReverseRunes`
12. `Initials`
13. `RepeatWord`
14. `JoinWithComma`
15. `IsPalindrome`

Проверка:

```bash
make test-text
make run-text
```

## 06. Constants / iota / switch

Файл: `internal/constants/constants.go`

1. `AppName`
2. `MaxAttempts`
3. `StatusText`
4. `IsFinalStatus`
5. `NextStatus`
6. `RoleText`
7. `CanEdit`
8. `HTTPStatusText`
9. `DayType`
10. `PriorityText`
11. `IsKnownStatus`
12. `PaymentStateText`
13. `TrafficLightAction`
14. `GradeText`
15. `EventTypeText`

Проверка:

```bash
make test-constants
make run-constants
```

## 07. Conversion / named types / strconv

Файл: `internal/conversion/conversion.go`

1. `IntToInt64`
2. `Int64ToInt`
3. `RubToKop`
4. `KopToRub`
5. `UserIDToString`
6. `ParseInt`
7. `ParseAndDouble`
8. `IntToString`
9. `FloatToString`
10. `ParseBoolText`
11. `BoolToText`
12. `SumIntAndInt64`
13. `PriceRubStringToKop`
14. `SafeParsePositive`
15. `FormatUser`

Проверка:

```bash
make test-conversion
make run-conversion
```

## 08. Pointers

Файл: `internal/pointers/pointers.go`

1. `ValueOrDefault`
2. `Increment`
3. `SetValue`
4. `Swap`
5. `ResetToZero`
6. `AddToValue`
7. `MaxPointer`
8. `IsNil`
9. `CopyValue`
10. `DoubleInPlace`
11. `NewInt`
12. `DivideInto`
13. `ApplyDiscountInPlace`
14. `ChoosePointer`
15. `PointToLarger`

Проверка:

```bash
make test-pointers
make run-pointers
```

## 09–11. Объединяющие задачи

### Calculator

Файл: `internal/calculator/calculator.go`

`Calculate(a, b int, operation string) (int, error)` должен поддерживать операции `+`, `-`, `*`, `/`, `%`.

Проверка:

```bash
make test-calculator
make run-calculator
```

### Profile

Файл: `internal/profile/profile.go`

`BuildUserCard(name string, age int, active bool) string` объединяет строки, числа, bool и условия.

Проверка:

```bash
make test-profile
make run-profile
```

### Order

Файл: `internal/order/order.go`

`OrderSummary(status int, priceRub int, paid bool) string` объединяет `const/iota`, `switch`, `int`, `bool` и преобразование рублей в копейки.

Проверка:

```bash
make test-order
make run-order
```

## Финальная проверка

```bash
make ci
```
