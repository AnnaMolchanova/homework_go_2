package order

import (
	"fmt"
)

// Блок order — объединяющая задача.
// Здесь нужно закрепить:
// constants/iota, switch, int conversion, bool,
// вычисление цены в копейках и сборку итоговой строки.

// Статусы заказа.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// OrderSummary собирает краткое описание заказа.
// цену в рублях — в копейки, затем соберите итоговую строку.
// Неизвестный статус и некорректная цена должны обрабатываться безопасно.
func OrderSummary(status int, priceRub int, paid bool) string {
	statusText := "unknown"
	switch status {
	case StatusNew:
		statusText = "new"
	case StatusPaid:
		statusText = "paid"
	case StatusCanceled:
		statusText = "canceled"
	}
	paymentText := "not_paid"
	if paid {
		paymentText = "paid"
	}
	if priceRub < 0 {
		priceRub = 0
	}
	priceKop := priceRub * 100
	return fmt.Sprintf("status=%s payment=%s price_kop=%d", statusText, paymentText, priceKop)
}
