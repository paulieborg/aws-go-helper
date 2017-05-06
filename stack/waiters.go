package stack

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

type StackWaiterProvider interface {
	WaitCreate(*string)
	WaitDelete(*string)
	WaitUpdate(*string)
}

func StackWaiter(ss *StackService) StackWaiterProvider {
	return &StackService{
		ss.Context,
		ss.Service,
	}
}

func (s *StackService) WaitCreate(stack_name *string) {

	input := cf.DescribeStacksInput{StackName: stack_name}

	s.Service.WaitUntilStackCreateCompleteWithContext(
		s.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}

func (s *StackService) WaitDelete(stack_name *string) {

	input := cf.DescribeStacksInput{StackName: stack_name}

	s.Service.WaitUntilStackDeleteCompleteWithContext(
		s.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}

func (s *StackService) WaitUpdate(stack_name *string) {

	filter := cf.DescribeStacksInput{StackName: stack_name}

	s.Service.WaitUntilStackUpdateCompleteWithContext(
		s.Context,
		&filter,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
