package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func waitUpdate(
	ctx aws.Context,
	svc *cf.CloudFormation,
	stackInfo cf.DescribeStacksInput) {

	svc.WaitUntilStackUpdateCompleteWithContext(
		ctx,
		&stackInfo,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
	)

	return
}
