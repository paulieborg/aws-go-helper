package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func Describe(p_args ProvisionArgs, ) (d *cf.DescribeStacksOutput, err error) {
	input := cf.DescribeStacksInput{StackName: &p_args.Stack_name}
	return p_args.Session.DescribeStacksWithContext(p_args.Context, &input)
}
