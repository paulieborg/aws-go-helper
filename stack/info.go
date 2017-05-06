package stack

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type StackInfoProvider interface {
	Exists(*string) (bool)
	Rollback(*string) (bool)
	Describe(*string) (*cloudformation.DescribeStacksOutput)
}

func StackInfo(ss *StackService) StackInfoProvider {
	return &StackService{
		ss.Context,
		ss.Service,
	}
}

func (s *StackService) Exists(stack_name *string) (bool) {
	ds := s.Describe(stack_name)
	if len(ds.Stacks) > 0 {
		return true
	}
	return false
}

func (s *StackService) Rollback(stack_name *string) (bool) {
	ds := s.Describe(stack_name)
	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}
	return false
}

func (s *StackService) Describe(stack_name *string) (*cloudformation.DescribeStacksOutput) {
	input := cloudformation.DescribeStacksInput{StackName: stack_name}
	output, _ := s.Service.DescribeStacksWithContext(s.Context, &input)
	return output
}
