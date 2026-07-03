package calculator

import "fmt"

func Example() string {
	add, _ := Calculate(10, 5, "+")
	div, _ := Calculate(17, 5, "/")
	return fmt.Sprintf("add=%d div=%d", add, div)
}
