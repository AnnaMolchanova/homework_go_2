package float

import "fmt"

func Example() string {
	return fmt.Sprintf(
		"discount=%.2f tax=%.2f avg=%.1f rounded=%.2f",
		DiscountPrice(1000, 15),
		AddTax(199.90, 20),
		Average(10, 15),
		Round2(12.345),
	)
}
