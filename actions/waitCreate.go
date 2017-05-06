package actions

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (c *CF) waitCreate(p ProvisionArgs) {

	input := cf.DescribeStacksInput{StackName: &p.Stack_name}

	c.Service.WaitUntilStackCreateCompleteWithContext(
		c.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
