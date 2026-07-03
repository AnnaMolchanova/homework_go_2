package basesbytes

// Блок basesbytes закрепляет системы счисления,
// byte как alias для uint8, форматирование чисел
// и простые операции с ASCII-байтами.
//
// В этом файле важно не просто получить правильные числа,
// а потренироваться записывать и выводить их в разных системах счисления.

// Decimal42 возвращает число 42 в десятичной записи.
//
// TODO: задайте число обычной десятичной записью.
func Decimal42() int {
	return 0
}

// Hex2A возвращает число 42 в шестнадцатеричной записи.
//
// TODO: задайте то же число через hex-литерал.
func Hex2A() int {
	return 0
}

// Binary101010 возвращает число 42 в двоичной записи.
//
// TODO: задайте то же число через binary-литерал.
func Binary101010() int {
	return 0
}

// Octal52 возвращает число 42 в восьмеричной записи.
//
// TODO: задайте то же число через octal-литерал.
func Octal52() int {
	return 0
}

// SameNumber проверяет, что 42 можно записать разными способами.
//
// TODO: сравните значения из функций выше и убедитесь, что это одно число.
func SameNumber() bool {
	return false
}

// MaxByte возвращает максимальное значение byte.
//
// TODO: верните верхнюю границу диапазона byte.
func MaxByte() byte {
	return 0
}

// FormatDecimal форматирует число в десятичной системе.
//
// TODO: преобразуйте число в строку в десятичном формате.
func FormatDecimal(n uint64) string {
	return ""
}

// FormatBinary форматирует число в двоичной системе.
//
// TODO: преобразуйте число в строку в двоичном формате.
func FormatBinary(n uint64) string {
	return ""
}

// FormatOctal форматирует число в восьмеричной системе.
//
// TODO: преобразуйте число в строку в восьмеричном формате.
func FormatOctal(n uint64) string {
	return ""
}

// FormatHex форматирует число в шестнадцатеричной системе.
//
// TODO: преобразуйте число в строку в hex-формате с маленькими буквами.
func FormatHex(n uint64) string {
	return ""
}

// FormatAllBases выводит одно число в разных системах счисления.
//
// TODO: соберите строку с decimal, binary, octal и hex представлениями числа.
// Точный формат результата смотрите в тестах.
func FormatAllBases(n uint64) string {
	return ""
}

// IsASCII проверяет, входит ли byte в ASCII-диапазон.
//
// TODO: определите, относится ли значение byte к ASCII.
func IsASCII(b byte) bool {
	return false
}

// ToUpperASCII переводит маленькую ASCII-букву в большую.
//
// TODO: изменяйте только маленькие латинские буквы, остальные byte возвращайте без изменений.
func ToUpperASCII(b byte) byte {
	return 0
}

// ToLowerASCII переводит большую ASCII-букву в маленькую.
//
// TODO: изменяйте только большие латинские буквы, остальные byte возвращайте без изменений.
func ToLowerASCII(b byte) byte {
	return 0
}

// PackTwoBytes объединяет два byte в одно uint16.
//
// TODO: соберите uint16 из старшего и младшего байта.
func PackTwoBytes(high, low byte) uint16 {
	return 0
}
