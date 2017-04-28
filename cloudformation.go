package main

import (
    "bytes"
    "flag"
    "fmt"

    "encoding/json"
    "io/ioutil"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    cf "github.com/aws/aws-sdk-go/service/cloudformation"
//  "github.com/aws/aws-sdk-go/aws/awserr"
//  "github.com/aws/aws-sdk-go/aws/request"
)

var (
    name = flag.String("n", "", "Stack name.")
    tmpl = flag.String("t", "network/template.yml", "Template file path.")
    params = flag.String("p", "network/params.json", "Parameters file path.")
    timeout = flag.Int64("x", 0, "Timeout in minutes.")
)

const capability string = "CAPABILITY_NAMED_IAM"

func main() {
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

    r, err := createStack(svc, cfParams, string(t))
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", r)

    stackInfo := cf.DescribeStacksInput{StackName: name}
    svc.WaitUntilStackCreateComplete(&stackInfo)
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
            ParameterKey: &pkey,
            ParameterValue: &pval,
        })
    }

    return
}

// createStack attempts to bring up a CloudFormation stack
func createStack(svc *cf.CloudFormation, p []*cf.Parameter, t string) (r *cf.CreateStackOutput, err error) {
    stack := &cf.CreateStackInput{
        StackName: aws.String(*name),
        Capabilities: []*string{
            aws.String(capability),
        },
        Parameters: p,
        TemplateBody: aws.String(t),
        TimeoutInMinutes: aws.Int64(*timeout),
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