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
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Unknown"
	} else {
		name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
	}

	group := "minor"
	if age >= 18 {
		group = "adult"
	}

	status := "inactive"
	if active {
		status = "active"
	}
	return fmt.Sprintf("name=%s age=%d group=%s status=%s", name, age, group, status)
}
