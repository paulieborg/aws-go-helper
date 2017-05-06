package actions

import (
	"log"
	"github.com/paulieborg/aws-go-helper/stack"
)

// Provision a CloudFormation stack
func Delete(service stack.StackService, config stack.StackConfig) (*string) {

	var response = "DELETE_COMPLETE"

	sp := stack.StackController(&service)
	sw := stack.StackWaiter(&service)

	_, err := sp.Delete(&config.Stack_name)

	if err != nil {
		log.Fatal(err)
	} else {
		sw.WaitDelete(&config.Stack_name)
	}

	return &response
}
