package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

var capability string = "CAPABILITY_NAMED_IAM"

type StackArgs struct {
	Context    aws.Context
	Session    *cf.CloudFormation
	Parameters []*cf.Parameter
	Stack_name string
	Template   []byte
	Bucket     string
	Timeout    int64
}

// Provision a CloudFormation stack
func Provision(stack *StackArgs) (*string) {

	var result *cf.DescribeStacksOutput
	var err error

	ds := describe(stack)

	if exists(ds) && rollback(ds) {
		err = Delete(stack)

		if err != nil {
			handleError(err)
		} else {
			result, err = create(stack)
		}

	} else if exists(ds) {
		result, err = update(stack)
		handleError(err)
	} else {
		result, err = create(stack)
		handleError(err)
	}

	return result.Stacks[0].StackStatus
}

func exists(ds *cf.DescribeStacksOutput) (bool) {

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}

func rollback(ds *cf.DescribeStacksOutput) (bool) {

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}