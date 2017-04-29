package actions

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func (p_args *ProvisionArgs) waitUpdate() {

	filter := cf.DescribeStacksInput{StackName: &p_args.Stack_name}

	p_args.Session.WaitUntilStackUpdateCompleteWithContext(
		p_args.Context,
		&filter,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
