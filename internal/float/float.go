package float

// Блок float закрепляет работу с float64:
// дробные числа, проценты, округление, форматирование,
// преобразование float64 в int и простые вычисления.

// Add складывает два дробных числа.
//
// TODO: реализуйте сложение двух входных значений.
func Add(a, b float64) float64 {
	return 0
}

// Subtract вычитает второе дробное число из первого.
//
// TODO: реализуйте вычитание значения b из значения a.
func Subtract(a, b float64) float64 {
	return 0
}

// Multiply умножает два дробных числа.
//
// TODO: реализуйте умножение двух входных значений.
func Multiply(a, b float64) float64 {
	return 0
}

// Divide делит первое дробное число на второе.
//
// TODO: реализуйте безопасное деление дробных чисел.
// При невозможности деления функция должна вернуть 0.
func Divide(a, b float64) float64 {
	return 0
}

// DiscountPrice считает цену после скидки.
//
// TODO: примените процент скидки к цене.
// Некорректная или слишком большая скидка должна обрабатываться безопасно.
func DiscountPrice(price, percent float64) float64 {
	return 0
}

// AddTax считает цену с налогом.
//
// TODO: примените процент налога к цене.
// Некорректный налог не должен увеличивать цену.
func AddTax(price, taxPercent float64) float64 {
	return 0
}

// CelsiusToFahrenheit переводит градусы Цельсия в Фаренгейты.
//
// TODO: реализуйте перевод температуры по стандартной формуле.
func CelsiusToFahrenheit(celsius float64) float64 {
	return 0
}

// FahrenheitToCelsius переводит градусы Фаренгейта в Цельсии.
//
// TODO: реализуйте обратный перевод температуры по стандартной формуле.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return 0
}

// Average считает среднее арифметическое двух дробных чисел.
//
// TODO: посчитайте среднее значение для двух входных чисел.
func Average(a, b float64) float64 {
	return 0
}

// Round2 округляет число до 2 знаков после точки.
//
// TODO: округлите value до двух знаков после точки.
func Round2(value float64) float64 {
	return 0
}

// FormatPrice форматирует цену с двумя знаками после точки.
//
// TODO: верните строковое представление цены с двумя цифрами после точки.
func FormatPrice(price float64) string {
	return ""
}

// PercentOf считает процент от числа.
//
// TODO: посчитайте указанную процентную часть от total.
func PercentOf(total, percent float64) float64 {
	return 0
}

// GrowthPercent считает рост в процентах между старым и новым значением.
//
// TODO: посчитайте процентное изменение между oldValue и newValue.
// Нулевое старое значение должно обрабатываться безопасно.
func GrowthPercent(oldValue, newValue float64) float64 {
	return 0
}

// IsPositive проверяет, что число строго больше нуля.
//
// TODO: определите, является ли значение положительным.
func IsPositive(value float64) bool {
	return false
}

// FloatToInt преобразует float64 в int.
//
// TODO: выполните явное преобразование к int и проверьте поведение на дробных числах.
func FloatToInt(value float64) int {
	return 0
}
