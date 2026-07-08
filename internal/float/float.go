package float

import (
	"fmt"
	"math"
)

// Add складывает два дробных числа.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract вычитает второе дробное число из первого.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply умножает два дробных числа.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide делит первое дробное число на второе.
// При невозможности деления функция должна вернуть 0.
func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// DiscountPrice считает цену после скидки.
// Некорректная или слишком большая скидка должна обрабатываться безопасно.
func DiscountPrice(price, percent float64) float64 {
	if percent < 0 {
		return price
	}
	if percent > 100 {
		return 0
	}
	return price - price*percent/100
}

// AddTax считает цену с налогом.
// Некорректный налог не должен увеличивать цену.
func AddTax(price, taxPercent float64) float64 {
	if taxPercent < 0 {
		return price
	}
	return price + price*taxPercent/100
}

// CelsiusToFahrenheit переводит градусы Цельсия в Фаренгейты.
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// FahrenheitToCelsius переводит градусы Фаренгейта в Цельсии.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Average считает среднее арифметическое двух дробных чисел.
func Average(a, b float64) float64 {
	return (a + b) / 2
}

// Round2 округляет число до 2 знаков после точки.
func Round2(value float64) float64 {
	return math.Round(value*100) / 100
}

// FormatPrice форматирует цену с двумя знаками после точки.
func FormatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

// PercentOf считает процент от числа.
func PercentOf(total, percent float64) float64 {
	return total * percent / 100
}

// GrowthPercent считает рост в процентах между старым и новым значением.
// Нулевое старое значение должно обрабатываться безопасно.
func GrowthPercent(oldValue, newValue float64) float64 {
	if oldValue == 0 {
		return 0
	}
	return (newValue - oldValue) / oldValue * 100
}

// IsPositive проверяет, что число строго больше нуля.
func IsPositive(value float64) bool {
	return value > 0
}

// FloatToInt преобразует float64 в int.
func FloatToInt(value float64) int {
	return int(value)
}
