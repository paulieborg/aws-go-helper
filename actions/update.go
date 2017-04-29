package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// updateStack attempts to update an existing CloudFormation stack
func (p_args *ProvisionArgs) update() (err error) {

	stack := &cf.UpdateStackInput{
		StackName: aws.String(p_args.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: p_args.Parameters,
	}

	if p_args.Bucket == "" {
		stack.TemplateBody = aws.String(string(p_args.Template))
	} else {
		stack.TemplateURL = aws.String(p_args.Bucket)
	}

	_, err = p_args.Session.UpdateStackWithContext(p_args.Context, stack)
	helpers.ErrorHandler(err)

	p_args.waitUpdate()

	return

}
