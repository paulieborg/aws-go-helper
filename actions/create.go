package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
)

// createStack attempts to bring up a CloudFormation stack
func create(stack *StackArgs) (*cf.DescribeStacksOutput, error) {

	var err error = nil

	stackInput := &cf.CreateStackInput{
		StackName: aws.String(stack.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       stack.Parameters,
		TimeoutInMinutes: aws.Int64(stack.Timeout),
	}

	if stack.Bucket == "" {
		stackInput.TemplateBody = aws.String(string(stack.Template))
	} else {
		path, err := s3upload(stack)
		if err != nil {
			return nil, err
		}
		stackInput.TemplateURL = aws.String(path)
	}

	_, err = stack.Session.CreateStackWithContext(stack.Context, stackInput)

	if err != nil {
		return nil, err
	}

	err = waitCreate(stack)

	if err != nil {
		return nil, err
	}

	return describe(stack), err
}
