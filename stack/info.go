package stack

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type StackInfoProvider interface {
	Exists(*string) (bool, error)
	Rollback(*string) (bool, error)
	Describe(*string) (*cloudformation.DescribeStacksOutput, error)
}

func StackInfo(ss *Service) StackInfoProvider {
	return &Service{
		ss.Context,
		ss.CFAPI,
	}
}

func (s *Service) Exists(n *string) (bool, error) {
	ds, err := s.Describe(n)

	if err != nil {
		return false, err
	}

	return len(ds.Stacks) > 0, err
}

func (s *Service) Rollback(n *string) (bool, error) {
	ds, err := s.Describe(n)

	if err != nil {
		return false, err
	}

	return *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE", err
}

func (s *Service) Describe(n *string) (*cloudformation.DescribeStacksOutput, error) {
	in := cloudformation.DescribeStacksInput{StackName: n}
	out, err := s.CFAPI.DescribeStacksWithContext(s.Context, &in)
	return out, err
}
