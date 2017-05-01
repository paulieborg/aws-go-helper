package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (s *StackArgs) Describe() (*cf.DescribeStacksOutput, error) {
	input := cf.DescribeStacksInput{StackName: &s.Stack_name}
	return s.Session.DescribeStacksWithContext(s.Context, &input)
}
