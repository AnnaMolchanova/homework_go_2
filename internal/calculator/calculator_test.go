package calculator

import "testing"

func TestCalculateAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 10, b: 5, want: 15},
		{name: "with zero", a: 10, b: 0, want: 10},
		{name: "negative and positive", a: -10, b: 5, want: -5},
		{name: "two negative numbers", a: -10, b: -5, want: -15},
		{name: "large numbers", a: 1000, b: 2500, want: 3500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, "+")
			if err != nil {
				t.Fatalf("Calculate(%d, %d, +) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Fatalf("Calculate(%d, %d, +)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculateSubtract(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 10, b: 5, want: 5},
		{name: "result is zero", a: 10, b: 10, want: 0},
		{name: "negative result", a: 5, b: 10, want: -5},
		{name: "subtract negative", a: 10, b: -5, want: 15},
		{name: "two negative numbers", a: -10, b: -5, want: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, "-")
			if err != nil {
				t.Fatalf("Calculate(%d, %d, -) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Fatalf("Calculate(%d, %d, -)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculateMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 10, b: 5, want: 50},
		{name: "multiply by zero", a: 10, b: 0, want: 0},
		{name: "negative and positive", a: -10, b: 5, want: -50},
		{name: "two negative numbers", a: -10, b: -5, want: 50},
		{name: "multiply by one", a: 99, b: 1, want: 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, "*")
			if err != nil {
				t.Fatalf("Calculate(%d, %d, *) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Fatalf("Calculate(%d, %d, *)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculateDivide(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "integer division", a: 17, b: 5, want: 3},
		{name: "exact division", a: 20, b: 5, want: 4},
		{name: "negative division", a: -20, b: 5, want: -4},
		{name: "division by larger number", a: 3, b: 10, want: 0},
		{name: "two negative numbers", a: -20, b: -5, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, "/")
			if err != nil {
				t.Fatalf("Calculate(%d, %d, /) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Fatalf("Calculate(%d, %d, /)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculateRemainder(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "regular remainder", a: 17, b: 5, want: 2},
		{name: "no remainder", a: 20, b: 5, want: 0},
		{name: "number smaller than divisor", a: 3, b: 10, want: 3},
		{name: "negative dividend", a: -17, b: 5, want: -2},
		{name: "two negative numbers", a: -17, b: -5, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, "%")
			if err != nil {
				t.Fatalf("Calculate(%d, %d, %%) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Fatalf("Calculate(%d, %d, %%)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculateErrors(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		operation string
	}{
		{name: "division by zero", a: 1, b: 0, operation: "/"},
		{name: "remainder by zero", a: 1, b: 0, operation: "%"},
		{name: "unknown question mark", a: 1, b: 2, operation: "?"},
		{name: "unknown word", a: 1, b: 2, operation: "add"},
		{name: "empty operation", a: 1, b: 2, operation: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, tt.operation)
			if err == nil {
				t.Fatalf(
					"Calculate(%d, %d, %q)=%d, want error",
					tt.a,
					tt.b,
					tt.operation,
					got,
				)
			}
		})
	}
}

func TestCalculateOriginalCases(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		op   string
		want int
	}{
		{name: "add", a: 10, b: 5, op: "+", want: 15},
		{name: "subtract", a: 10, b: 5, op: "-", want: 5},
		{name: "multiply", a: 10, b: 5, op: "*", want: 50},
		{name: "divide", a: 17, b: 5, op: "/", want: 3},
		{name: "mod", a: 17, b: 5, op: "%", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.b, tt.op)
			if err != nil {
				t.Fatalf("Calculate(%d, %d, %q) unexpected error: %v", tt.a, tt.b, tt.op, err)
			}
			if got != tt.want {
				t.Fatalf("Calculate(%d, %d, %q)=%d, want %d", tt.a, tt.b, tt.op, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "add=15 div=3"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
