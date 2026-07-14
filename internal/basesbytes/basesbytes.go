package basesbytes

import "strconv"

// Decimal42 возвращает число 42 в десятичной записи.
func Decimal42() int {
	return 42
}

// Hex2A возвращает число 42 в шестнадцатеричной записи.
func Hex2A() int {
	return 0x2A
}

// Binary101010 возвращает число 42 в двоичной записи.
func Binary101010() int {
	return 0b101010
}

// Octal52 возвращает число 42 в восьмеричной записи.
func Octal52() int {
	return 0o52
}

// SameNumber проверяет, что 42 можно записать разными способами.
func SameNumber() bool {
	return Decimal42() == Hex2A() &&
		Decimal42() == Binary101010() &&
		Decimal42() == Octal52()
}

// MaxByte возвращает максимальное значение byte.
func MaxByte() byte {
	return 255
}

// FormatDecimal форматирует число в десятичной системе.
func FormatDecimal(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// FormatBinary форматирует число в двоичной системе.
func FormatBinary(n uint64) string {
	return strconv.FormatUint(n, 2)
}

// FormatOctal форматирует число в восьмеричной системе.
func FormatOctal(n uint64) string {
	return strconv.FormatUint(n, 8)
}

// FormatHex форматирует число в шестнадцатеричной системе.
func FormatHex(n uint64) string {
	return strconv.FormatUint(n, 16)
}

// FormatAllBases выводит одно число в разных системах счисления.
func FormatAllBases(n uint64) string {
	return "dec=" + FormatDecimal(n) +
		" bin=" + FormatBinary(n) +
		" oct=" + FormatOctal(n) +
		" hex=" + FormatHex(n)
}

// IsASCII проверяет, входит ли byte в ASCII-диапазон.
func IsASCII(b byte) bool {
	return b <= 127
}

// ToUpperASCII переводит маленькую ASCII-букву в большую.
func ToUpperASCII(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}

	return b
}

// ToLowerASCII переводит большую ASCII-букву в маленькую.
func ToLowerASCII(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}

	return b
}

// PackTwoBytes объединяет два byte в одно uint16.
func PackTwoBytes(high, low byte) uint16 {
	return uint16(high)<<8 | uint16(low)
}
