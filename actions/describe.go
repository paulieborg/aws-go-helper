package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (s *Stack) Describe(p ProvisionArgs) (*cf.DescribeStacksOutput) {

	sess := cf.New(session.Must(session.NewSession()))

	input := cf.DescribeStacksInput{StackName: &p.Stack_name}
	d, _ := sess.DescribeStacksWithContext(s.Context, &input)
	return d
}
