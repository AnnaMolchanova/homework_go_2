package pointers

// Блок pointers закрепляет работу с указателями:
// оператор &, оператор *, nil, изменение значения через указатель,
// возврат указателя и безопасные проверки перед разыменованием.

// ValueOrDefault возвращает значение по указателю или значение по умолчанию.
//
// TODO: безопасно прочитайте значение по указателю или используйте fallback.
func ValueOrDefault(p *int, def int) int {
	return 0
}

// Increment увеличивает значение по указателю на 1.
//
// TODO: измените исходное значение через указатель и верните новый результат.
func Increment(p *int) int {
	return 0
}

// SetValue записывает новое значение по указателю.
//
// TODO: запишите value по адресу, если это возможно, и верните успешность операции.
func SetValue(p *int, value int) bool {
	return false
}

// Swap меняет местами два значения по указателям.
//
// TODO: поменяйте два значения местами через указатели.
func Swap(a, b *int) bool {
	return false
}

// ResetToZero сбрасывает значение по указателю в 0.
//
// TODO: сбросьте исходное значение через указатель и верните успешность операции.
func ResetToZero(p *int) bool {
	return false
}

// AddToValue прибавляет delta к значению по указателю.
//
// TODO: измените исходное значение на delta и верните новое значение.
func AddToValue(p *int, delta int) int {
	return 0
}

// MaxPointer возвращает указатель на большее значение.
//
// TODO: выберите указатель на большее значение с безопасной обработкой nil.
func MaxPointer(a, b *int) *int {
	return nil
}

// IsNil проверяет, равен ли указатель nil.
//
// TODO: реализуйте проверку nil-указателя.
func IsNil(p *int) bool {
	return false
}

// CopyValue возвращает копию значения по указателю.
//
// TODO: безопасно прочитайте значение по указателю.
func CopyValue(p *int) int {
	return 0
}

// DoubleInPlace умножает значение по указателю на 2.
//
// TODO: измените исходное значение через указатель.
func DoubleInPlace(p *int) bool {
	return false
}

// NewInt создаёт новое int-значение и возвращает указатель на него.
//
// TODO: создайте новое значение и верните указатель на него.
func NewInt(value int) *int {
	return nil
}

// DivideInto делит a на b и записывает результат в out.
//
// TODO: безопасно выполните целочисленное деление и запишите результат в out.
func DivideInto(out *int, a, b int) bool {
	return false
}

// ApplyDiscountInPlace применяет скидку к цене по указателю.
//
// TODO: измените цену через указатель с учётом процента скидки.
// Некорректные значения должны обрабатываться безопасно.
func ApplyDiscountInPlace(price *int, percent int) bool {
	return false
}

// ChoosePointer выбирает первый доступный указатель.
//
// TODO: выберите primary, если он доступен, иначе используйте fallback.
func ChoosePointer(primary, fallback *int) *int {
	return nil
}

// PointToLarger возвращает указатель на большее значение.
//
// TODO: выберите указатель на большее значение с безопасной обработкой nil.
func PointToLarger(a, b *int) *int {
	return nil
}
