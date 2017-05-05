package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (c *Context) Describe(p ProvisionArgs) (*cf.DescribeStacksOutput) {
	input := cf.DescribeStacksInput{StackName: &p.Stack_name}
	d, _ := c.Service.DescribeStacksWithContext(c.Context, &input)
	return d
}
