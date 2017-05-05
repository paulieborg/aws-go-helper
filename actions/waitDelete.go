package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func (c *Context) waitDelete(stack_name *string) {

	input := cf.DescribeStacksInput{StackName: stack_name}

	c.Service.WaitUntilStackDeleteCompleteWithContext(
		c.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
