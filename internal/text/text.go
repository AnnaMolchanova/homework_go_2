package text

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// ByteLen возвращает длину строки в байтах.
func ByteLen(s string) int {
	return len(s)
}

// RuneLen возвращает количество Unicode-символов в строке.
func RuneLen(s string) int {
	return utf8.RuneCountInString(s)
}

// FirstRune возвращает первый Unicode-символ строки.
func FirstRune(s string) string {
	for _, r := range s {
		return string(r)
	}
	return ""
}

// LastRune возвращает последний Unicode-символ строки.
func LastRune(s string) string {
	var last rune
	for _, r := range s {
		last = r
	}
	if last == 0 {
		return ""
	}
	return string(last)
}

// Trim убирает пробелы по краям строки.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ToLower переводит строку в нижний регистр.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper переводит строку в верхний регистр.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// NormalizeEmail нормализует email.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// ContainsWord проверяет, содержит ли text подстроку word.
func ContainsWord(text, word string) bool {
	return strings.Contains(text, word)
}

// ReplaceFirstRune заменяет первый Unicode-символ строки.
// Пустая строка должна обрабатываться безопасно.
func ReplaceFirstRune(s string, r rune) string {
	if s == "" {
		return ""
	}
	_, size := utf8.DecodeRuneInString(s)
	return string(r) + s[size:]
}

// ReverseRunes разворачивает строку по Unicode-символам.
func ReverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Initials возвращает инициалы имени и фамилии.
// Пробелы по краям не должны влиять на результат.
func Initials(firstName, lastName string) string {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	result := ""

	if firstName != "" {
		result += strings.ToUpper(FirstRune(firstName))
	}

	if lastName != "" {
		result += strings.ToUpper(FirstRune(lastName))
	}

	return result
}

// RepeatWord повторяет слово count раз.
// Неположительное количество повторений должно давать пустую строку.
func RepeatWord(word string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(word, count)
}

// JoinWithComma объединяет строки через запятую.
func JoinWithComma(values []string) string {
	return strings.Join(values, ",")
}

// IsPalindrome проверяет, является ли строка палиндромом.
// Регистр и пробелы не должны мешать проверке.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	cleaned := ""
	for _, r := range s {
		if !unicode.IsSpace(r) {
			cleaned += string(r)
		}
	}
	return cleaned == ReverseRunes(cleaned)
}
