package command

import (
	"fmt"
)

func verifyHost() error {
	if currentContext.Host == "" {
		return fmt.Errorf("host is not set")
	}

	return nil
}
