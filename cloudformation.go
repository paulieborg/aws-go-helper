package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type paramFactory struct {
	paramKey   string
	paramValue string
}

type params_type map[string]string

func main() {

	var name, template_file, parameters string
	var capability string = "CAPABILITY_NAMED_IAM"

	flag.StringVar(&name, "n", "", "Name")
	flag.StringVar(&template_file, "t", "network/template.yml", "Template File Name.")
	flag.StringVar(&parameters, "p", "network/params.json", "Params File Name.")
	flag.Parse()

	var data params_type

	template, err := ioutil.ReadFile(template_file)
	check(err)

	params, err := ioutil.ReadFile(parameters)
	check(err)

	err = json.Unmarshal(params, &data)
	check(err)

	var key string
	var val string

	var cf_params []*cloudformation.Parameter

	for key, val = range data {

		cf_params = append(cf_params, formatParams(&paramFactory{paramKey: key, paramValue: val}))
	}

	sess := session.Must(session.NewSession())
	svc := cloudformation.New(sess)

	stack := &cloudformation.CreateStackInput{
		StackName: aws.String(name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       cf_params,
		TemplateBody:     aws.String(string(template)),
		TimeoutInMinutes: aws.Int64(5),
	}

	resp, err := svc.CreateStack(stack)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Println(resp)

	stackInfo := cloudformation.DescribeStacksInput{StackName: &name}
	svc.WaitUntilStackCreateComplete(&stackInfo)

}

func formatParams(data *paramFactory) *cloudformation.Parameter {

	var cf_params *cloudformation.Parameter

	cf_params = &cloudformation.Parameter{
		ParameterKey:   &data.paramKey,
		ParameterValue: &data.paramValue,
	}

	return cf_params
}
