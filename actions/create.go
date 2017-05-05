package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"log"
)

// createStack attempts to bring up a CloudFormation stack
func (s *Stack) create(p ProvisionArgs) {

	sess := cf.New(session.Must(session.NewSession()))

	stackInput := &cf.CreateStackInput{
		StackName: aws.String(p.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       p.Parameters,
		TimeoutInMinutes: aws.Int64(p.Timeout),
	}

	if p.BucketName == "" {
		stackInput.TemplateBody = aws.String(string(p.Template))
	} else {
		stackInput.TemplateURL = aws.String(s.s3upload(p))
	}

	_, err := sess.CreateStackWithContext(s.Context, stackInput)

	if err != nil {
		log.Fatal(err)
	} else {
		s.waitCreate(p)
	}

	return
}
