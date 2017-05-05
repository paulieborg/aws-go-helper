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
func (c *Context) update(p ProvisionArgs) {

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
		stack.TemplateURL = aws.String(c.s3upload(p))
	}

	_, err := c.Service.UpdateStackWithContext(c.Context, stack)

	if err != nil {
		if strings.Contains(err.Error(), "ValidationError: No updates are to be performed.") {
			fmt.Println("No updates are to be performed.")
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}

	c.waitUpdate(p)

}
