package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
)

// createStack attempts to bring up a CloudFormation stack
func AwsCreateStack(
	svc *cf.CloudFormation,
	parameter []*cf.Parameter,
	name string,
	template string,
	compatibility string,
	timeout int64, ) (r *cf.CreateStackOutput, err error) {

	stack := &cf.CreateStackInput{
		StackName: aws.String(name),
		Capabilities: []*string{
			aws.String(compatibility),
		},
		Parameters:       parameter,
		TemplateBody:     aws.String(template),
		TimeoutInMinutes: aws.Int64(timeout),
	}

	return svc.CreateStack(stack)
}
