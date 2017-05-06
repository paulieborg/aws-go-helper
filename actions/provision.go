package actions

import (
	"github.com/paulieborg/aws-go-helper/stack"
	"strings"
	"fmt"
	"os"
	"log"
)

// Provision a CloudFormation stack
func Provision(service stack.StackService, config stack.StackConfig) (*string) {

	si := stack.StackInfo(&service)
	sp := stack.StackController(&service)
	sw := stack.StackWaiter(&service)

	if si.Exists(&config.Stack_name) && si.Rollback(&config.Stack_name) {
		sp.Delete(&config.Stack_name)
		_, err := sp.Create(&config)

		if err != nil {
			log.Fatal(err)
		} else {
			sw.WaitCreate(&config.Stack_name)
		}

	} else if si.Exists(&config.Stack_name) {
		_, err := sp.Update(&config)

		if err != nil {
			if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
				fmt.Println("No updates are to be performed.")
				os.Exit(0)
			} else {
				log.Fatal(err)
			}
		} else {
			sw.WaitUpdate(&config.Stack_name)
		}

	} else {

		_, err := sp.Create(&config)

		if err != nil {
			log.Fatal(err)
		} else {
			sw.WaitCreate(&config.Stack_name)
		}
	}

	return si.Describe(&config.Stack_name).Stacks[0].StackStatus
}
