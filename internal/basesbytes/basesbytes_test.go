package basesbytes

import "testing"

func TestDecimal42(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "equals 42", got: Decimal42(), want: 42},
		{name: "same as hex", got: Decimal42(), want: Hex2A()},
		{name: "same as binary", got: Decimal42(), want: Binary101010()},
		{name: "same as octal", got: Decimal42(), want: Octal52()},
		{name: "arithmetic check", got: Decimal42() + 8, want: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestHex2A(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "equals 42", got: Hex2A(), want: 42},
		{name: "same as decimal", got: Hex2A(), want: Decimal42()},
		{name: "same as binary", got: Hex2A(), want: Binary101010()},
		{name: "same as octal", got: Hex2A(), want: Octal52()},
		{name: "arithmetic check", got: Hex2A() * 2, want: 84},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestBinary101010(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "equals 42", got: Binary101010(), want: 42},
		{name: "same as decimal", got: Binary101010(), want: Decimal42()},
		{name: "same as hex", got: Binary101010(), want: Hex2A()},
		{name: "same as octal", got: Binary101010(), want: Octal52()},
		{name: "arithmetic check", got: Binary101010() - 2, want: 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestOctal52(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "equals 42", got: Octal52(), want: 42},
		{name: "same as decimal", got: Octal52(), want: Decimal42()},
		{name: "same as hex", got: Octal52(), want: Hex2A()},
		{name: "same as binary", got: Octal52(), want: Binary101010()},
		{name: "arithmetic check", got: Octal52() / 2, want: 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestSameNumber(t *testing.T) {
	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{name: "same number is true", got: SameNumber(), want: true},
		{name: "decimal equals hex", got: Decimal42() == Hex2A(), want: true},
		{name: "decimal equals binary", got: Decimal42() == Binary101010(), want: true},
		{name: "decimal equals octal", got: Decimal42() == Octal52(), want: true},
		{name: "all values are 42", got: Decimal42() == 42 && Hex2A() == 42 && Binary101010() == 42 && Octal52() == 42, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %t, want %t", tt.got, tt.want)
			}
		})
	}
}

func TestMaxByte(t *testing.T) {
	tests := []struct {
		name string
		got  byte
		want byte
	}{
		{name: "max byte", got: MaxByte(), want: byte(255)},
		{name: "greater than zero", got: MaxByte(), want: 255},
		{name: "same as uint8 max", got: MaxByte(), want: ^byte(0)},
		{name: "minus one equals 254", got: MaxByte() - 1, want: 254},
		{name: "half check", got: MaxByte() / 2, want: 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestFormatDecimal(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want string
	}{
		{name: "zero", n: 0, want: "0"},
		{name: "small number", n: 7, want: "7"},
		{name: "byte max", n: 255, want: "255"},
		{name: "power of two", n: 1024, want: "1024"},
		{name: "large number", n: 123456789, want: "123456789"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDecimal(tt.n)
			if got != tt.want {
				t.Fatalf("FormatDecimal(%d)=%q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestFormatBinary(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want string
	}{
		{name: "zero", n: 0, want: "0"},
		{name: "one", n: 1, want: "1"},
		{name: "two", n: 2, want: "10"},
		{name: "byte max", n: 255, want: "11111111"},
		{name: "sixty four", n: 64, want: "1000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatBinary(tt.n)
			if got != tt.want {
				t.Fatalf("FormatBinary(%d)=%q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestFormatOctal(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want string
	}{
		{name: "zero", n: 0, want: "0"},
		{name: "eight", n: 8, want: "10"},
		{name: "byte max", n: 255, want: "377"},
		{name: "sixty four", n: 64, want: "100"},
		{name: "forty two", n: 42, want: "52"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatOctal(tt.n)
			if got != tt.want {
				t.Fatalf("FormatOctal(%d)=%q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestFormatHex(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want string
	}{
		{name: "zero", n: 0, want: "0"},
		{name: "ten", n: 10, want: "a"},
		{name: "byte max", n: 255, want: "ff"},
		{name: "sixty four", n: 64, want: "40"},
		{name: "forty two", n: 42, want: "2a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatHex(tt.n)
			if got != tt.want {
				t.Fatalf("FormatHex(%d)=%q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestFormatAllBases(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want string
	}{
		{name: "zero", n: 0, want: "dec=0 bin=0 oct=0 hex=0"},
		{name: "one", n: 1, want: "dec=1 bin=1 oct=1 hex=1"},
		{name: "sixty four", n: 64, want: "dec=64 bin=1000000 oct=100 hex=40"},
		{name: "byte max", n: 255, want: "dec=255 bin=11111111 oct=377 hex=ff"},
		{name: "forty two", n: 42, want: "dec=42 bin=101010 oct=52 hex=2a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatAllBases(tt.n)
			if got != tt.want {
				t.Fatalf("FormatAllBases(%d)=%q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestIsASCII(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want bool
	}{
		{name: "capital letter", b: 'A', want: true},
		{name: "lowercase letter", b: 'z', want: true},
		{name: "digit", b: '7', want: true},
		{name: "ascii max", b: 127, want: true},
		{name: "not ascii", b: 200, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsASCII(tt.b)
			if got != tt.want {
				t.Fatalf("IsASCII(%d)=%t, want %t", tt.b, got, tt.want)
			}
		})
	}
}

func TestToUpperASCII(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want byte
	}{
		{name: "lowercase m", b: 'm', want: 'M'},
		{name: "lowercase a", b: 'a', want: 'A'},
		{name: "lowercase z", b: 'z', want: 'Z'},
		{name: "digit unchanged", b: '7', want: '7'},
		{name: "uppercase unchanged", b: 'Q', want: 'Q'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUpperASCII(tt.b)
			if got != tt.want {
				t.Fatalf("ToUpperASCII(%q)=%q, want %q", tt.b, got, tt.want)
			}
		})
	}
}

func TestToLowerASCII(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want byte
	}{
		{name: "uppercase M", b: 'M', want: 'm'},
		{name: "uppercase A", b: 'A', want: 'a'},
		{name: "uppercase Z", b: 'Z', want: 'z'},
		{name: "digit unchanged", b: '7', want: '7'},
		{name: "lowercase unchanged", b: 'q', want: 'q'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLowerASCII(tt.b)
			if got != tt.want {
				t.Fatalf("ToLowerASCII(%q)=%q, want %q", tt.b, got, tt.want)
			}
		})
	}
}

func TestPackTwoBytes(t *testing.T) {
	tests := []struct {
		name string
		high byte
		low  byte
		want uint16
	}{
		{name: "one and two", high: 1, low: 2, want: 258},
		{name: "zero and zero", high: 0, low: 0, want: 0},
		{name: "zero and max", high: 0, low: 255, want: 255},
		{name: "one and zero", high: 1, low: 0, want: 256},
		{name: "max and max", high: 255, low: 255, want: 65535},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PackTwoBytes(tt.high, tt.low)
			if got != tt.want {
				t.Fatalf("PackTwoBytes(%d, %d)=%d, want %d", tt.high, tt.low, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "same=true byte=255 bin=11111111 hex=ff packed=258"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
