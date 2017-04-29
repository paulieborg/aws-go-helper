package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// createStack attempts to bring up a CloudFormation stack
func (p_args *ProvisionArgs) create() (err error) {

	stack := &cf.CreateStackInput{
		StackName: aws.String(p_args.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       p_args.Parameters,
		TimeoutInMinutes: aws.Int64(p_args.Timeout),
	}

	if p_args.Bucket == "" {
		stack.TemplateBody = aws.String(string(p_args.Template))
	} else {
		stack.TemplateURL = aws.String(p_args.Bucket)
	}

	_, err = p_args.Session.CreateStackWithContext(p_args.Context, stack)
	helpers.ErrorHandler(err)

	p_args.waitCreate()
	return
}
