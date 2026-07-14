package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

// Блок conversion закрепляет:
// явные преобразования типов, named types,
// strconv, парсинг строк, форматирование значений
// и обработку ошибок при преобразовании.

// Rub — пользовательский тип для рублей.
type Rub int

// Kop — пользовательский тип для копеек.
type Kop int

// UserID — пользовательский тип для идентификатора пользователя.
type UserID int64

// IntToInt64 явно преобразует int в int64.
func IntToInt64(n int) int64 {
	return int64(n)
}

// Int64ToInt явно преобразует int64 в int.
// Для этой учебной задачи переполнение не обрабатываем.
func Int64ToInt(n int64) int {
	return int(n)
}

// RubToKop переводит рубли в копейки.
func RubToKop(rub Rub) Kop {
	return Kop(rub * 100)
}

// KopToRub переводит копейки в рубли через целочисленное деление.
func KopToRub(kop Kop) Rub {
	return Rub(kop / 100)
}

// UserIDToString преобразует UserID в string.
func UserIDToString(id UserID) string {
	return strconv.FormatInt(int64(id), 10)
}

// ParseInt преобразует строку в int.
func ParseInt(text string) (int, error) {
	return strconv.Atoi(text)
}

// ParseAndDouble парсит строку в int и умножает результат на 2.
func ParseAndDouble(text string) (int, error) {
	n, err := ParseInt(text)
	if err != nil {
		return 0, err
	}
	return n * 2, nil
}

// IntToString преобразует int в string.
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// FloatToString форматирует float64 с двумя знаками после точки.
func FloatToString(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

// ParseBoolText преобразует строку в bool.
func ParseBoolText(text string) (bool, error) {
	return strconv.ParseBool(text)
}

// BoolToText преобразует bool в string.
func BoolToText(value bool) string {
	return strconv.FormatBool(value)
}

// SumIntAndInt64 складывает int и int64.
func SumIntAndInt64(a int, b int64) int64 {
	return int64(a) + b
}

// PriceRubStringToKop преобразует строку с рублями в копейки.
func PriceRubStringToKop(text string) (Kop, error) {
	rub, err := ParseInt(text)
	if err != nil {
		return 0, err
	}
	if rub < 0 {
		return 0, errors.New("negative price")
	}
	return RubToKop(Rub(rub)), nil
}

// SafeParsePositive безопасно парсит положительное число.
func SafeParsePositive(text string) int {
	n, err := ParseInt(text)
	if err != nil {
		return 0
	}
	if n <= 0 {
		return 0
	}
	return n
}

// FormatUser форматирует пользователя в строку.
func FormatUser(id UserID, name string) string {
	return "user:" + UserIDToString(id) + ":" + name
}
