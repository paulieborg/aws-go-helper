package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// createStack attempts to bring up a CloudFormation stack
func (s *StackArgs) create() (err error) {

	stackInput := &cf.CreateStackInput{
		StackName: aws.String(s.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       s.Parameters,
		TimeoutInMinutes: aws.Int64(s.Timeout),
	}

	if s.Bucket == "" {
		stackInput.TemplateBody = aws.String(string(s.Template))
	} else {
		path, err := s.s3upload()
		helpers.ErrorHandler(err)
		stackInput.TemplateURL = aws.String(path)
	}

	_, err = s.Session.CreateStackWithContext(s.Context, stackInput)
	helpers.ErrorHandler(err)

	s.waitCreate()
	return
}
