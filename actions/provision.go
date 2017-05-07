package actions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/paulieborg/aws-go-helper/stack"
)

// Provision a CloudFormation stack
func Provision(svc stack.Service, cfg stack.Config) (*string) {
	si := stack.StackInfo(&svc)
	sp := stack.Controller(&svc)
	sw := stack.StackWaiter(&svc)

	if si.Exists(&cfg.StackName) && si.Rollback(&cfg.StackName) {
		sp.Delete(&cfg.StackName)
		_, err := sp.Create(&cfg)

		if err != nil {
			log.Fatal(err)
		} else {
			sw.WaitCreate(&cfg.StackName)
		}

	} else if si.Exists(&cfg.StackName) {
		_, err := sp.Update(&cfg)

		if err != nil {
			if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
				fmt.Println("No updates are to be performed.")
				os.Exit(0)
			} else {
				log.Fatal(err)
			}
		} else {
			sw.WaitUpdate(&cfg.StackName)
		}
	} else {
		_, err := sp.Create(&cfg)

		if err != nil {
			log.Fatal(err)
		} else {
			sw.WaitCreate(&cfg.StackName)
		}
	}

	return si.Describe(&cfg.StackName).Stacks[0].StackStatus
}
