package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// updateStack attempts to update an existing CloudFormation stack
func update(
	p_args ProvisionArgs) (err error) {

	stack := &cf.UpdateStackInput{
		StackName: aws.String(p_args.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:   p_args.Parameters,
		TemplateBody: aws.String(p_args.Template),
	}

	_, err = p_args.Session.UpdateStackWithContext(p_args.Context, stack)
	helpers.ErrorHandler(err)

	waitUpdate(p_args)

	return

}
