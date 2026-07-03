package boolean

// Блок boolean закрепляет работу с bool:
// сравнения, логические операторы &&, ||, !,
// простые условия и функции, которые отвечают true или false.

// CanEnter проверяет, может ли человек войти.
//
// TODO: вход разрешён только совершеннолетнему человеку с билетом.
func CanEnter(age int, hasTicket bool) bool {
	return false
}

// IsAdult проверяет, является ли человек совершеннолетним.
//
// TODO: определите совершеннолетие по возрасту.
func IsAdult(age int) bool {
	return false
}

// CanBuyAlcohol проверяет, можно ли купить алкоголь.
//
// TODO: для учебной задачи используйте возрастное ограничение из тестов.
func CanBuyAlcohol(age int) bool {
	return false
}

// CanRest проверяет, можно ли отдыхать.
//
// TODO: отдых возможен в выходной или праздничный день.
func CanRest(isWeekend, isHoliday bool) bool {
	return false
}

// IsWorkingDay проверяет, является ли день рабочим.
//
// TODO: обработайте английские названия дней недели.
// Неизвестные значения считаются нерабочими.
func IsWorkingDay(day string) bool {
	return false
}

// HasAccess проверяет доступ пользователя.
//
// TODO: доступ есть у пользователя с административными правами или у владельца ресурса.
func HasAccess(isAdmin, isOwner bool) bool {
	return false
}

// CanApplyDiscount проверяет, можно ли применить скидку.
//
// TODO: скидка доступна VIP-пользователю или при достаточной сумме заказа.
func CanApplyDiscount(isVIP bool, total int) bool {
	return false
}

// ShouldNotify проверяет, нужно ли отправлять уведомление.
//
// TODO: уведомление отправляется только при выполнении обоих входных условий.
func ShouldNotify(emailVerified, notificationsEnabled bool) bool {
	return false
}

// IsValidScore проверяет корректность оценки.
//
// TODO: оценка должна входить в допустимый диапазон.
func IsValidScore(score int) bool {
	return false
}

// IsInRange проверяет, что value находится в диапазоне [min, max].
//
// TODO: проверьте попадание значения в диапазон вместе с границами.
func IsInRange(value, min, max int) bool {
	return false
}

// IsLeapYear проверяет, является ли год високосным.
//
// TODO: реализуйте стандартное правило високосного года.
// Проверьте обычные годы, века и годы, кратные 400.
func IsLeapYear(year int) bool {
	return false
}

// CanWithdraw проверяет, можно ли снять деньги.
//
// TODO: учитывайте блокировку аккаунта, сумму снятия и доступный баланс.
func CanWithdraw(balance, amount int, blocked bool) bool {
	return false
}

// LoginAllowed проверяет, разрешён ли вход.
//
// TODO: вход разрешён только при успешной проверке пароля и второго фактора.
func LoginAllowed(passwordOK, otpOK bool) bool {
	return false
}

// IsEmpty проверяет, является ли строка пустой.
//
// TODO: отличайте пустую строку от строки с пробелами или другими символами.
func IsEmpty(text string) bool {
	return false
}

// Not возвращает противоположное bool-значение.
//
// TODO: реализуйте логическое отрицание.
func Not(flag bool) bool {
	return false
}
