package profile

import (
	"fmt"
	"strings"
)

// Блок profile — объединяющая задача.
// Здесь нужно закрепить:
// string normalization, int, bool, if,
// форматирование строки и простую бизнес-логику.

// BuildUserCard собирает карточку пользователя.
// Пустое после очистки имя должно заменяться значением по умолчанию.
// Точный формат итоговой строки смотрите в тестах.
func BuildUserCard(name string, age int, active bool) string {
	trimSpaceName := strings.TrimSpace(name)

	titleName := strings.Title(strings.ToLower(trimSpaceName))
	if titleName == "" {
		titleName = "Unknown"
	}
	resultAdult := "adult"
	if age < 18 {
		resultAdult = "minor"
	}
	resultStatus := "active"
	if !active {
		resultStatus = "inactive"
	}
	return fmt.Sprintf("name=%s age=%d group=%s status=%s", titleName, age, resultAdult, resultStatus)
}
