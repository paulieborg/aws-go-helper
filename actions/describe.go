package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (s *StackArgs) Describe() (*cf.DescribeStacksOutput) {
	input := cf.DescribeStacksInput{StackName: &s.Stack_name}
	d, _ := s.Session.DescribeStacksWithContext(s.Context, &input)
	return d
}
