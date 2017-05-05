package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (s *Stack) waitCreate(p ProvisionArgs) {

	sess := cf.New(session.Must(session.NewSession()))
	input := cf.DescribeStacksInput{StackName: &p.Stack_name}

	sess.WaitUntilStackCreateCompleteWithContext(
		s.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
