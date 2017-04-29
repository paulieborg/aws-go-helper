package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/paulieborg/aws-go-helper/helpers"
)

var capability string = "CAPABILITY_NAMED_IAM"

type ProvisionArgs struct {
	Context    aws.Context
	Session    *cf.CloudFormation
	Parameters []*cf.Parameter
	Stack_name string
	Template   string
	Bucket   string
	Timeout    int64
}

// Provision a CloudFormation stack
func Provision(p_args ProvisionArgs, ) (err error) {

	if exists(p_args) && rolledback(p_args) {
		Delete(p_args)
		helpers.ErrorHandler(err)
		WaitDelete(p_args)
		create(p_args)
	} else if exists(p_args) {
		err = update(p_args)
	} else {
		err = create(p_args)
	}

	return
}
