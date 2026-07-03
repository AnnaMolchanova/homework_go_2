package text

// Блок text закрепляет работу со строками:
// string, byte, rune, immutable string, Unicode,
// пакет strings и простые операции над текстом.

// ByteLen возвращает длину строки в байтах.
//
// TODO: посчитайте размер строки в байтах.
func ByteLen(s string) int {
	return 0
}

// RuneLen возвращает количество Unicode-символов в строке.
//
// TODO: посчитайте количество rune в строке.
func RuneLen(s string) int {
	return 0
}

// FirstRune возвращает первый Unicode-символ строки.
//
// TODO: безопасно верните первый символ строки с учётом Unicode.
func FirstRune(s string) string {
	return ""
}

// LastRune возвращает последний Unicode-символ строки.
//
// TODO: безопасно верните последний символ строки с учётом Unicode.
func LastRune(s string) string {
	return ""
}

// Trim убирает пробелы по краям строки.
//
// TODO: очистите строку от внешних пробельных символов.
func Trim(s string) string {
	return ""
}

// ToLower переводит строку в нижний регистр.
//
// TODO: приведите строку к нижнему регистру.
func ToLower(s string) string {
	return ""
}

// ToUpper переводит строку в верхний регистр.
//
// TODO: приведите строку к верхнему регистру.
func ToUpper(s string) string {
	return ""
}

// NormalizeEmail нормализует email.
//
// TODO: очистите email от внешних пробелов и приведите к единому регистру.
func NormalizeEmail(email string) string {
	return ""
}

// ContainsWord проверяет, содержит ли text подстроку word.
//
// TODO: проверьте наличие word внутри text.
func ContainsWord(text, word string) bool {
	return false
}

// ReplaceFirstRune заменяет первый Unicode-символ строки.
//
// TODO: верните новую строку, где первый символ заменён на r.
// Пустая строка должна обрабатываться безопасно.
func ReplaceFirstRune(s string, r rune) string {
	return ""
}

// ReverseRunes разворачивает строку по Unicode-символам.
//
// TODO: разверните строку без поломки Unicode-символов.
func ReverseRunes(s string) string {
	return ""
}

// Initials возвращает инициалы имени и фамилии.
//
// TODO: соберите инициалы из имени и фамилии.
// Пробелы по краям не должны влиять на результат.
func Initials(firstName, lastName string) string {
	return ""
}

// RepeatWord повторяет слово count раз.
//
// TODO: повторите слово нужное количество раз.
// Неположительное количество повторений должно давать пустую строку.
func RepeatWord(word string, count int) string {
	return ""
}

// JoinWithComma объединяет строки через запятую.
//
// TODO: объедините элементы с разделителем-запятой.
func JoinWithComma(values []string) string {
	return ""
}

// IsPalindrome проверяет, является ли строка палиндромом.
//
// TODO: сравните строку с её развёрнутой версией.
// Регистр и пробелы не должны мешать проверке.
func IsPalindrome(s string) bool {
	return false
}
