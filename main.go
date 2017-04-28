package main

import (
	"flag"
	"fmt"

	"io/ioutil"

	"github.com/paulieborg/aws-go-helper/parsers"
	"github.com/paulieborg/aws-go-helper/actions"
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

var (
	name    = flag.String("n", "", "Stack name.")
	tmpl    = flag.String("t", "network/template.yml", "Template file path.")
	params  = flag.String("p", "network/params.json", "Parameters file path.")
	timeout = flag.Int64("x", 5, "Timeout in minutes.")
)

const capability string = "CAPABILITY_NAMED_IAM"

func main() {
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

	r, err := actions.AwsCreateStack(svc, cfParams, *name, string(t), capability, *timeout)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", r)

	stackInfo := cf.DescribeStacksInput{StackName: name}
	svc.WaitUntilStackCreateComplete(&stackInfo)
}

func readFile(f string) (t []byte, err error) {
	return ioutil.ReadFile(f)
}
