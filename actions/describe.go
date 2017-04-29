package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func Describe(ctx aws.Context, svc *cf.CloudFormation, input cf.DescribeStacksInput) (d *cf.DescribeStacksOutput, err error) {
	return svc.DescribeStacksWithContext(ctx, &input)
}
