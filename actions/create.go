package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
)

var capability string = "CAPABILITY_NAMED_IAM"

// createStack attempts to bring up a CloudFormation stack
func Create(
	ctx aws.Context,
	svc *cf.CloudFormation,
	parameter []*cf.Parameter,
	name string,
	template string,
	timeout int64, ) (r *cf.CreateStackOutput, err error) {

	stack := &cf.CreateStackInput{
		StackName: aws.String(name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       parameter,
		TemplateBody:     aws.String(template),
		TimeoutInMinutes: aws.Int64(timeout),
	}

	return svc.CreateStackWithContext(ctx, stack)
}
