package text

import "testing"

func TestByteLen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "cyrillic", s: "Привет", want: 12},
		{name: "ascii", s: "hello", want: 5},
		{name: "empty", s: "", want: 0},
		{name: "mixed", s: "Go Привет", want: 15},
		{name: "emoji", s: "🙂", want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteLen(tt.s)
			if got != tt.want {
				t.Fatalf("ByteLen(%q)=%d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestRuneLen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "cyrillic", s: "Привет", want: 6},
		{name: "ascii", s: "hello", want: 5},
		{name: "empty", s: "", want: 0},
		{name: "mixed", s: "Go Привет", want: 9},
		{name: "emoji", s: "🙂", want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuneLen(tt.s)
			if got != tt.want {
				t.Fatalf("RuneLen(%q)=%d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestFirstRune(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "cyrillic", s: "Привет", want: "П"},
		{name: "ascii", s: "hello", want: "h"},
		{name: "empty", s: "", want: ""},
		{name: "emoji first", s: "🙂go", want: "🙂"},
		{name: "one rune", s: "Я", want: "Я"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstRune(tt.s)
			if got != tt.want {
				t.Fatalf("FirstRune(%q)=%q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestLastRune(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "cyrillic", s: "Привет", want: "т"},
		{name: "ascii", s: "hello", want: "o"},
		{name: "empty", s: "", want: ""},
		{name: "emoji last", s: "go🙂", want: "🙂"},
		{name: "one rune", s: "Я", want: "Я"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastRune(tt.s)
			if got != tt.want {
				t.Fatalf("LastRune(%q)=%q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "spaces", s: "  Go  ", want: "Go"},
		{name: "tabs", s: "\tGo\t", want: "Go"},
		{name: "new lines", s: "\nGo\n", want: "Go"},
		{name: "no spaces", s: "Go", want: "Go"},
		{name: "only spaces", s: "   ", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trim(tt.s)
			if got != tt.want {
				t.Fatalf("Trim(%q)=%q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "mixed ascii", s: "GoLang", want: "golang"},
		{name: "upper ascii", s: "GO", want: "go"},
		{name: "already lower", s: "go", want: "go"},
		{name: "cyrillic", s: "ПРИВЕТ", want: "привет"},
		{name: "empty", s: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLower(tt.s)
			if got != tt.want {
				t.Fatalf("ToLower(%q)=%q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "lower ascii", s: "go", want: "GO"},
		{name: "mixed ascii", s: "GoLang", want: "GOLANG"},
		{name: "already upper", s: "GO", want: "GO"},
		{name: "cyrillic", s: "привет", want: "ПРИВЕТ"},
		{name: "empty", s: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUpper(tt.s)
			if got != tt.want {
				t.Fatalf("ToUpper(%q)=%q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{name: "spaces and upper", email: "  MARIA@EXAMPLE.COM  ", want: "maria@example.com"},
		{name: "already normalized", email: "user@test.com", want: "user@test.com"},
		{name: "mixed case", email: "User@Test.Com", want: "user@test.com"},
		{name: "tabs", email: "\tADMIN@SITE.RU\t", want: "admin@site.ru"},
		{name: "empty", email: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeEmail(tt.email)
			if got != tt.want {
				t.Fatalf("NormalizeEmail(%q)=%q, want %q", tt.email, got, tt.want)
			}
		})
	}
}

func TestContainsWord(t *testing.T) {
	tests := []struct {
		name string
		text string
		word string
		want bool
	}{
		{name: "contains word", text: "hello go", word: "go", want: true},
		{name: "does not contain word", text: "hello rust", word: "go", want: false},
		{name: "contains cyrillic", text: "привет go", word: "привет", want: true},
		{name: "empty word", text: "hello", word: "", want: true},
		{name: "case sensitive", text: "Hello Go", word: "go", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsWord(tt.text, tt.word)
			if got != tt.want {
				t.Fatalf("ContainsWord(%q, %q)=%t, want %t", tt.text, tt.word, got, tt.want)
			}
		})
	}
}

func TestReplaceFirstRune(t *testing.T) {
	tests := []struct {
		name string
		s    string
		r    rune
		want string
	}{
		{name: "cyrillic word", s: "кот", r: 'р', want: "рот"},
		{name: "ascii word", s: "cat", r: 'b', want: "bat"},
		{name: "empty", s: "", r: 'x', want: ""},
		{name: "emoji", s: "🙂go", r: 'G', want: "Ggo"},
		{name: "one rune", s: "я", r: 'т', want: "т"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceFirstRune(tt.s, tt.r)
			if got != tt.want {
				t.Fatalf("ReplaceFirstRune(%q, %q)=%q, want %q", tt.s, tt.r, got, tt.want)
			}
		})
	}
}

func TestReverseRunes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "cyrillic", s: "кот", want: "ток"},
		{name: "ascii", s: "go", want: "og"},
		{name: "empty", s: "", want: ""},
		{name: "one rune", s: "я", want: "я"},
		{name: "emoji", s: "go🙂", want: "🙂og"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseRunes(tt.s)
			if got != tt.want {
				t.Fatalf("ReverseRunes(%q)=%q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestInitials(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		want      string
	}{
		{name: "latin names", firstName: "Maria", lastName: "Ivanova", want: "MI"},
		{name: "lowercase cyrillic", firstName: "алексей", lastName: "петров", want: "АП"},
		{name: "with spaces", firstName: "  maria  ", lastName: "  ivanova  ", want: "MI"},
		{name: "empty first name", firstName: "", lastName: "Ivanova", want: "I"},
		{name: "empty last name", firstName: "Maria", lastName: "", want: "M"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Initials(tt.firstName, tt.lastName)
			if got != tt.want {
				t.Fatalf("Initials(%q, %q)=%q, want %q", tt.firstName, tt.lastName, got, tt.want)
			}
		})
	}
}

func TestRepeatWord(t *testing.T) {
	tests := []struct {
		name  string
		word  string
		count int
		want  string
	}{
		{name: "three times", word: "go", count: 3, want: "gogogo"},
		{name: "zero times", word: "go", count: 0, want: ""},
		{name: "negative count", word: "go", count: -1, want: ""},
		{name: "one time", word: "go", count: 1, want: "go"},
		{name: "cyrillic", word: "да", count: 2, want: "дада"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RepeatWord(tt.word, tt.count)
			if got != tt.want {
				t.Fatalf("RepeatWord(%q, %d)=%q, want %q", tt.word, tt.count, got, tt.want)
			}
		})
	}
}

func TestJoinWithComma(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		want   string
	}{
		{name: "three values", values: []string{"go", "test", "ci"}, want: "go,test,ci"},
		{name: "one value", values: []string{"go"}, want: "go"},
		{name: "empty slice", values: []string{}, want: ""},
		{name: "nil slice", values: nil, want: ""},
		{name: "values with spaces", values: []string{"go", "unit test", "ci"}, want: "go,unit test,ci"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinWithComma(tt.values)
			if got != tt.want {
				t.Fatalf("JoinWithComma(%#v)=%q, want %q", tt.values, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "cyrillic phrase", s: "А роза упала на лапу Азора", want: true},
		{name: "simple cyrillic", s: "топот", want: true},
		{name: "ascii with case", s: "Madam", want: true},
		{name: "not palindrome", s: "Привет", want: false},
		{name: "spaces only", s: "   ", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.s)
			if got != tt.want {
				t.Fatalf("IsPalindrome(%q)=%t, want %t", tt.s, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "bytes=12 runes=6 first=П changed=рот email=maria@example.com"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
