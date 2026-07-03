package integer

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 17, b: 5, want: 22},
		{name: "with zero", a: 10, b: 0, want: 10},
		{name: "negative and positive", a: -10, b: 3, want: -7},
		{name: "two negative numbers", a: -4, b: -6, want: -10},
		{name: "large numbers", a: 1000, b: 2500, want: 3500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Add(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 17, b: 5, want: 12},
		{name: "result is zero", a: 10, b: 10, want: 0},
		{name: "negative result", a: 5, b: 12, want: -7},
		{name: "subtract negative", a: 8, b: -2, want: 10},
		{name: "two negative numbers", a: -10, b: -3, want: -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Subtract(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 7, b: 6, want: 42},
		{name: "multiply by zero", a: 100, b: 0, want: 0},
		{name: "negative and positive", a: -5, b: 4, want: -20},
		{name: "two negative numbers", a: -3, b: -7, want: 21},
		{name: "multiply by one", a: 99, b: 1, want: 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Multiply(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
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
		{name: "division by zero", a: 10, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Divide(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Divide(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestRemainder(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "basic remainder", a: 17, b: 5, want: 2},
		{name: "no remainder", a: 20, b: 5, want: 0},
		{name: "number smaller than divisor", a: 3, b: 10, want: 3},
		{name: "negative dividend", a: -17, b: 5, want: -2},
		{name: "division by zero", a: 10, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Remainder(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Remainder(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "even positive", n: 42, want: true},
		{name: "odd positive", n: 17, want: false},
		{name: "zero is even", n: 0, want: true},
		{name: "even negative", n: -8, want: true},
		{name: "odd negative", n: -9, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.n)
			if got != tt.want {
				t.Fatalf("IsEven(%d)=%t, want %t", tt.n, got, tt.want)
			}
		})
	}
}

func TestLastDigit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "positive number", n: 987, want: 7},
		{name: "negative number", n: -987, want: 7},
		{name: "single digit", n: 5, want: 5},
		{name: "zero", n: 0, want: 0},
		{name: "ends with zero", n: 120, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastDigit(tt.n)
			if got != tt.want {
				t.Fatalf("LastDigit(%d)=%d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "first is greater", a: 20, b: 10, want: 20},
		{name: "second is greater", a: 10, b: 20, want: 20},
		{name: "equal numbers", a: 7, b: 7, want: 7},
		{name: "negative numbers", a: -10, b: -3, want: -3},
		{name: "negative and positive", a: -1, b: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Max(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "first is lower", a: 10, b: 20, want: 10},
		{name: "second is lower", a: 20, b: 10, want: 10},
		{name: "equal numbers", a: 7, b: 7, want: 7},
		{name: "negative numbers", a: -10, b: -3, want: -10},
		{name: "negative and positive", a: -1, b: 1, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Min(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name  string
		value int
		min   int
		max   int
		want  int
	}{
		{name: "lower than min", value: -5, min: 0, max: 100, want: 0},
		{name: "inside range", value: 55, min: 0, max: 100, want: 55},
		{name: "greater than max", value: 120, min: 0, max: 100, want: 100},
		{name: "equal min", value: 0, min: 0, max: 100, want: 0},
		{name: "equal max", value: 100, min: 0, max: 100, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Clamp(tt.value, tt.min, tt.max)
			if got != tt.want {
				t.Fatalf("Clamp(%d, %d, %d)=%d, want %d", tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestSumThree(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{name: "positive numbers", a: 1, b: 2, c: 3, want: 6},
		{name: "with zero", a: 10, b: 0, c: 5, want: 15},
		{name: "negative numbers", a: -1, b: -2, c: -3, want: -6},
		{name: "mixed numbers", a: -10, b: 5, c: 20, want: 15},
		{name: "all zero", a: 0, b: 0, c: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumThree(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Fatalf("SumThree(%d, %d, %d)=%d, want %d", tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "integer average with remainder", a: 5, b: 8, want: 6},
		{name: "exact average", a: 10, b: 20, want: 15},
		{name: "with zero", a: 0, b: 10, want: 5},
		{name: "negative numbers", a: -10, b: -20, want: -15},
		{name: "mixed numbers", a: -5, b: 5, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Average(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Average(%d, %d)=%d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIntToInt64(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{name: "positive number", n: 99, want: int64(99)},
		{name: "zero", n: 0, want: int64(0)},
		{name: "negative number", n: -99, want: int64(-99)},
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

func TestNonNegativeToUint(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want uint
	}{
		{name: "positive number", n: 10, want: uint(10)},
		{name: "zero", n: 0, want: uint(0)},
		{name: "negative number", n: -10, want: uint(0)},
		{name: "large positive", n: 999, want: uint(999)},
		{name: "minus one", n: -1, want: uint(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NonNegativeToUint(tt.n)
			if got != tt.want {
				t.Fatalf("NonNegativeToUint(%d)=%d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountPages(t *testing.T) {
	tests := []struct {
		name       string
		totalItems int
		pageSize   int
		want       int
	}{
		{name: "exact pages", totalItems: 100, pageSize: 10, want: 10},
		{name: "needs extra page", totalItems: 101, pageSize: 10, want: 11},
		{name: "zero items", totalItems: 0, pageSize: 10, want: 0},
		{name: "invalid page size", totalItems: 100, pageSize: 0, want: 0},
		{name: "one item per page", totalItems: 5, pageSize: 1, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountPages(tt.totalItems, tt.pageSize)
			if got != tt.want {
				t.Fatalf("CountPages(%d, %d)=%d, want %d", tt.totalItems, tt.pageSize, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "add=22 div=3 mod=2 pages=11 even=true clamp=100"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
