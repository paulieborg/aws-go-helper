package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"time"
)

func AwsWaitDelete(ctx aws.Context, svc *cf.CloudFormation, input cf.DescribeStacksInput) {

	svc.WaitUntilStackDeleteCompleteWithContext(
		ctx,
		&input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
	)

	return
}
