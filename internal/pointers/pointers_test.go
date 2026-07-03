package pointers

import "testing"

func TestValueOrDefault(t *testing.T) {
	tests := []struct {
		name string
		p    *int
		def  int
		want int
	}{
		{name: "nil pointer", p: nil, def: 5, want: 5},
		{name: "regular value", p: NewInt(10), def: 5, want: 10},
		{name: "zero value", p: NewInt(0), def: 5, want: 0},
		{name: "negative value", p: NewInt(-10), def: 5, want: -10},
		{name: "negative default", p: nil, def: -1, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValueOrDefault(tt.p, tt.def)
			if got != tt.want {
				t.Fatalf("ValueOrDefault(%v, %d)=%d, want %d", tt.p, tt.def, got, tt.want)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	t.Run("regular value", func(t *testing.T) {
		v := 10
		got := Increment(&v)

		if got != 11 || v != 11 {
			t.Fatalf("got %d value %d, want 11 and 11", got, v)
		}
	})

	t.Run("zero value", func(t *testing.T) {
		v := 0
		got := Increment(&v)

		if got != 1 || v != 1 {
			t.Fatalf("got %d value %d, want 1 and 1", got, v)
		}
	})

	t.Run("negative value", func(t *testing.T) {
		v := -2
		got := Increment(&v)

		if got != -1 || v != -1 {
			t.Fatalf("got %d value %d, want -1 and -1", got, v)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		got := Increment(nil)

		if got != 0 {
			t.Fatalf("got %d, want 0", got)
		}
	})

	t.Run("called twice", func(t *testing.T) {
		v := 5
		Increment(&v)
		got := Increment(&v)

		if got != 7 || v != 7 {
			t.Fatalf("got %d value %d, want 7 and 7", got, v)
		}
	})
}

func TestSetValue(t *testing.T) {
	t.Run("set positive", func(t *testing.T) {
		v := 1
		ok := SetValue(&v, 7)

		if !ok || v != 7 {
			t.Fatalf("ok=%t value=%d, want true and 7", ok, v)
		}
	})

	t.Run("set zero", func(t *testing.T) {
		v := 10
		ok := SetValue(&v, 0)

		if !ok || v != 0 {
			t.Fatalf("ok=%t value=%d, want true and 0", ok, v)
		}
	})

	t.Run("set negative", func(t *testing.T) {
		v := 10
		ok := SetValue(&v, -5)

		if !ok || v != -5 {
			t.Fatalf("ok=%t value=%d, want true and -5", ok, v)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		ok := SetValue(nil, 7)

		if ok {
			t.Fatalf("ok=%t, want false", ok)
		}
	})

	t.Run("overwrite old value", func(t *testing.T) {
		v := 100
		SetValue(&v, 50)
		ok := SetValue(&v, 25)

		if !ok || v != 25 {
			t.Fatalf("ok=%t value=%d, want true and 25", ok, v)
		}
	})
}

func TestSwap(t *testing.T) {
	t.Run("regular values", func(t *testing.T) {
		a, b := 1, 2
		ok := Swap(&a, &b)

		if !ok || a != 2 || b != 1 {
			t.Fatalf("ok=%t a=%d b=%d, want true 2 1", ok, a, b)
		}
	})

	t.Run("negative values", func(t *testing.T) {
		a, b := -1, -2
		ok := Swap(&a, &b)

		if !ok || a != -2 || b != -1 {
			t.Fatalf("ok=%t a=%d b=%d, want true -2 -1", ok, a, b)
		}
	})

	t.Run("same values", func(t *testing.T) {
		a, b := 5, 5
		ok := Swap(&a, &b)

		if !ok || a != 5 || b != 5 {
			t.Fatalf("ok=%t a=%d b=%d, want true 5 5", ok, a, b)
		}
	})

	t.Run("first nil", func(t *testing.T) {
		b := 2
		ok := Swap(nil, &b)

		if ok || b != 2 {
			t.Fatalf("ok=%t b=%d, want false and 2", ok, b)
		}
	})

	t.Run("second nil", func(t *testing.T) {
		a := 1
		ok := Swap(&a, nil)

		if ok || a != 1 {
			t.Fatalf("ok=%t a=%d, want false and 1", ok, a)
		}
	})
}

func TestResetToZero(t *testing.T) {
	tests := []struct {
		name  string
		start int
		want  int
	}{
		{name: "positive", start: 99, want: 0},
		{name: "negative", start: -99, want: 0},
		{name: "already zero", start: 0, want: 0},
		{name: "one", start: 1, want: 0},
		{name: "large", start: 100000, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.start
			ok := ResetToZero(&v)

			if !ok || v != tt.want {
				t.Fatalf("ok=%t value=%d, want true and %d", ok, v, tt.want)
			}
		})
	}

	t.Run("nil pointer", func(t *testing.T) {
		if ResetToZero(nil) {
			t.Fatalf("ResetToZero(nil)=true, want false")
		}
	})
}

func TestAddToValue(t *testing.T) {
	tests := []struct {
		name  string
		start int
		delta int
		want  int
	}{
		{name: "add positive", start: 10, delta: 5, want: 15},
		{name: "add negative", start: 10, delta: -3, want: 7},
		{name: "add to zero", start: 0, delta: 5, want: 5},
		{name: "add zero", start: 10, delta: 0, want: 10},
		{name: "negative start", start: -10, delta: 5, want: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.start
			got := AddToValue(&v, tt.delta)

			if got != tt.want || v != tt.want {
				t.Fatalf("got %d value %d, want %d", got, v, tt.want)
			}
		})
	}

	t.Run("nil pointer", func(t *testing.T) {
		got := AddToValue(nil, 5)

		if got != 0 {
			t.Fatalf("got %d, want 0", got)
		}
	})
}

func TestMaxPointer(t *testing.T) {
	t.Run("second is larger", func(t *testing.T) {
		a, b := 10, 20
		got := MaxPointer(&a, &b)

		if got != &b {
			t.Fatalf("want pointer to b")
		}
	})

	t.Run("first is larger", func(t *testing.T) {
		a, b := 30, 20
		got := MaxPointer(&a, &b)

		if got != &a {
			t.Fatalf("want pointer to a")
		}
	})

	t.Run("equal returns first", func(t *testing.T) {
		a, b := 20, 20
		got := MaxPointer(&a, &b)

		if got != &a {
			t.Fatalf("want pointer to a")
		}
	})

	t.Run("first nil", func(t *testing.T) {
		b := 20
		got := MaxPointer(nil, &b)

		if got != &b {
			t.Fatalf("want pointer to b")
		}
	})

	t.Run("both nil", func(t *testing.T) {
		got := MaxPointer(nil, nil)

		if got != nil {
			t.Fatalf("got %v, want nil", got)
		}
	})
}

func TestIsNil(t *testing.T) {
	tests := []struct {
		name string
		p    *int
		want bool
	}{
		{name: "nil pointer", p: nil, want: true},
		{name: "positive value", p: NewInt(1), want: false},
		{name: "zero value", p: NewInt(0), want: false},
		{name: "negative value", p: NewInt(-1), want: false},
		{name: "new int", p: NewInt(100), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNil(tt.p)
			if got != tt.want {
				t.Fatalf("IsNil(%v)=%t, want %t", tt.p, got, tt.want)
			}
		})
	}
}

func TestCopyValue(t *testing.T) {
	tests := []struct {
		name string
		p    *int
		want int
	}{
		{name: "positive value", p: NewInt(12), want: 12},
		{name: "zero value", p: NewInt(0), want: 0},
		{name: "negative value", p: NewInt(-12), want: -12},
		{name: "nil pointer", p: nil, want: 0},
		{name: "large value", p: NewInt(100000), want: 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CopyValue(tt.p)
			if got != tt.want {
				t.Fatalf("CopyValue(%v)=%d, want %d", tt.p, got, tt.want)
			}
		})
	}
}

func TestDoubleInPlace(t *testing.T) {
	tests := []struct {
		name  string
		start int
		want  int
	}{
		{name: "positive", start: 6, want: 12},
		{name: "zero", start: 0, want: 0},
		{name: "negative", start: -6, want: -12},
		{name: "one", start: 1, want: 2},
		{name: "large", start: 1000, want: 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.start
			ok := DoubleInPlace(&v)

			if !ok || v != tt.want {
				t.Fatalf("ok=%t value=%d, want true and %d", ok, v, tt.want)
			}
		})
	}

	t.Run("nil pointer", func(t *testing.T) {
		if DoubleInPlace(nil) {
			t.Fatalf("DoubleInPlace(nil)=true, want false")
		}
	})
}

func TestNewInt(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "positive", value: 15, want: 15},
		{name: "zero", value: 0, want: 0},
		{name: "negative", value: -15, want: -15},
		{name: "one", value: 1, want: 1},
		{name: "large", value: 100000, want: 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt(tt.value)
			if got == nil {
				t.Fatalf("NewInt(%d)=nil, want pointer", tt.value)
			}
			if *got != tt.want {
				t.Fatalf("*NewInt(%d)=%d, want %d", tt.value, *got, tt.want)
			}
		})
	}
}

func TestDivideInto(t *testing.T) {
	t.Run("regular division", func(t *testing.T) {
		out := 0
		ok := DivideInto(&out, 17, 5)

		if !ok || out != 3 {
			t.Fatalf("ok=%t out=%d, want true and 3", ok, out)
		}
	})

	t.Run("exact division", func(t *testing.T) {
		out := 0
		ok := DivideInto(&out, 20, 5)

		if !ok || out != 4 {
			t.Fatalf("ok=%t out=%d, want true and 4", ok, out)
		}
	})

	t.Run("negative division", func(t *testing.T) {
		out := 0
		ok := DivideInto(&out, -20, 5)

		if !ok || out != -4 {
			t.Fatalf("ok=%t out=%d, want true and -4", ok, out)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		out := 99
		ok := DivideInto(&out, 1, 0)

		if ok || out != 99 {
			t.Fatalf("ok=%t out=%d, want false and unchanged 99", ok, out)
		}
	})

	t.Run("nil output", func(t *testing.T) {
		ok := DivideInto(nil, 10, 2)

		if ok {
			t.Fatalf("DivideInto(nil, 10, 2)=true, want false")
		}
	})
}

func TestApplyDiscountInPlace(t *testing.T) {
	tests := []struct {
		name    string
		price   int
		percent int
		want    int
		ok      bool
	}{
		{name: "regular discount", price: 1000, percent: 15, want: 850, ok: true},
		{name: "zero discount", price: 1000, percent: 0, want: 1000, ok: true},
		{name: "full discount", price: 1000, percent: 100, want: 0, ok: true},
		{name: "more than full discount", price: 1000, percent: 120, want: 0, ok: true},
		{name: "negative discount", price: 1000, percent: -10, want: 1000, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price := tt.price
			ok := ApplyDiscountInPlace(&price, tt.percent)

			if ok != tt.ok || price != tt.want {
				t.Fatalf("ok=%t price=%d, want ok=%t price=%d", ok, price, tt.ok, tt.want)
			}
		})
	}

	t.Run("nil price", func(t *testing.T) {
		if ApplyDiscountInPlace(nil, 15) {
			t.Fatalf("ApplyDiscountInPlace(nil, 15)=true, want false")
		}
	})
}

func TestChoosePointer(t *testing.T) {
	t.Run("primary exists", func(t *testing.T) {
		a, b := 1, 2
		got := ChoosePointer(&a, &b)

		if got != &a {
			t.Fatalf("want primary pointer")
		}
	})

	t.Run("primary nil fallback exists", func(t *testing.T) {
		b := 2
		got := ChoosePointer(nil, &b)

		if got != &b {
			t.Fatalf("want fallback pointer")
		}
	})

	t.Run("both nil", func(t *testing.T) {
		got := ChoosePointer(nil, nil)

		if got != nil {
			t.Fatalf("got %v, want nil", got)
		}
	})

	t.Run("primary zero value exists", func(t *testing.T) {
		a, b := 0, 2
		got := ChoosePointer(&a, &b)

		if got != &a {
			t.Fatalf("want primary pointer even when value is zero")
		}
	})

	t.Run("fallback zero value exists", func(t *testing.T) {
		b := 0
		got := ChoosePointer(nil, &b)

		if got != &b {
			t.Fatalf("want fallback pointer even when value is zero")
		}
	})
}

func TestPointToLarger(t *testing.T) {
	t.Run("second is larger", func(t *testing.T) {
		a, b := 1, 2
		got := PointToLarger(&a, &b)

		if got != &b {
			t.Fatalf("want pointer to b")
		}
	})

	t.Run("first is larger", func(t *testing.T) {
		a, b := 3, 2
		got := PointToLarger(&a, &b)

		if got != &a {
			t.Fatalf("want pointer to a")
		}
	})

	t.Run("equal returns first", func(t *testing.T) {
		a, b := 2, 2
		got := PointToLarger(&a, &b)

		if got != &a {
			t.Fatalf("want pointer to a")
		}
	})

	t.Run("first nil", func(t *testing.T) {
		b := 2
		got := PointToLarger(nil, &b)

		if got != &b {
			t.Fatalf("want pointer to b")
		}
	})

	t.Run("both nil", func(t *testing.T) {
		got := PointToLarger(nil, nil)

		if got != nil {
			t.Fatalf("got %v, want nil", got)
		}
	})
}

func TestExample(t *testing.T) {
	want := "afterInc=11 value=22 nil=5"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
