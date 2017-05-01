package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func (s *StackArgs) Delete() (d *cf.DeleteStackOutput, err error) {
	input := cf.DeleteStackInput{StackName: &s.Stack_name}
	return s.Session.DeleteStackWithContext(s.Context, &input)
}
