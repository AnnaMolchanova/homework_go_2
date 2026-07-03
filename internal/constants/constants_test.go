package constants

import "testing"

func TestAppName(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "app name", got: AppName(), want: "go-homework-2"},
		{name: "not empty", got: AppName(), want: "go-homework-2"},
		{name: "stable value", got: AppName(), want: appName},
		{name: "same call result", got: AppName(), want: AppName()},
		{name: "expected literal", got: appName, want: "go-homework-2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}

func TestMaxAttempts(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "max attempts", got: MaxAttempts(), want: 3},
		{name: "stable value", got: MaxAttempts(), want: maxAttempts},
		{name: "literal check", got: maxAttempts, want: 3},
		{name: "same call result", got: MaxAttempts(), want: MaxAttempts()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %#v, want %#v", tt.got, tt.want)
			}
		})
	}
	if MaxAttempts() <= 0 {
		t.Fatalf("got %#v, want %#v", MaxAttempts() > 0, true)
	}
}

func TestStatusText(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   string
	}{
		{name: "new", status: StatusNew, want: "new"},
		{name: "paid", status: StatusPaid, want: "paid"},
		{name: "canceled", status: StatusCanceled, want: "canceled"},
		{name: "unknown negative", status: -1, want: "unknown"},
		{name: "unknown large", status: 100, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StatusText(tt.status)
			if got != tt.want {
				t.Fatalf("StatusText(%d)=%q, want %q", tt.status, got, tt.want)
			}
		})
	}
}

func TestIsFinalStatus(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{name: "new is not final", status: StatusNew, want: false},
		{name: "paid is final", status: StatusPaid, want: true},
		{name: "canceled is final", status: StatusCanceled, want: true},
		{name: "unknown negative", status: -1, want: false},
		{name: "unknown large", status: 100, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFinalStatus(tt.status)
			if got != tt.want {
				t.Fatalf("IsFinalStatus(%d)=%t, want %t", tt.status, got, tt.want)
			}
		})
	}
}

func TestNextStatus(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   int
	}{
		{name: "new to paid", status: StatusNew, want: StatusPaid},
		{name: "paid stays paid", status: StatusPaid, want: StatusPaid},
		{name: "canceled stays canceled", status: StatusCanceled, want: StatusCanceled},
		{name: "unknown negative to new", status: -1, want: StatusNew},
		{name: "unknown large to new", status: 100, want: StatusNew},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextStatus(tt.status)
			if got != tt.want {
				t.Fatalf("NextStatus(%d)=%d, want %d", tt.status, got, tt.want)
			}
		})
	}
}

func TestRoleText(t *testing.T) {
	tests := []struct {
		name string
		role int
		want string
	}{
		{name: "guest", role: RoleGuest, want: "guest"},
		{name: "user", role: RoleUser, want: "user"},
		{name: "admin", role: RoleAdmin, want: "admin"},
		{name: "unknown negative", role: -1, want: "unknown"},
		{name: "unknown large", role: 100, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RoleText(tt.role)
			if got != tt.want {
				t.Fatalf("RoleText(%d)=%q, want %q", tt.role, got, tt.want)
			}
		})
	}
}

func TestCanEdit(t *testing.T) {
	tests := []struct {
		name string
		role int
		want bool
	}{
		{name: "guest can not edit", role: RoleGuest, want: false},
		{name: "user can not edit", role: RoleUser, want: false},
		{name: "admin can edit", role: RoleAdmin, want: true},
		{name: "unknown negative", role: -1, want: false},
		{name: "unknown large", role: 100, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanEdit(tt.role)
			if got != tt.want {
				t.Fatalf("CanEdit(%d)=%t, want %t", tt.role, got, tt.want)
			}
		})
	}
}

func TestHTTPStatusText(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{name: "ok", code: 200, want: "OK"},
		{name: "created", code: 201, want: "Created"},
		{name: "bad request", code: 400, want: "Bad Request"},
		{name: "not found", code: 404, want: "Not Found"},
		{name: "unknown", code: 418, want: "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HTTPStatusText(tt.code)
			if got != tt.want {
				t.Fatalf("HTTPStatusText(%d)=%q, want %q", tt.code, got, tt.want)
			}
		})
	}
}

func TestDayType(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want string
	}{
		{name: "monday", day: 1, want: "working"},
		{name: "friday", day: 5, want: "working"},
		{name: "saturday", day: 6, want: "weekend"},
		{name: "sunday", day: 7, want: "weekend"},
		{name: "unknown", day: 0, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DayType(tt.day)
			if got != tt.want {
				t.Fatalf("DayType(%d)=%q, want %q", tt.day, got, tt.want)
			}
		})
	}
}

func TestPriorityText(t *testing.T) {
	tests := []struct {
		name     string
		priority int
		want     string
	}{
		{name: "low", priority: PriorityLow, want: "low"},
		{name: "medium", priority: PriorityMedium, want: "medium"},
		{name: "high", priority: PriorityHigh, want: "high"},
		{name: "unknown zero", priority: 0, want: "unknown"},
		{name: "unknown large", priority: 100, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PriorityText(tt.priority)
			if got != tt.want {
				t.Fatalf("PriorityText(%d)=%q, want %q", tt.priority, got, tt.want)
			}
		})
	}
}

func TestIsKnownStatus(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{name: "new", status: StatusNew, want: true},
		{name: "paid", status: StatusPaid, want: true},
		{name: "canceled", status: StatusCanceled, want: true},
		{name: "unknown negative", status: -1, want: false},
		{name: "unknown large", status: 100, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsKnownStatus(tt.status)
			if got != tt.want {
				t.Fatalf("IsKnownStatus(%d)=%t, want %t", tt.status, got, tt.want)
			}
		})
	}
}

func TestPaymentStateText(t *testing.T) {
	tests := []struct {
		name     string
		paid     bool
		canceled bool
		want     string
	}{
		{name: "pending", paid: false, canceled: false, want: "pending"},
		{name: "paid", paid: true, canceled: false, want: "paid"},
		{name: "canceled", paid: false, canceled: true, want: "canceled"},
		{name: "canceled has priority", paid: true, canceled: true, want: "canceled"},
		{name: "default pending", paid: false, canceled: false, want: "pending"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PaymentStateText(tt.paid, tt.canceled)
			if got != tt.want {
				t.Fatalf("PaymentStateText(%t, %t)=%q, want %q", tt.paid, tt.canceled, got, tt.want)
			}
		})
	}
}

func TestTrafficLightAction(t *testing.T) {
	tests := []struct {
		name  string
		color string
		want  string
	}{
		{name: "red", color: "red", want: "stop"},
		{name: "yellow", color: "yellow", want: "wait"},
		{name: "green", color: "green", want: "go"},
		{name: "unknown", color: "blue", want: "unknown"},
		{name: "empty", color: "", want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrafficLightAction(tt.color)
			if got != tt.want {
				t.Fatalf("TrafficLightAction(%q)=%q, want %q", tt.color, got, tt.want)
			}
		})
	}
}

func TestGradeText(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{name: "excellent", score: 95, want: "excellent"},
		{name: "good", score: 80, want: "good"},
		{name: "passed", score: 60, want: "passed"},
		{name: "retry", score: 30, want: "retry"},
		{name: "invalid greater than max", score: 101, want: "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GradeText(tt.score)
			if got != tt.want {
				t.Fatalf("GradeText(%d)=%q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestEventTypeText(t *testing.T) {
	tests := []struct {
		name      string
		eventType int
		want      string
	}{
		{name: "created", eventType: EventCreated, want: "created"},
		{name: "updated", eventType: EventUpdated, want: "updated"},
		{name: "deleted", eventType: EventDeleted, want: "deleted"},
		{name: "unknown negative", eventType: -1, want: "unknown"},
		{name: "unknown large", eventType: 100, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EventTypeText(tt.eventType)
			if got != tt.want {
				t.Fatalf("EventTypeText(%d)=%q, want %q", tt.eventType, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "app=go-homework-2 attempts=3 status=paid role=admin final=true"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
