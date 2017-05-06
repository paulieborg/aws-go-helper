package actions

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type StackProvider interface {
	exists(*string) (bool)
	describe(*string) (*cloudformation.DescribeStacksOutput)
	rollback(*string) (bool)
}

func StackInfo(c *CF) StackProvider {
	return &CF{
		c.Context,
		c.Service,
	}
}

func (c *CF) exists(stack_name *string) (bool) {
	ds := c.describe(stack_name)
	if len(ds.Stacks) > 0 {
		return true
	}
	return false
}

func (c *CF) rollback(stack_name *string) (bool) {
	ds := c.describe(stack_name)
	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}
	return false
}

func (c *CF) describe(stack_name *string) (*cloudformation.DescribeStacksOutput) {
	input := cloudformation.DescribeStacksInput{StackName: stack_name}
	output, _ := c.Service.DescribeStacksWithContext(c.Context, &input)
	return output
}
