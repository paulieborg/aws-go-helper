package actions

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (s *StackArgs) waitUpdate() {

	filter := cf.DescribeStacksInput{StackName: &s.Stack_name}

	s.Session.WaitUntilStackUpdateCompleteWithContext(
		s.Context,
		&filter,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
