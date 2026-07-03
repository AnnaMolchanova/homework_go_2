package order

// Блок order — объединяющая задача.
// Здесь нужно закрепить:
// constants/iota, switch, int conversion, bool,
// вычисление цены в копейках и сборку итоговой строки.

// Статусы заказа.
//
// TODO: связанные статусы должны иметь последовательные значения.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// OrderSummary собирает краткое описание заказа.
//
// TODO: преобразуйте статус в текст, bool-флаг оплаты — в текстовое состояние,
// цену в рублях — в копейки, затем соберите итоговую строку.
// Неизвестный статус и некорректная цена должны обрабатываться безопасно.
func OrderSummary(status int, priceRub int, paid bool) string {
	return ""
}
