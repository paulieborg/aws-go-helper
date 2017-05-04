package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func describe(stack *StackArgs) (*cf.DescribeStacksOutput) {
	input := cf.DescribeStacksInput{StackName: &stack.Stack_name}
	d, _ := stack.Session.DescribeStacksWithContext(stack.Context, &input)
	return d
}
