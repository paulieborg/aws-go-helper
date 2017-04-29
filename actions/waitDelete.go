package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func (p_args *ProvisionArgs) WaitDelete() {

	input := cf.DescribeStacksInput{StackName: &p_args.Stack_name}

	p_args.Session.WaitUntilStackDeleteCompleteWithContext(
		p_args.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)

	return
}
