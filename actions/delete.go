package actions

import (
	"github.com/paulieborg/aws-go-helper/stack"
)

// Delete ...
func Delete(svc stack.Service, name *string) (err error) {
	ctrl := stack.Controller(&svc)
	waiter := stack.Waiter(&svc)

	_, err = ctrl.Delete(name)

	if err != nil {
		return
	}

	return waiter.WaitDelete(name)
}
