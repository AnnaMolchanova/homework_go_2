package constants

// Блок constants закрепляет:
// const, iota, enum-like значения, switch,
// статусы, роли, приоритеты, типы событий и простые текстовые маппинги.

const appName = "go-homework-2"

const maxAttempts = 3

// Статусы оплаты.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// Роли пользователя.
const (
	RoleGuest = iota
	RoleUser
	RoleAdmin
)

// Приоритеты задачи.
const (
	PriorityLow = iota + 1
	PriorityMedium
	PriorityHigh
)

// Типы событий.
const (
	EventCreated = iota
	EventUpdated
	EventDeleted
)

// AppName возвращает имя приложения.
func AppName() string {
	return appName
}

// MaxAttempts возвращает максимальное количество попыток.
func MaxAttempts() int {
	return maxAttempts
}

// StatusText возвращает текстовое представление статуса.
func StatusText(status int) string {
	switch status {
	case StatusNew:
		return "new"
	case StatusPaid:
		return "paid"
	case StatusCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}

// IsFinalStatus проверяет, является ли статус финальным.
func IsFinalStatus(status int) bool {
	return status == StatusPaid || status == StatusCanceled
}

// NextStatus возвращает следующий статус.
func NextStatus(status int) int {
	switch status {
	case StatusNew:
		return StatusPaid
	case StatusPaid:
		return StatusPaid
	case StatusCanceled:
		return StatusCanceled
	default:
		return StatusNew
	}
}

// RoleText возвращает текстовое представление роли.
func RoleText(role int) string {
	switch role {
	case RoleGuest:
		return "guest"
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}

// CanEdit проверяет, может ли пользователь редактировать данные.
func CanEdit(role int) bool {
	return role == RoleAdmin
}

// HTTPStatusText возвращает текст HTTP-статуса.
func HTTPStatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	default:
		return "Unknown"
	}
}

// DayType возвращает тип дня недели.
func DayType(day int) string {
	switch day {
	case 1, 2, 3, 4, 5:
		return "working"
	case 6, 7:
		return "weekend"
	default:
		return "unknown"
	}
}

// PriorityText возвращает текстовое представление приоритета.
func PriorityText(priority int) string {
	switch priority {
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	default:
		return "unknown"
	}
}

// IsKnownStatus проверяет, известен ли статус.
func IsKnownStatus(status int) bool {
	return status == StatusNew || status == StatusPaid || status == StatusCanceled
}

// PaymentStateText возвращает текст состояния оплаты.
func PaymentStateText(paid, canceled bool) string {
	if canceled {
		return "canceled"
	}

	if paid {
		return "paid"
	}

	return "pending"
}

// TrafficLightAction возвращает действие по цвету светофора.
func TrafficLightAction(color string) string {
	switch color {
	case "red":
		return "stop"
	case "yellow":
		return "wait"
	case "green":
		return "go"
	default:
		return "unknown"
	}
}

// GradeText возвращает текстовую оценку по score.
func GradeText(score int) string {
	if score < 0 || score > 100 {
		return "invalid"
	}

	if score >= 90 {
		return "excellent"
	}

	if score >= 75 {
		return "good"
	}

	if score >= 60 {
		return "passed"
	}

	return "retry"
}

// EventTypeText возвращает текстовое представление типа события.
func EventTypeText(eventType int) string {
	switch eventType {
	case EventCreated:
		return "created"
	case EventUpdated:
		return "updated"
	case EventDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}
