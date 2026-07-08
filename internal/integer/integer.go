package integer

// Add складывает два целых числа.
func Add(a, b int) int {
	return a + b
}

// Subtract вычитает второе число из первого.
func Subtract(a, b int) int {
	return a - b
}

// Multiply умножает два целых числа.
func Multiply(a, b int) int {
	return a * b
}

// Divide делит первое число на второе через целочисленное деление.
// При невозможности деления функция возвращает 0.
func Divide(a, b int) int {
	if b == 0 {
		return 0
	}

	return a / b
}

// Remainder возвращает остаток от деления первого числа на второе.
// При невозможности деления функция возвращает 0.
func Remainder(a, b int) int {
	if b == 0 {
		return 0
	}

	return a % b
}

// IsEven проверяет, является ли число чётным.
func IsEven(n int) bool {
	return n%2 == 0
}

// LastDigit возвращает последнюю цифру числа.
// Для отрицательных чисел результат должен быть положительным.
func LastDigit(n int) int {
	digit := n % 10

	if digit < 0 {
		return -digit
	}

	return digit
}

// Max возвращает большее из двух чисел.
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Min возвращает меньшее из двух чисел.
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Clamp ограничивает значение диапазоном [min, max].
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}

// SumThree складывает три целых числа.
func SumThree(a, b, c int) int {
	return a + b + c
}

// Average возвращает среднее арифметическое двух целых чисел.
func Average(a, b int) int {
	return (a + b) / 2
}

// IntToInt64 явно преобразует int в int64.
func IntToInt64(n int) int64 {
	return int64(n)
}

// NonNegativeToUint преобразует неотрицательное int-число в uint.
// Отрицательные значения дают 0.
func NonNegativeToUint(n int) uint {
	if n < 0 {
		return 0
	}

	return uint(n)
}

// CountPages считает количество страниц для списка элементов.
// Количество страниц считается с округлением вверх.
func CountPages(totalItems, pageSize int) int {
	if totalItems <= 0 || pageSize <= 0 {
		return 0
	}

	return (totalItems + pageSize - 1) / pageSize
}
