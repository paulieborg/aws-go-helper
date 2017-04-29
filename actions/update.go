package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"strings"
	"fmt"
)

// updateStack attempts to update an existing CloudFormation stack
func update(
	ctx aws.Context,
	svc *cf.CloudFormation,
	parameter []*cf.Parameter,
	name string,
	template string) (err error) {

	stack := &cf.UpdateStackInput{
		StackName: aws.String(name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:   parameter,
		TemplateBody: aws.String(template),
	}

	_, updateError := svc.UpdateStackWithContext(ctx, stack)

	if updateError != nil {
		if strings.Contains(updateError.Error(), "ValidationError: No updates are to be performed.") {
			fmt.Print("No updates are to be performed.\n")
			return
		}
		return updateError
	}

	waitUpdate(ctx, svc, cf.DescribeStacksInput{StackName: &name})

	return

}
