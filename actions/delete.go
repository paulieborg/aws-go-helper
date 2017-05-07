package actions

import (
	"log"

	"github.com/paulieborg/aws-go-helper/stack"
)

// Delete ...
func Delete(svc stack.Service, cfg stack.Config) (resp string) {
	resp = "DELETE_COMPLETE"

	ctrl := stack.Controller(&svc)
	waiter := stack.StackWaiter(&svc)

	_, err := ctrl.Delete(&cfg.StackName)

	if err != nil {
		log.Fatal(err)
	} else {
		waiter.WaitDelete(&cfg.StackName)
	}

	return resp
}
