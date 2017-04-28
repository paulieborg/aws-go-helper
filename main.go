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
	"io/ioutil"
	"strings"
)

var (
	action  = flag.String("a", "create", "create or delete")
	name    = flag.String("n", "", "Stack name.")
	tmpl    = flag.String("t", "network/test-template.yml", "Template file path.")
	params  = flag.String("p", "network/test-params.json", "Parameters file path.")
	timeout = flag.Int64("x", 5, "Timeout in minutes.")
	verbose = flag.Bool("v", false, "Verbose Output")
)

const capability string = "CAPABILITY_NAMED_IAM"

func main() {
	ctx := context.Background()

	flag.Parse()

	t, err := readFile(*tmpl)
	if err != nil {
		panic(err)
	}

	p, err := readFile(*params)
	if err != nil {
		panic(err)
	}
	cfParams, err := parsers.ParseParams(p)
	if err != nil {
		panic(err)
	}

	sess := session.Must(session.NewSession())
	svc := cf.New(sess)

	if ( *action == "create") {

		r, err := actions.AwsCreateStack(ctx, svc, cfParams, *name, string(t), capability, *timeout)
		if err != nil {
			panic(err)
		}

		if *verbose {
			fmt.Printf("%+v\n", r)
		}

		actions.AwsWaitCreateStack(ctx, svc, cf.DescribeStacksInput{StackName: name})
		ds, err := actions.AwsDescribeStacks(ctx, svc, cf.DescribeStacksInput{StackName: name})
		if err != nil {
			panic(err)
		}

		if (strings.TrimRight(aws.StringValue(ds.Stacks[0].StackStatus), "\n") == "CREATE_COMPLETE") && *verbose {
			fmt.Printf("Stack - %+v\n", aws.StringValue(ds.Stacks[0].StackStatus))
		}

	} else if ( *action == "delete" ) {
		_, err = actions.AwsDeleteStack(ctx, svc, cf.DeleteStackInput{StackName: name})
		if err != nil {
			panic(err)
		}

		actions.AwsWaitDelete(ctx, svc, cf.DescribeStacksInput{StackName: name})

	} else {
		fmt.Printf("Unknown action '%+v'\n", action)
	}
}

func readFile(f string) (t []byte, err error) {
	return ioutil.ReadFile(f)
}
