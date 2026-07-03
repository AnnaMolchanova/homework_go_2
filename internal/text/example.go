package text

import "fmt"

func Example() string {
	changed := ReplaceFirstRune("кот", 'р')
	return fmt.Sprintf(
		"bytes=%d runes=%d first=%s changed=%s email=%s",
		ByteLen("Привет"),
		RuneLen("Привет"),
		FirstRune("Привет"),
		changed,
		NormalizeEmail("  MARIA@EXAMPLE.COM  "),
	)
}
