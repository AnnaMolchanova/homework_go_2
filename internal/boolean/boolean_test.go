package boolean

import "testing"

func TestCanEnter(t *testing.T) {
	tests := []struct {
		name      string
		age       int
		hasTicket bool
		want      bool
	}{
		{name: "adult with ticket", age: 20, hasTicket: true, want: true},
		{name: "minor with ticket", age: 17, hasTicket: true, want: false},
		{name: "adult without ticket", age: 20, hasTicket: false, want: false},
		{name: "exact adult age with ticket", age: 18, hasTicket: true, want: true},
		{name: "minor without ticket", age: 10, hasTicket: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanEnter(tt.age, tt.hasTicket)
			if got != tt.want {
				t.Fatalf("CanEnter(%d, %t)=%t, want %t", tt.age, tt.hasTicket, got, tt.want)
			}
		})
	}
}

func TestIsAdult(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{name: "less than adult", age: 17, want: false},
		{name: "exact adult age", age: 18, want: true},
		{name: "greater than adult", age: 25, want: true},
		{name: "zero age", age: 0, want: false},
		{name: "negative age", age: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAdult(tt.age)
			if got != tt.want {
				t.Fatalf("IsAdult(%d)=%t, want %t", tt.age, got, tt.want)
			}
		})
	}
}

func TestCanBuyAlcohol(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{name: "minor", age: 16, want: false},
		{name: "seventeen", age: 17, want: false},
		{name: "eighteen", age: 18, want: true},
		{name: "adult", age: 30, want: true},
		{name: "zero age", age: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanBuyAlcohol(tt.age)
			if got != tt.want {
				t.Fatalf("CanBuyAlcohol(%d)=%t, want %t", tt.age, got, tt.want)
			}
		})
	}
}

func TestCanRest(t *testing.T) {
	tests := []struct {
		name      string
		isWeekend bool
		isHoliday bool
		want      bool
	}{
		{name: "weekend only", isWeekend: true, isHoliday: false, want: true},
		{name: "holiday only", isWeekend: false, isHoliday: true, want: true},
		{name: "weekend and holiday", isWeekend: true, isHoliday: true, want: true},
		{name: "working day", isWeekend: false, isHoliday: false, want: false},
		{name: "no rest flags", isWeekend: false, isHoliday: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanRest(tt.isWeekend, tt.isHoliday)
			if got != tt.want {
				t.Fatalf("CanRest(%t, %t)=%t, want %t", tt.isWeekend, tt.isHoliday, got, tt.want)
			}
		})
	}
}

func TestIsWorkingDay(t *testing.T) {
	tests := []struct {
		name string
		day  string
		want bool
	}{
		{name: "monday", day: "monday", want: true},
		{name: "friday", day: "friday", want: true},
		{name: "saturday", day: "saturday", want: false},
		{name: "sunday", day: "sunday", want: false},
		{name: "unknown day", day: "holiday", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsWorkingDay(tt.day)
			if got != tt.want {
				t.Fatalf("IsWorkingDay(%q)=%t, want %t", tt.day, got, tt.want)
			}
		})
	}
}

func TestHasAccess(t *testing.T) {
	tests := []struct {
		name    string
		isAdmin bool
		isOwner bool
		want    bool
	}{
		{name: "admin only", isAdmin: true, isOwner: false, want: true},
		{name: "owner only", isAdmin: false, isOwner: true, want: true},
		{name: "admin and owner", isAdmin: true, isOwner: true, want: true},
		{name: "no access", isAdmin: false, isOwner: false, want: false},
		{name: "default user", isAdmin: false, isOwner: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasAccess(tt.isAdmin, tt.isOwner)
			if got != tt.want {
				t.Fatalf("HasAccess(%t, %t)=%t, want %t", tt.isAdmin, tt.isOwner, got, tt.want)
			}
		})
	}
}

func TestCanApplyDiscount(t *testing.T) {
	tests := []struct {
		name  string
		isVIP bool
		total int
		want  bool
	}{
		{name: "vip with small total", isVIP: true, total: 100, want: true},
		{name: "not vip with large total", isVIP: false, total: 5000, want: true},
		{name: "not vip with greater total", isVIP: false, total: 7000, want: true},
		{name: "not vip with small total", isVIP: false, total: 4999, want: false},
		{name: "vip with zero total", isVIP: true, total: 0, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanApplyDiscount(tt.isVIP, tt.total)
			if got != tt.want {
				t.Fatalf("CanApplyDiscount(%t, %d)=%t, want %t", tt.isVIP, tt.total, got, tt.want)
			}
		})
	}
}

func TestShouldNotify(t *testing.T) {
	tests := []struct {
		name                 string
		emailVerified        bool
		notificationsEnabled bool
		want                 bool
	}{
		{name: "verified and enabled", emailVerified: true, notificationsEnabled: true, want: true},
		{name: "not verified", emailVerified: false, notificationsEnabled: true, want: false},
		{name: "disabled notifications", emailVerified: true, notificationsEnabled: false, want: false},
		{name: "both false", emailVerified: false, notificationsEnabled: false, want: false},
		{name: "verified only is not enough", emailVerified: true, notificationsEnabled: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldNotify(tt.emailVerified, tt.notificationsEnabled)
			if got != tt.want {
				t.Fatalf(
					"ShouldNotify(%t, %t)=%t, want %t",
					tt.emailVerified,
					tt.notificationsEnabled,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsValidScore(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  bool
	}{
		{name: "zero", score: 0, want: true},
		{name: "middle", score: 55, want: true},
		{name: "max", score: 100, want: true},
		{name: "less than zero", score: -1, want: false},
		{name: "greater than max", score: 101, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidScore(tt.score)
			if got != tt.want {
				t.Fatalf("IsValidScore(%d)=%t, want %t", tt.score, got, tt.want)
			}
		})
	}
}

func TestIsInRange(t *testing.T) {
	tests := []struct {
		name  string
		value int
		min   int
		max   int
		want  bool
	}{
		{name: "inside range", value: 5, min: 1, max: 10, want: true},
		{name: "equal min", value: 1, min: 1, max: 10, want: true},
		{name: "equal max", value: 10, min: 1, max: 10, want: true},
		{name: "less than min", value: 0, min: 1, max: 10, want: false},
		{name: "greater than max", value: 11, min: 1, max: 10, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsInRange(tt.value, tt.min, tt.max)
			if got != tt.want {
				t.Fatalf("IsInRange(%d, %d, %d)=%t, want %t", tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{name: "regular leap year", year: 2024, want: true},
		{name: "regular not leap year", year: 2023, want: false},
		{name: "century not leap year", year: 1900, want: false},
		{name: "century leap year", year: 2000, want: true},
		{name: "future leap year", year: 2400, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.want {
				t.Fatalf("IsLeapYear(%d)=%t, want %t", tt.year, got, tt.want)
			}
		})
	}
}

func TestCanWithdraw(t *testing.T) {
	tests := []struct {
		name    string
		balance int
		amount  int
		blocked bool
		want    bool
	}{
		{name: "enough money", balance: 1000, amount: 700, blocked: false, want: true},
		{name: "not enough money", balance: 500, amount: 700, blocked: false, want: false},
		{name: "blocked account", balance: 1000, amount: 700, blocked: true, want: false},
		{name: "amount equals balance", balance: 1000, amount: 1000, blocked: false, want: true},
		{name: "invalid amount", balance: 1000, amount: 0, blocked: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanWithdraw(tt.balance, tt.amount, tt.blocked)
			if got != tt.want {
				t.Fatalf("CanWithdraw(%d, %d, %t)=%t, want %t", tt.balance, tt.amount, tt.blocked, got, tt.want)
			}
		})
	}
}

func TestLoginAllowed(t *testing.T) {
	tests := []struct {
		name       string
		passwordOK bool
		otpOK      bool
		want       bool
	}{
		{name: "password and otp ok", passwordOK: true, otpOK: true, want: true},
		{name: "password failed", passwordOK: false, otpOK: true, want: false},
		{name: "otp failed", passwordOK: true, otpOK: false, want: false},
		{name: "both failed", passwordOK: false, otpOK: false, want: false},
		{name: "only password is not enough", passwordOK: true, otpOK: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoginAllowed(tt.passwordOK, tt.otpOK)
			if got != tt.want {
				t.Fatalf("LoginAllowed(%t, %t)=%t, want %t", tt.passwordOK, tt.otpOK, got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{name: "empty string", text: "", want: true},
		{name: "regular text", text: "go", want: false},
		{name: "space is not empty", text: " ", want: false},
		{name: "newline is not empty", text: "\n", want: false},
		{name: "cyrillic text", text: "Привет", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmpty(tt.text)
			if got != tt.want {
				t.Fatalf("IsEmpty(%q)=%t, want %t", tt.text, got, tt.want)
			}
		})
	}
}

func TestNot(t *testing.T) {
	tests := []struct {
		name string
		flag bool
		want bool
	}{
		{name: "false to true", flag: false, want: true},
		{name: "true to false", flag: true, want: false},
		{name: "literal expression true", flag: 10 > 5, want: false},
		{name: "literal expression false", flag: 10 < 5, want: true},
		{name: "zero equals one", flag: 0 == 1, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Not(tt.flag)
			if got != tt.want {
				t.Fatalf("Not(%t)=%t, want %t", tt.flag, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "enter=true rest=true access=true leap=true withdraw=true"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
