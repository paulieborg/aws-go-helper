package stack

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type StackInfoProvider interface {
	Exists(*string) (bool)
	Rollback(*string) (bool)
	Describe(*string) (*cloudformation.DescribeStacksOutput)
}

func StackInfo(ss *Service) StackInfoProvider {
	return &Service{
		ss.Context,
		ss.CFAPI,
	}
}

func (s *Service) Exists(n *string) (bool) {
	ds := s.Describe(n)
    return len(ds.Stacks) > 0
}

func (s *Service) Rollback(n *string) (bool) {
	ds := s.Describe(n)
    return *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE"
}

func (s *Service) Describe(n *string) (*cloudformation.DescribeStacksOutput) {
	in     := cloudformation.DescribeStacksInput{StackName: n}
	out, _ := s.CFAPI.DescribeStacksWithContext(s.Context, &in)
	return out
}
