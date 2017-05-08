package stack

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
)

type WaiterProvider interface {
	WaitCreate(*string) error
	WaitDelete(*string) error
	WaitUpdate(*string) error
}

func Waiter(s *Service) WaiterProvider {
	return &Service{
		s.Context,
		s.CFAPI,
		s.S3API,
	}
}

func (s *Service) WaitCreate(n *string) error {
	in := cfapi.DescribeStacksInput{StackName: n}

	err := s.CFAPI.WaitUntilStackCreateCompleteWithContext(
		s.Context,
		&in,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return err
}

func (s *Service) WaitDelete(n *string) error {
	in := cfapi.DescribeStacksInput{StackName: n}

	err := s.CFAPI.WaitUntilStackDeleteCompleteWithContext(
		s.Context,
		&in,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return err
}

func (s *Service) WaitUpdate(n *string) error {
	flt := cfapi.DescribeStacksInput{StackName: n}

	err := s.CFAPI.WaitUntilStackUpdateCompleteWithContext(
		s.Context,
		&flt,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return err

}
