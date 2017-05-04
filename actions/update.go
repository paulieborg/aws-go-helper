package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"strings"
	"fmt"
	"os"
)

// updateStack attempts to update an existing CloudFormation stack
func update(stack *StackArgs) (*cf.DescribeStacksOutput, error) {

	var err error = nil

	stackInput := &cf.UpdateStackInput{
		StackName: aws.String(stack.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: stack.Parameters,
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

	_, err = stack.Session.UpdateStackWithContext(stack.Context, stackInput)

	if err != nil {
		if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
			fmt.Println("No updates are to be performed.")
			os.Exit(0)
		} else {
			return nil, err
		}
	}

	err = waitUpdate(stack)

	if err != nil {
		return nil, err
	}

	return describe(stack), err

}
