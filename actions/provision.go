package actions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/paulieborg/aws-go-helper/stack"
)

// Provision a CloudFormation stack
func Provision(svc stack.Service, cfg stack.Config) (*string, error) {
	si := stack.StackInfo(&svc)
	sp := stack.Controller(&svc)
	sw := stack.StackWaiter(&svc)

	exists, _ := si.Exists(&cfg.StackName)
	rollback, _ := si.Rollback(&cfg.StackName)

	if exists && rollback {
		sp.Delete(&cfg.StackName)
		_, err := sp.Create(&cfg)

		if err != nil {
			log.Fatal(err)
		} else {
			sw.WaitCreate(&cfg.StackName)
		}

	} else if exists {
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

	describe, err := si.Describe(&cfg.StackName)

	if err != nil {
		return nil, err
	}

	return describe.Stacks[0].StackStatus, err
}
