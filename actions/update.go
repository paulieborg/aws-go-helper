package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"strings"
	"fmt"
	"os"
)

// updateStack attempts to update an existing CloudFormation stack
func (s *StackArgs) update() {

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
		stack.TemplateURL = aws.String(s.s3upload())
	}

	_, err := s.Session.UpdateStackWithContext(s.Context, stack)

	if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
		fmt.Printf("%v\n", err.Error())
		os.Exit(0)
	} else if err != nil {
		log.Fatal(err)
	} else {
		s.waitUpdate()
	}

}
