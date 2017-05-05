package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (c *Context) waitCreate(p ProvisionArgs) {

	sess := cf.New(session.Must(session.NewSession()))
	input := cf.DescribeStacksInput{StackName: &p.Stack_name}

	sess.WaitUntilStackCreateCompleteWithContext(
		c.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
