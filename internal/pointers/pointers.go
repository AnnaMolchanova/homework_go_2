package pointers

// Блок pointers закрепляет работу с указателями:
// оператор &, оператор *, nil, изменение значения через указатель,
// возврат указателя и безопасные проверки перед разыменованием.

// ValueOrDefault возвращает значение по указателю или значение по умолчанию.
func ValueOrDefault(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}

// Increment увеличивает значение по указателю на 1.
func Increment(p *int) int {
	if p == nil {
		return 0
	}
	*p = *p + 1
	return *p
}

// SetValue записывает новое значение по указателю.
func SetValue(p *int, value int) bool {
	if p == nil {
		return false
	}
	*p = value
	return true
}

// Swap меняет местами два значения по указателям.
func Swap(a, b *int) bool {
	if a == nil || b == nil {
		return false
	}
	temp := *a
	*a = *b
	*b = temp
	return true
}

// ResetToZero сбрасывает значение по указателю в 0.
func ResetToZero(p *int) bool {
	if p == nil {
		return false
	}
	*p = 0
	return true
}

// AddToValue прибавляет delta к значению по указателю.
func AddToValue(p *int, delta int) int {
	if p == nil {
		return 0
	}
	*p = *p + delta
	return *p
}

// MaxPointer возвращает указатель на большее значение.
func MaxPointer(a, b *int) *int {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if *a >= *b {
		return a
	}
	return b
}

// IsNil проверяет, равен ли указатель nil.
func IsNil(p *int) bool {
	return p == nil
}

// CopyValue возвращает копию значения по указателю.
func CopyValue(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// DoubleInPlace умножает значение по указателю на 2.
func DoubleInPlace(p *int) bool {
	if p == nil {
		return false
	}
	*p = *p * 2
	return true
}

// NewInt создаёт новое int-значение и возвращает указатель на него.
func NewInt(value int) *int {
	return &value
}

// DivideInto делит a на b и записывает результат в out.
func DivideInto(out *int, a, b int) bool {
	if out == nil || b == 0 {
		return false
	}
	*out = a / b
	return true
}

// ApplyDiscountInPlace применяет скидку к цене по указателю.
// Некорректные значения должны обрабатываться безопасно.
func ApplyDiscountInPlace(price *int, percent int) bool {
	if price == nil {
		return false
	}
	if percent < 0 {
		return false
	}
	if percent >= 100 {
		*price = 0
		return true
	}
	*price = *price - (*price * percent / 100)
	return true
}

// ChoosePointer выбирает первый доступный указатель.
func ChoosePointer(primary, fallback *int) *int {
	if primary != nil {
		return primary
	}
	return fallback
}

// PointToLarger возвращает указатель на большее значение.
func PointToLarger(a, b *int) *int {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if *a >= *b {
		return a
	}
	return b
}
