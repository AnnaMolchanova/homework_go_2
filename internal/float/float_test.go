package float

import (
	"math"
	"testing"
)

func assertFloat(t *testing.T, got, want float64) {
	t.Helper()

	if math.Abs(got-want) > 0.000001 {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "positive numbers", a: 1.5, b: 2.25, want: 3.75},
		{name: "with zero", a: 10.5, b: 0, want: 10.5},
		{name: "negative and positive", a: -5.5, b: 2.5, want: -3},
		{name: "two negative numbers", a: -1.25, b: -2.75, want: -4},
		{name: "small decimals", a: 0.1, b: 0.2, want: 0.3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, Add(tt.a, tt.b), tt.want)
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "positive numbers", a: 10.5, b: 0.5, want: 10},
		{name: "result is zero", a: 5.5, b: 5.5, want: 0},
		{name: "negative result", a: 3.25, b: 5.25, want: -2},
		{name: "subtract negative", a: 7.5, b: -2.5, want: 10},
		{name: "two negative numbers", a: -10.5, b: -2.5, want: -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, Subtract(tt.a, tt.b), tt.want)
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "positive numbers", a: 2.5, b: 4, want: 10},
		{name: "multiply by zero", a: 100.5, b: 0, want: 0},
		{name: "negative and positive", a: -3.5, b: 2, want: -7},
		{name: "two negative numbers", a: -2.5, b: -4, want: 10},
		{name: "fraction result", a: 0.5, b: 0.5, want: 0.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, Multiply(tt.a, tt.b), tt.want)
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "fraction result", a: 10, b: 4, want: 2.5},
		{name: "exact division", a: 20, b: 5, want: 4},
		{name: "negative division", a: -9, b: 3, want: -3},
		{name: "divide by larger number", a: 3, b: 10, want: 0.3},
		{name: "division by zero", a: 10, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, Divide(tt.a, tt.b), tt.want)
		})
	}
}

func TestDiscountPrice(t *testing.T) {
	tests := []struct {
		name    string
		price   float64
		percent float64
		want    float64
	}{
		{name: "regular discount", price: 1000, percent: 15, want: 850},
		{name: "zero discount", price: 1000, percent: 0, want: 1000},
		{name: "negative discount", price: 1000, percent: -10, want: 1000},
		{name: "full discount", price: 1000, percent: 100, want: 0},
		{name: "small price", price: 199.90, percent: 20, want: 159.92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, DiscountPrice(tt.price, tt.percent), tt.want)
		})
	}
}

func TestAddTax(t *testing.T) {
	tests := []struct {
		name       string
		price      float64
		taxPercent float64
		want       float64
	}{
		{name: "regular tax", price: 199.90, taxPercent: 20, want: 239.88},
		{name: "zero tax", price: 100, taxPercent: 0, want: 100},
		{name: "negative tax", price: 100, taxPercent: -10, want: 100},
		{name: "ten percent", price: 1000, taxPercent: 10, want: 1100},
		{name: "fraction price", price: 10.50, taxPercent: 20, want: 12.60},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, AddTax(tt.price, tt.taxPercent), tt.want)
		})
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name    string
		celsius float64
		want    float64
	}{
		{name: "zero celsius", celsius: 0, want: 32},
		{name: "boiling point", celsius: 100, want: 212},
		{name: "freezing equal point", celsius: -40, want: -40},
		{name: "room temperature", celsius: 20, want: 68},
		{name: "fraction celsius", celsius: 37.5, want: 99.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, CelsiusToFahrenheit(tt.celsius), tt.want)
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name       string
		fahrenheit float64
		want       float64
	}{
		{name: "freezing point", fahrenheit: 32, want: 0},
		{name: "boiling point", fahrenheit: 212, want: 100},
		{name: "equal point", fahrenheit: -40, want: -40},
		{name: "room temperature", fahrenheit: 68, want: 20},
		{name: "fraction result", fahrenheit: 99.5, want: 37.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, FahrenheitToCelsius(tt.fahrenheit), tt.want)
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "integer-like values", a: 10, b: 15, want: 12.5},
		{name: "equal values", a: 7.5, b: 7.5, want: 7.5},
		{name: "with zero", a: 0, b: 10, want: 5},
		{name: "negative values", a: -10, b: -20, want: -15},
		{name: "mixed values", a: -5, b: 5, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, Average(tt.a, tt.b), tt.want)
		})
	}
}

func TestRound2(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{name: "round up", value: 12.345, want: 12.35},
		{name: "round down", value: 12.344, want: 12.34},
		{name: "already two digits", value: 12.30, want: 12.30},
		{name: "negative round", value: -12.345, want: -12.35},
		{name: "zero", value: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, Round2(tt.value), tt.want)
		})
	}
}

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		name  string
		price float64
		want  string
	}{
		{name: "one digit after point", price: 12.3, want: "12.30"},
		{name: "two digits after point", price: 12.34, want: "12.34"},
		{name: "round up", price: 12.345, want: "12.35"},
		{name: "zero", price: 0, want: "0.00"},
		{name: "integer price", price: 100, want: "100.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatPrice(tt.price)
			if got != tt.want {
				t.Fatalf("FormatPrice(%v)=%q, want %q", tt.price, got, tt.want)
			}
		})
	}
}

func TestPercentOf(t *testing.T) {
	tests := []struct {
		name    string
		total   float64
		percent float64
		want    float64
	}{
		{name: "ten percent", total: 250, percent: 10, want: 25},
		{name: "zero percent", total: 250, percent: 0, want: 0},
		{name: "full percent", total: 250, percent: 100, want: 250},
		{name: "half percent", total: 200, percent: 12.5, want: 25},
		{name: "negative percent", total: 200, percent: -10, want: -20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, PercentOf(tt.total, tt.percent), tt.want)
		})
	}
}

func TestGrowthPercent(t *testing.T) {
	tests := []struct {
		name     string
		oldValue float64
		newValue float64
		want     float64
	}{
		{name: "growth", oldValue: 100, newValue: 125, want: 25},
		{name: "decrease", oldValue: 100, newValue: 80, want: -20},
		{name: "no growth", oldValue: 100, newValue: 100, want: 0},
		{name: "double growth", oldValue: 50, newValue: 100, want: 100},
		{name: "old value zero", oldValue: 0, newValue: 100, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFloat(t, GrowthPercent(tt.oldValue, tt.newValue), tt.want)
		})
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  bool
	}{
		{name: "positive", value: 0.1, want: true},
		{name: "negative", value: -1, want: false},
		{name: "zero", value: 0, want: false},
		{name: "large positive", value: 1000.5, want: true},
		{name: "large negative", value: -1000.5, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPositive(tt.value)
			if got != tt.want {
				t.Fatalf("IsPositive(%v)=%t, want %t", tt.value, got, tt.want)
			}
		})
	}
}

func TestFloatToInt(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  int
	}{
		{name: "positive fraction", value: 12.99, want: 12},
		{name: "negative fraction", value: -12.99, want: -12},
		{name: "integer-like value", value: 42.0, want: 42},
		{name: "small positive", value: 0.99, want: 0},
		{name: "small negative", value: -0.99, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatToInt(tt.value)
			if got != tt.want {
				t.Fatalf("FloatToInt(%v)=%d, want %d", tt.value, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "discount=850.00 tax=239.88 avg=12.5 rounded=12.35"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
