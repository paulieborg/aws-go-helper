package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/aws"
)

func AwsDeleteStack(ctx aws.Context, svc *cf.CloudFormation, input cf.DeleteStackInput) (d *cf.DeleteStackOutput, err error) {
	return svc.DeleteStackWithContext(ctx, &input)
}
