package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (c *Context) waitUpdate(p ProvisionArgs) {

	sess := cf.New(session.Must(session.NewSession()))
	filter := cf.DescribeStacksInput{StackName: &p.Stack_name}

	sess.WaitUntilStackUpdateCompleteWithContext(
		c.Context,
		&filter,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
