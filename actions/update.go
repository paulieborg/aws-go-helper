package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"strings"
	"fmt"
	"os"
)

// updateStack attempts to update an existing CloudFormation stack
func (s *Stack) update(p ProvisionArgs) {

	sess := cf.New(session.Must(session.NewSession()))

	stack := &cf.UpdateStackInput{
		StackName: aws.String(p.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: p.Parameters,
	}

	if p.BucketName == "" {
		stack.TemplateBody = aws.String(string(p.Template))
	} else {
		stack.TemplateURL = aws.String(s.s3upload(p))
	}

	_, err := sess.UpdateStackWithContext(s.Context, stack)

	if err != nil {
		if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
			fmt.Printf("%v\n", err.Error())
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}

	s.waitUpdate(p)

}
