package constants

import "fmt"

func Example() string {
	return fmt.Sprintf(
		"app=%s attempts=%d status=%s role=%s final=%t",
		AppName(),
		MaxAttempts(),
		StatusText(StatusPaid),
		RoleText(RoleAdmin),
		IsFinalStatus(StatusCanceled),
	)
}
