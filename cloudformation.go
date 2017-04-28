package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"time"

	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	//  "github.com/aws/aws-sdk-go/aws/awserr"
)

var (
	name    = flag.String("n", "", "Stack name.")
	tmpl    = flag.String("t", "network/test-template.yml", "Template file path.")
	params  = flag.String("p", "network/test-params.json", "Parameters file path.")
	timeout = flag.Int64("x", 5, "Timeout in minutes.")
)

const capability string = "CAPABILITY_NAMED_IAM"
const ctx_timeout time.Duration = 0

func main() {
	ctx := context.Background()

	flag.Parse()

	t, err := ioutil.ReadFile(*tmpl)
	if err != nil {
		panic(err)
	}

	p, err := ioutil.ReadFile(*params)
	if err != nil {
		panic(err)
	}

	cfParams, err := parseParams(p)
	if err != nil {
		panic(err)
	}

	sess := session.Must(session.NewSession())
	svc := cf.New(sess)

	r, err := createStack(ctx, svc, cfParams, string(t))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", r)

	waitCreate(ctx, svc, cf.DescribeStacksInput{StackName: name})

}

// parseParams takes a simple JSON blob of parameters and converts it to a slice
// of CloudFormation parameter structs
func parseParams(params []byte) (p []*cf.Parameter, err error) {
	var sp map[string]string

	err = json.NewDecoder(bytes.NewReader(params)).Decode(&sp)
	if err != nil {
		return
	}

	for k, v := range sp {
		pkey := k
		pval := v
		p = append(p, &cf.Parameter{
			ParameterKey:   &pkey,
			ParameterValue: &pval,
		})
	}

	return
}

// createStack attempts to bring up a CloudFormation stack
func createStack(ctx aws.Context, svc *cf.CloudFormation, p []*cf.Parameter, t string) (r *cf.CreateStackOutput, err error) {
	stack := &cf.CreateStackInput{
		StackName: aws.String(*name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       p,
		TemplateBody:     aws.String(t),
		TimeoutInMinutes: aws.Int64(*timeout),
	}

	return svc.CreateStackWithContext(ctx, stack)
}

func waitCreate(ctx aws.Context, svc *cf.CloudFormation, stackInfo cf.DescribeStacksInput) {

	svc.WaitUntilStackCreateCompleteWithContext(
		ctx,
		&stackInfo,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
	)

	return
}

//  if err != nil {
//      if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
//          fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
//      } else {
//          fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
//      }
//      os.Exit(1)
//  }
