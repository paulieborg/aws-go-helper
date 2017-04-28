package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	//  "github.com/aws/aws-sdk-go/aws/awserr"
	//  "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type simpleParams map[string]string

var (
	name   = flag.String("n", "", "Stack name.")
	tmpl   = flag.String("t", "network/template.yml", "Template file path.")
	params = flag.String("p", "network/params.json", "Parameters file path.")
)

var capability string = "CAPABILITY_NAMED_IAM"

func main() {
	flag.Parse()

	t, err := getReader(*tmpl)
	if err != nil {
		panic(err)
	}
	raw_t, err := ioutil.ReadAll(t)
	if err != nil {
		panic(err)
	}

	p, err := getReader(*params)
	if err != nil {
		panic(err)
	}

	cf_params, err := parseParams(p)
	if err != nil {
		panic(err)
	}

	sess := session.Must(session.NewSession())
	svc := cloudformation.New(sess)
	r, err := createStack(svc, cf_params, string(raw_t))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", r)
}

func getReader(p string) (r *bufio.Reader, err error) {
	f, err := os.Open(p)
	if err != nil {
		return
	}

	return bufio.NewReader(f), nil
}

func parseParams(b *bufio.Reader) (p []*cloudformation.Parameter, err error) {
	var s_params simpleParams
	err = json.NewDecoder(b).Decode(&s_params)
	if err != nil {
		return
	}

	for k, v := range s_params {
		pkey := k
		pval := v
		p = append(p, &cloudformation.Parameter{
			ParameterKey:   &pkey,
			ParameterValue: &pval,
		})
	}

	return
}

func createStack(svc *cloudformation.CloudFormation, p []*cloudformation.Parameter, t string) (r *cloudformation.CreateStackOutput, err error) {
	stack := &cloudformation.CreateStackInput{
		StackName: aws.String(*name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       p,
		TemplateBody:     aws.String(t),
		TimeoutInMinutes: aws.Int64(5),
	}

	return svc.CreateStack(stack)
}

//  if err != nil {
//      if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
//          fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
//      } else {
//          fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
//      }
//      os.Exit(1)
//  }
// 
//  fmt.Println(resp)
// 
//  stackInfo := cloudformation.DescribeStacksInput{StackName: &name}
//  svc.WaitUntilStackCreateComplete(&stackInfo)
