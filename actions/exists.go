package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func exists(ctx aws.Context, svc *cf.CloudFormation, name string) (e bool) {

	ds, _ := Describe(ctx, svc, cf.DescribeStacksInput{StackName: &name})

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}
