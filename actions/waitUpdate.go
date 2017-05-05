package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (s *Stack) waitUpdate(p ProvisionArgs) {

	sess := cf.New(session.Must(session.NewSession()))
	filter := cf.DescribeStacksInput{StackName: &p.Stack_name}

	sess.WaitUntilStackUpdateCompleteWithContext(
		s.Context,
		&filter,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
