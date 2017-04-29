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
	Template   []byte
	Bucket     string
	Timeout    int64
}

// Provision a CloudFormation stack
func (p_args *ProvisionArgs) Provision() (err error) {

	if p_args.exists() && p_args.rollback() {
		p_args.Delete()
		helpers.ErrorHandler(err)
		p_args.WaitDelete()
		p_args.create()
	} else if p_args.exists() {
		err = p_args.update()
	} else {
		err = p_args.create()
	}

	return
}
