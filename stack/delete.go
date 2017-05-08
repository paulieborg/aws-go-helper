package stack

import (
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Delete does ...
func (svc *Service) Delete(n *string) (*cfapi.DeleteStackOutput, error) {
	d := cfapi.DeleteStackInput{StackName: n}
	return svc.CFAPI.DeleteStackWithContext(svc.Context, &d)
}
