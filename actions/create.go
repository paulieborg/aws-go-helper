package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
)

// createStack attempts to bring up a CloudFormation stack
func create(
	ctx aws.Context,
	svc *cf.CloudFormation,
	parameter []*cf.Parameter,
	name string,
	template string,
	timeout int64, ) (err error) {

	stack := &cf.CreateStackInput{
		StackName: aws.String(name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       parameter,
		TemplateBody:     aws.String(template),
		TimeoutInMinutes: aws.Int64(timeout),
	}

	_, err = svc.CreateStackWithContext(ctx, stack)

	waitCreate(ctx, svc, cf.DescribeStacksInput{StackName: &name})

	return err
}
