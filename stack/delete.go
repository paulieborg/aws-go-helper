package stack

import (
	cf    "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Delete does ...
func (svc *Service) Delete(n *string) (*cf.DeleteStackOutput, error) {
	d := cf.DeleteStackInput{StackName: n}
	return svc.CFAPI.DeleteStackWithContext(svc.Context, &d)
}