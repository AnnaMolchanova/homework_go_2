package integer

import "fmt"

// Example собирает несколько функций из блока integer в одну строку.
// Main должен только вызвать Example и напечатать результат.
func Example() string {
	return fmt.Sprintf(
		"add=%d div=%d mod=%d pages=%d even=%t clamp=%d",
		Add(17, 5),
		Divide(17, 5),
		Remainder(17, 5),
		CountPages(101, 10),
		IsEven(42),
		Clamp(120, 0, 100),
	)
}
