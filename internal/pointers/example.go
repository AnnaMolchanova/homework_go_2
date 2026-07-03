package pointers

import "fmt"

func Example() string {
	value := 10
	afterInc := Increment(&value)
	nilValue := ValueOrDefault(nil, 5)
	DoubleInPlace(&value)
	return fmt.Sprintf("afterInc=%d value=%d nil=%d", afterInc, value, nilValue)
}
