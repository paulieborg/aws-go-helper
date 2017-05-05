package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func (s *StackArgs) waitDelete() {

	input := cf.DescribeStacksInput{StackName: &s.Stack_name}

	s.Session.WaitUntilStackDeleteCompleteWithContext(
		s.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
