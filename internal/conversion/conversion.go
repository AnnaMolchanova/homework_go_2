package conversion

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
//
// TODO: выполните явное преобразование типа.
func IntToInt64(n int) int64 {
	return 0
}

// Int64ToInt явно преобразует int64 в int.
//
// TODO: выполните явное преобразование типа.
// Для этой учебной задачи переполнение не обрабатываем.
func Int64ToInt(n int64) int {
	return 0
}

// RubToKop переводит рубли в копейки.
//
// TODO: переведите значение из рублей в копейки и сохраните тип результата.
func RubToKop(rub Rub) Kop {
	return 0
}

// KopToRub переводит копейки в рубли через целочисленное деление.
//
// TODO: переведите значение из копеек в рубли и сохраните тип результата.
func KopToRub(kop Kop) Rub {
	return 0
}

// UserIDToString преобразует UserID в string.
//
// TODO: преобразуйте пользовательский идентификатор в строку.
func UserIDToString(id UserID) string {
	return ""
}

// ParseInt преобразует строку в int.
//
// TODO: распарсите строку как int и корректно верните ошибку.
func ParseInt(text string) (int, error) {
	return 0, nil
}

// ParseAndDouble парсит строку в int и умножает результат на 2.
//
// TODO: распарсите строку, обработайте ошибку и верните удвоенное число.
func ParseAndDouble(text string) (int, error) {
	return 0, nil
}

// IntToString преобразует int в string.
//
// TODO: преобразуйте целое число в строку.
func IntToString(n int) string {
	return ""
}

// FloatToString форматирует float64 с двумя знаками после точки.
//
// TODO: преобразуйте дробное число в строку с фиксированным количеством знаков после точки.
func FloatToString(value float64) string {
	return ""
}

// ParseBoolText преобразует строку в bool.
//
// TODO: распарсите строковое bool-значение и корректно верните ошибку.
func ParseBoolText(text string) (bool, error) {
	return false, nil
}

// BoolToText преобразует bool в string.
//
// TODO: преобразуйте bool в его строковое представление.
func BoolToText(value bool) string {
	return ""
}

// SumIntAndInt64 складывает int и int64.
//
// TODO: сложите значения разных целочисленных типов через явное преобразование.
func SumIntAndInt64(a int, b int64) int64 {
	return 0
}

// PriceRubStringToKop преобразует строку с рублями в копейки.
//
// TODO: распарсите цену в рублях, проверьте корректность и верните цену в копейках.
func PriceRubStringToKop(text string) (Kop, error) {
	return 0, nil
}

// SafeParsePositive безопасно парсит положительное число.
//
// TODO: верните положительное число из строки или 0 при некорректном значении.
func SafeParsePositive(text string) int {
	return 0
}

// FormatUser форматирует пользователя в строку.
//
// TODO: соберите строковое представление пользователя по формату из тестов.
func FormatUser(id UserID, name string) string {
	return ""
}
