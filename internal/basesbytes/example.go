package basesbytes

import "fmt"

func Example() string {
	return fmt.Sprintf(
		"same=%t byte=%d bin=%s hex=%s packed=%d",
		SameNumber(),
		MaxByte(),
		FormatBinary(255),
		FormatHex(255),
		PackTwoBytes(1, 2),
	)
}
