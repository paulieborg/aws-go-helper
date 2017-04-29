package helpers

import (
	"strings"
	"fmt"
	"os"
)

func ErrorHandler(err error) {

	if err != nil {

		if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
			fmt.Print("No updates are to be performed.\n")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "ExpiredToken: The security token included in the request is expired") {
			fmt.Print("The security token included in the request is expired.\n")
			os.Exit(0)
		} else {
			panic(err)
		}
	}

	return
}
