package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func waitDelete(stack *StackArgs) error {
	input := cf.DescribeStacksInput{StackName: &stack.Stack_name}

	return stack.Session.WaitUntilStackDeleteCompleteWithContext(
		stack.Context,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(15*time.Second)),
	)
}
