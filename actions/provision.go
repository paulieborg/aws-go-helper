package actions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/paulieborg/aws-go-helper/stack"
)

// Provision a CloudFormation stack

func Provision(svc stack.Service, cfg stack.Config) (status string, err error) {
	i := stack.Info(&svc)
	ctrl := stack.Controller(&svc)
	waiter := stack.Waiter(&svc)
	exists, _ := i.Exists(&cfg.StackName)
	rollback, _ := i.Rollback(&cfg.StackName)

	switch {
	case exists && rollback:
		ctrl.Delete(&cfg.StackName) //	should the user	be warned about	this?
		_, err = ctrl.Create(&cfg)
		if err != nil {
			log.Fatal(err)
			return
		}
		waiter.WaitCreate(&cfg.StackName)

	case exists:
		_, err = ctrl.Update(&cfg)
		if err != nil {
			// TODO if AWS change this text, this'll break
			if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
				fmt.Println("No updates are to be performed")
				os.Exit(0)
			}
			return
		}
		waiter.WaitUpdate(&cfg.StackName)

	default:
		_, err = ctrl.Create(&cfg)
		if err != nil {
			log.Fatal(err)
			return
		}
		waiter.WaitCreate(&cfg.StackName)
	}

	describe, err := i.Describe(&cfg.StackName)

	if err != nil {
		return
	}

	return *describe.Stacks[0].StackStatus, err
}
