package	actions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/paulieborg/aws-go-helper/stack"
)

// Provision a CloudFormation stack
func Provision(svc stack.Service, cfg stack.Config)	(status	*string, err error)	{
	i	   := stack.Info(&svc)
	ctrl   := stack.Controller(&svc)
	waiter := stack.Waiter(&svc)

	switch {
	case i.Exists(&cfg.StackName) && i.Rollback(&cfg.StackName):
		ctrl.Delete(&cfg.StackName)	 //	should the user	be warned about	this?
		_, err = ctrl.Create(&cfg)
		if err != nil {
			log.Fatal(err)
			return
		}
		waiter.WaitCreate(&cfg.StackName)

	case i.Exists(&cfg.StackName):
		_, err = ctrl.Update(&cfg)
		if err != nil {
    		// TODO if AWS change this text, this'll break
    		if strings.Contains(err.Error(), "Validation Error: No updates are to be performed.") {
        		fmt.Println(err.Error())
        		os.Exit(1)
    		}
			log.Fatal(err)
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

	return i.Describe(&cfg.StackName).Stacks[0].StackStatus, nil
}
