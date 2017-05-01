package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/paulieborg/aws-go-helper/helpers"
)

// updateStack attempts to update an existing CloudFormation stack
func (s *StackArgs) update() (err error) {

	stack := &cf.UpdateStackInput{
		StackName: aws.String(s.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: s.Parameters,
	}

	if s.Bucket == "" {
		stack.TemplateBody = aws.String(string(s.Template))
	} else {
		path, err := s.s3upload()
		helpers.ErrorHandler(err)
		stack.TemplateURL = aws.String(path)
	}

	_, err = s.Session.UpdateStackWithContext(s.Context, stack)
	helpers.ErrorHandler(err)

	s.waitUpdate()

	return

}
