package stack

import (
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
)

type InfoProvider interface {
	Exists(*string) (bool, error)
	Rollback(*string) (bool, error)
	Describe(*string) (*cfapi.DescribeStacksOutput, error)
}

// Info ...
func Info(s *Service) InfoProvider {
	return &Service{
		s.Context,
		s.CFAPI,
		s.S3API,
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

func (s *Service) Describe(n *string) (*cfapi.DescribeStacksOutput, error) {
	in := cfapi.DescribeStacksInput{StackName: n}
	out, err := s.CFAPI.DescribeStacksWithContext(s.Context, &in)
	return out, err
}
