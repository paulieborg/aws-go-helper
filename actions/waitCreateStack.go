package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func AwsWaitCreateStack(
	ctx aws.Context,
	svc *cf.CloudFormation,
	stackInfo cf.DescribeStacksInput) {

	svc.WaitUntilStackCreateCompleteWithContext(
		ctx,
		&stackInfo,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
	)

	return
}
