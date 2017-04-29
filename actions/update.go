package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// updateStack attempts to update an existing CloudFormation stack
func update(
	ctx aws.Context,
	svc *cf.CloudFormation,
	parameter []*cf.Parameter,
	name string,
	template string) (err error) {

	stack := &cf.UpdateStackInput{
		StackName: aws.String(name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:   parameter,
		TemplateBody: aws.String(template),
	}

	_, err = svc.UpdateStackWithContext(ctx, stack)
	helpers.ErrorHandler(err)

	waitUpdate(ctx, svc, cf.DescribeStacksInput{StackName: &name})

	return

}
