package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func WaitDelete(p_args ProvisionArgs, ) {

	input := cf.DescribeStacksInput{StackName: &p_args.Stack_name}

	p_args.Session.WaitUntilStackDeleteCompleteWithContext(
		p_args.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
	)

	return
}
