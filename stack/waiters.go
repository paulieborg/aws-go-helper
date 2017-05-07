package stack

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

type WaiterProvider interface {
	WaitCreate(*string)
	WaitDelete(*string)
	WaitUpdate(*string)
}

func Waiter(s *Service) WaiterProvider {
	return &Service{
		s.Context,
		s.CFAPI,
	}
}

func (s *Service) WaitCreate(n *string) {
	in := cf.DescribeStacksInput{StackName: n}

	s.CFAPI.WaitUntilStackCreateCompleteWithContext(
		s.Context,
		&in,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15 * time.Second)),
	)
}

func (s *Service) WaitDelete(n *string) {
	in := cf.DescribeStacksInput{StackName: n}

	s.CFAPI.WaitUntilStackDeleteCompleteWithContext(
		s.Context,
		&in,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15 * time.Second)),
	)
}

func (s *Service) WaitUpdate(n *string) {
	flt := cf.DescribeStacksInput{StackName: n}

	s.CFAPI.WaitUntilStackUpdateCompleteWithContext(
		s.Context,
		&flt,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15 * time.Second)),
	)
}
