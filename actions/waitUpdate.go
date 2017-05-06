package actions

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (c *CF) waitUpdate(p ProvisionArgs) {

	filter := cf.DescribeStacksInput{StackName: &p.Stack_name}

	c.Service.WaitUntilStackUpdateCompleteWithContext(
		c.Context,
		&filter,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
	return
}
