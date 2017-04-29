package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func Delete(p_args ProvisionArgs) (d *cf.DeleteStackOutput, err error) {
	input := cf.DeleteStackInput{StackName: &p_args.Stack_name}
	return p_args.Session.DeleteStackWithContext(p_args.Context, &input)
}
