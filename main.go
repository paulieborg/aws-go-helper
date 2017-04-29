package main

import (
	"context"
	"flag"
	"fmt"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/paulieborg/aws-go-helper/actions"
	"github.com/paulieborg/aws-go-helper/helpers"
	"github.com/paulieborg/aws-go-helper/parsers"
)

var (
	action   = flag.String("a", "create", "create or delete")
	name     = flag.String("n", "", "Stack name.")
	template = flag.String("t", "network/test-template.yml", "Template file path.")
	params   = flag.String("p", "network/test-params.json", "Parameters file path.")
	bucket   = flag.String("b", "", "Bucket containing template.")
	timeout  = flag.Int64("x", 5, "Timeout in minutes.")
)

func main() {
	ctx := context.Background()

	flag.Parse()

	t, err := readFile(*template)
	helpers.ErrorHandler(err)

	p, err := readFile(*params)
	helpers.ErrorHandler(err)

	cfParams, err := parsers.ParseParams(p)
	helpers.ErrorHandler(err)

	s := session.Must(session.NewSession())
	svc := cf.New(s)

	p_args := actions.ProvisionArgs{
		Context:          ctx,
		Session:          svc,
		Parameters:       cfParams,
		Stack_name:       *name,
		Template:         t,
		TemplateFileName: *template,
		Bucket:           *bucket,
		Timeout:          *timeout, }

	switch *action {
	case "provision":
		err := p_args.Provision()
		helpers.ErrorHandler(err)
		ds, err := p_args.Describe()
		helpers.ErrorHandler(err)
		fmt.Printf("Stack - %s\n", aws.StringValue(ds.Stacks[0].StackStatus))
	case "delete":
		_, err = p_args.Delete()
		helpers.ErrorHandler(err)
		p_args.WaitDelete()
	default:
		fmt.Printf("Unknown action '%s'\n", *action)
	}
}

func readFile(f string) (t []byte, err error) {
	return ioutil.ReadFile(f)
}
