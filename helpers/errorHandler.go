package helpers

import (
	"strings"
	"fmt"
	"os"
	"log"
)

func ErrorHandler(err error) {

	if err != nil {

		if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
			fmt.Printf("%v\n", err.Error())
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
}
