package boolean

import "fmt"

func Example() string {
	return fmt.Sprintf(
		"enter=%t rest=%t access=%t leap=%t withdraw=%t",
		CanEnter(20, true),
		CanRest(false, true),
		HasAccess(false, true),
		IsLeapYear(2024),
		CanWithdraw(1000, 700, false),
	)
}
