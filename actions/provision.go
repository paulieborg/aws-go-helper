package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
)

var capability string = "CAPABILITY_NAMED_IAM"

// Provision a CloudFormation stack
func Provision(
	ctx aws.Context,
	svc *cf.CloudFormation,
	parameter []*cf.Parameter,
	name string,
	template string,
	timeout int64, ) (err error) {

	if exists(ctx, svc, name) {
		err = update(ctx, svc, parameter, name, string(template))
	} else {
		err = create(ctx, svc, parameter, name, string(template), timeout)
	}

	return
}
