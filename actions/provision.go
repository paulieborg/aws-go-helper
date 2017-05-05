package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

var capability string = "CAPABILITY_NAMED_IAM"

type Stack struct {
	Context aws.Context
}

type ProvisionArgs struct {
	Stack_name string
	Parameters []*cf.Parameter
	Template   []byte
	BucketName string
	Timeout    int64
}

// Provision a CloudFormation stack
func (s *Stack) Provision(p ProvisionArgs) (string) {

	if s.exists(p) && s.rollback(p) {
		s.Delete(&p.Stack_name)
		s.create(p)
	} else if s.exists(p) {
		s.update(p)
	} else {
		s.create(p)
	}

	return aws.StringValue(s.Describe(p).Stacks[0].StackStatus)
}
