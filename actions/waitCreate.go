package actions

import (
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

func waitCreate(stack *StackArgs) error {
	input := cf.DescribeStacksInput{StackName: &stack.Stack_name}

	return stack.Session.WaitUntilStackCreateCompleteWithContext(
		stack.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
}
