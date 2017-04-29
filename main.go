package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/paulieborg/aws-go-helper/parsers"
	"github.com/paulieborg/aws-go-helper/actions"
	"github.com/paulieborg/aws-go-helper/helpers"
	"io/ioutil"
	"os"
)

var (
	action  = flag.String("a", "create", "create or delete")
	name    = flag.String("n", "", "Stack name.")
	tmpl    = flag.String("t", "network/test-template.yml", "Template file path.")
	params  = flag.String("p", "network/test-params.json", "Parameters file path.")
	timeout = flag.Int64("x", 5, "Timeout in minutes.")
)

func main() {
	ctx := context.Background()

	flag.Parse()

	t, err := readFile(*tmpl)
	helpers.ErrorHandler(err)

	p, err := readFile(*params)
	helpers.ErrorHandler(err)

	cfParams, err := parsers.ParseParams(p)
	helpers.ErrorHandler(err)

	sess := session.Must(session.NewSession())
	svc := cf.New(sess)

	if *action == "provision" {

		err := actions.Provision(ctx, svc, cfParams, *name, string(t), *timeout)
		helpers.ErrorHandler(err)

		ds, err := actions.Describe(ctx, svc, cf.DescribeStacksInput{StackName: name})
		helpers.ErrorHandler(err)

		fmt.Printf("Stack - %s\n", aws.StringValue(ds.Stacks[0].StackStatus))

	} else if *action == "delete" {
		_, err = actions.Delete(ctx, svc, cf.DeleteStackInput{StackName: name})
		helpers.ErrorHandler(err)

		actions.WaitDelete(ctx, svc, cf.DescribeStacksInput{StackName: name})

	} else {
		fmt.Printf("Unknown action '%s'\n", action)
	}
}

func readFile(f string) (t []byte, err error) {
	return ioutil.ReadFile(f)
}
