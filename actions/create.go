package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// createStack attempts to bring up a CloudFormation stack
func create(p_args ProvisionArgs) (err error) {

	stack := &cf.CreateStackInput{
		StackName: aws.String(p_args.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       p_args.Parameters,
		TemplateBody:     aws.String(p_args.Template),
		TimeoutInMinutes: aws.Int64(p_args.Timeout),
	}

	_, err = p_args.Session.CreateStackWithContext(p_args.Context, stack)
	helpers.ErrorHandler(err)

	waitCreate(p_args)
	return
}
