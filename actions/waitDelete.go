package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func (s *Stack) waitDelete(stack_name *string) {

	sess := cf.New(session.Must(session.NewSession()))
	input := cf.DescribeStacksInput{StackName: stack_name}

	sess.WaitUntilStackDeleteCompleteWithContext(
		s.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
