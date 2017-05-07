package actions

import (
	"log"

	"github.com/paulieborg/aws-go-helper/stack"
)

const status = "DELETE_COMPLETE"

// Delete ...
func Delete(svc stack.Service, cfg stack.Config) (status *string, err error) {
    ctrl := stack.Controller(&svc)
    waiter := stack.Waiter(&svc)

	_, err = ctrl.Delete(&cfg.StackName)

	if err != nil {
		log.Fatal(err)
		return
	}

    waiter.WaitDelete(&cfg.StackName)

	return
}
