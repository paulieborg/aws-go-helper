package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (c *Context) Describe(p ProvisionArgs) (*cf.DescribeStacksOutput) {

	sess := cf.New(session.Must(session.NewSession()))

	input := cf.DescribeStacksInput{StackName: &p.Stack_name}
	d, _ := sess.DescribeStacksWithContext(c.Context, &input)
	return d
}
