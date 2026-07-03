package conversion

import "fmt"

func Example() string {
	doubled, _ := ParseAndDouble("42")
	kop := RubToKop(1200)
	sum := SumIntAndInt64(10, 20)
	return fmt.Sprintf("double=%d kop=%d sum=%d user=%s", doubled, kop, sum, UserIDToString(1001))
}
