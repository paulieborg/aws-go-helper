package actions

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (s *StackArgs) waitCreate() {
	input := cf.DescribeStacksInput{StackName: &s.Stack_name}

	s.Session.WaitUntilStackCreateCompleteWithContext(
		s.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
