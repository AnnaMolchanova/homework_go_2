package conversion

import "testing"

func TestIntToInt64(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{name: "positive", n: 42, want: int64(42)},
		{name: "zero", n: 0, want: int64(0)},
		{name: "negative", n: -42, want: int64(-42)},
		{name: "large positive", n: 1_000_000, want: int64(1_000_000)},
		{name: "large negative", n: -1_000_000, want: int64(-1_000_000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToInt64(tt.n)
			if got != tt.want {
				t.Fatalf("IntToInt64(%d)=%d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestInt64ToInt(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int
	}{
		{name: "positive", n: int64(42), want: 42},
		{name: "zero", n: int64(0), want: 0},
		{name: "negative", n: int64(-42), want: -42},
		{name: "large positive", n: int64(1_000_000), want: 1_000_000},
		{name: "large negative", n: int64(-1_000_000), want: -1_000_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64ToInt(tt.n)
			if got != tt.want {
				t.Fatalf("Int64ToInt(%d)=%d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestRubToKop(t *testing.T) {
	tests := []struct {
		name string
		rub  Rub
		want Kop
	}{
		{name: "regular price", rub: 1200, want: 120000},
		{name: "zero", rub: 0, want: 0},
		{name: "one rub", rub: 1, want: 100},
		{name: "small price", rub: 15, want: 1500},
		{name: "negative value", rub: -10, want: -1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RubToKop(tt.rub)
			if got != tt.want {
				t.Fatalf("RubToKop(%d)=%d, want %d", tt.rub, got, tt.want)
			}
		})
	}
}

func TestKopToRub(t *testing.T) {
	tests := []struct {
		name string
		kop  Kop
		want Rub
	}{
		{name: "regular price", kop: 120000, want: 1200},
		{name: "zero", kop: 0, want: 0},
		{name: "one rub", kop: 100, want: 1},
		{name: "drops remainder", kop: 120050, want: 1200},
		{name: "negative value", kop: -1000, want: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KopToRub(tt.kop)
			if got != tt.want {
				t.Fatalf("KopToRub(%d)=%d, want %d", tt.kop, got, tt.want)
			}
		})
	}
}

func TestUserIDToString(t *testing.T) {
	tests := []struct {
		name string
		id   UserID
		want string
	}{
		{name: "regular id", id: 1001, want: "1001"},
		{name: "zero id", id: 0, want: "0"},
		{name: "one digit", id: 7, want: "7"},
		{name: "large id", id: 9000000000, want: "9000000000"},
		{name: "negative id", id: -1, want: "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UserIDToString(tt.id)
			if got != tt.want {
				t.Fatalf("UserIDToString(%d)=%q, want %q", tt.id, got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    int
		wantErr bool
	}{
		{name: "positive number", text: "42", want: 42, wantErr: false},
		{name: "zero", text: "0", want: 0, wantErr: false},
		{name: "negative number", text: "-42", want: -42, wantErr: false},
		{name: "invalid text", text: "oops", want: 0, wantErr: true},
		{name: "empty text", text: "", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.text)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseInt(%q) expected error, got nil", tt.text)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseInt(%q) unexpected error: %v", tt.text, err)
			}
			if got != tt.want {
				t.Fatalf("ParseInt(%q)=%d, want %d", tt.text, got, tt.want)
			}
		})
	}
}

func TestParseAndDouble(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    int
		wantErr bool
	}{
		{name: "positive number", text: "42", want: 84, wantErr: false},
		{name: "zero", text: "0", want: 0, wantErr: false},
		{name: "negative number", text: "-10", want: -20, wantErr: false},
		{name: "invalid text", text: "oops", want: 0, wantErr: true},
		{name: "empty text", text: "", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAndDouble(tt.text)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseAndDouble(%q) expected error, got nil", tt.text)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseAndDouble(%q) unexpected error: %v", tt.text, err)
			}
			if got != tt.want {
				t.Fatalf("ParseAndDouble(%q)=%d, want %d", tt.text, got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "positive", n: 42, want: "42"},
		{name: "zero", n: 0, want: "0"},
		{name: "negative", n: -42, want: "-42"},
		{name: "large positive", n: 1_000_000, want: "1000000"},
		{name: "one digit", n: 7, want: "7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToString(tt.n)
			if got != tt.want {
				t.Fatalf("IntToString(%d)=%q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  string
	}{
		{name: "one digit after point", value: 12.3, want: "12.30"},
		{name: "two digits after point", value: 12.34, want: "12.34"},
		{name: "round up", value: 12.345, want: "12.35"},
		{name: "zero", value: 0, want: "0.00"},
		{name: "negative", value: -12.3, want: "-12.30"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatToString(tt.value)
			if got != tt.want {
				t.Fatalf("FloatToString(%v)=%q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestParseBoolText(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    bool
		wantErr bool
	}{
		{name: "true", text: "true", want: true, wantErr: false},
		{name: "false", text: "false", want: false, wantErr: false},
		{name: "one", text: "1", want: true, wantErr: false},
		{name: "zero", text: "0", want: false, wantErr: false},
		{name: "invalid text", text: "yes", want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBoolText(tt.text)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseBoolText(%q) expected error, got nil", tt.text)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseBoolText(%q) unexpected error: %v", tt.text, err)
			}
			if got != tt.want {
				t.Fatalf("ParseBoolText(%q)=%t, want %t", tt.text, got, tt.want)
			}
		})
	}
}

func TestBoolToText(t *testing.T) {
	tests := []struct {
		name  string
		value bool
		want  string
	}{
		{name: "true", value: true, want: "true"},
		{name: "false", value: false, want: "false"},
		{name: "expression true", value: 10 > 5, want: "true"},
		{name: "expression false", value: 10 < 5, want: "false"},
		{name: "not false", value: !false, want: "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BoolToText(tt.value)
			if got != tt.want {
				t.Fatalf("BoolToText(%t)=%q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestSumIntAndInt64(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int64
		want int64
	}{
		{name: "positive numbers", a: 10, b: 20, want: 30},
		{name: "with zero int", a: 0, b: 20, want: 20},
		{name: "with zero int64", a: 10, b: 0, want: 10},
		{name: "negative int", a: -10, b: 20, want: 10},
		{name: "negative int64", a: 10, b: -20, want: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumIntAndInt64(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("SumIntAndInt64(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestPriceRubStringToKop(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    Kop
		wantErr bool
	}{
		{name: "regular price", text: "15", want: 1500, wantErr: false},
		{name: "zero price", text: "0", want: 0, wantErr: false},
		{name: "one rub", text: "1", want: 100, wantErr: false},
		{name: "invalid text", text: "oops", want: 0, wantErr: true},
		{name: "negative price", text: "-15", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PriceRubStringToKop(tt.text)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("PriceRubStringToKop(%q) expected error, got nil", tt.text)
				}
				return
			}

			if err != nil {
				t.Fatalf("PriceRubStringToKop(%q) unexpected error: %v", tt.text, err)
			}
			if got != tt.want {
				t.Fatalf("PriceRubStringToKop(%q)=%d, want %d", tt.text, got, tt.want)
			}
		})
	}
}

func TestSafeParsePositive(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "positive number", text: "25", want: 25},
		{name: "negative number", text: "-25", want: 0},
		{name: "zero", text: "0", want: 0},
		{name: "invalid text", text: "oops", want: 0},
		{name: "empty text", text: "", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SafeParsePositive(tt.text)
			if got != tt.want {
				t.Fatalf("SafeParsePositive(%q)=%d, want %d", tt.text, got, tt.want)
			}
		})
	}
}

func TestFormatUser(t *testing.T) {
	tests := []struct {
		name     string
		id       UserID
		userName string
		want     string
	}{
		{name: "regular user", id: 1001, userName: "Maria", want: "user:1001:Maria"},
		{name: "zero id", id: 0, userName: "Guest", want: "user:0:Guest"},
		{name: "empty name", id: 10, userName: "", want: "user:10:"},
		{name: "cyrillic name", id: 77, userName: "Мария", want: "user:77:Мария"},
		{name: "large id", id: 9000000000, userName: "Alex", want: "user:9000000000:Alex"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatUser(tt.id, tt.userName)
			if got != tt.want {
				t.Fatalf("FormatUser(%d, %q)=%q, want %q", tt.id, tt.userName, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "double=84 kop=120000 sum=30 user=1001"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
