package boolean

// CanEnter проверяет, может ли человек войти.
func CanEnter(age int, hasTicket bool) bool {
	return age >= 18 && hasTicket
}

// IsAdult проверяет, является ли человек совершеннолетним.
func IsAdult(age int) bool {
	return age >= 18
}

// CanBuyAlcohol проверяет, можно ли купить алкоголь.
func CanBuyAlcohol(age int) bool {
	return age >= 18
}

// CanRest проверяет, можно ли отдыхать.
func CanRest(isWeekend, isHoliday bool) bool {
	return isWeekend || isHoliday
}

// IsWorkingDay проверяет, является ли день рабочим.
// Неизвестные значения считаются нерабочими.
func IsWorkingDay(day string) bool {
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		return true
	default:
		return false
	}
}

// HasAccess проверяет доступ пользователя.
func HasAccess(isAdmin, isOwner bool) bool {
	return isAdmin || isOwner
}

// CanApplyDiscount проверяет, можно ли применить скидку.
func CanApplyDiscount(isVIP bool, total int) bool {
	return isVIP || total >= 5000
}

// ShouldNotify проверяет, нужно ли отправлять уведомление.
func ShouldNotify(emailVerified, notificationsEnabled bool) bool {
	return emailVerified && notificationsEnabled
}

// IsValidScore проверяет корректность оценки.
func IsValidScore(score int) bool {
	return score >= 0 && score <= 100
}

// IsInRange проверяет, что value находится в диапазоне [min, max].
func IsInRange(value, min, max int) bool {
	return value >= min && value <= max
}

// IsLeapYear проверяет, является ли год високосным.
// Проверьте обычные годы, века и годы, кратные 400.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

// CanWithdraw проверяет, можно ли снять деньги.
func CanWithdraw(balance, amount int, blocked bool) bool {
	return !blocked && amount > 0 && balance >= amount
}

// LoginAllowed проверяет, разрешён ли вход.
func LoginAllowed(passwordOK, otpOK bool) bool {
	return passwordOK && otpOK

}

// IsEmpty проверяет, является ли строка пустой.
func IsEmpty(text string) bool {
	return text == ""
}

// Not возвращает противоположное bool-значение.
func Not(flag bool) bool {
	return !flag
}
