package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

var capability string = "CAPABILITY_NAMED_IAM"

type ProvisionArgs struct {
	Context    aws.Context
	Session    *cf.CloudFormation
	Parameters []*cf.Parameter
	Stack_name string
	Template   string
	Timeout    int64
}

// Provision a CloudFormation stack
func Provision(p_args ProvisionArgs, ) (err error) {

	if exists(p_args) {
		err = update(p_args)
	} else {
		err = create(p_args)
	}

	return
}
