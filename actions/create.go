package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"log"
)

// createStack attempts to bring up a CloudFormation stack
func (s *StackArgs) create() {

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
		stackInput.TemplateURL = aws.String(s.s3upload())
	}

	_, err := s.Session.CreateStackWithContext(s.Context, stackInput)

	if err != nil {
		log.Fatal(err)
	} else {
		s.waitCreate()
	}

	return
}
