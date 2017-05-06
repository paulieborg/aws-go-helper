package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

var capability string = "CAPABILITY_NAMED_IAM"

type CF struct {
	Context aws.Context
	Service cfapi.CloudFormationAPI
}

type ProvisionArgs struct {
	Stack_name string
	Parameters []*cf.Parameter
	Template   []byte
	BucketName string
	Timeout    int64
}

// Provision a CloudFormation stack
func (c *CF) Provision(p ProvisionArgs) (string) {

	sp := StackInfo(c)

	var status string

	if sp.exists(&p.Stack_name) && sp.rollback(&p.Stack_name) {
		c.Delete(&p.Stack_name)
		status = *c.create(p)
	} else if sp.exists(&p.Stack_name) {
		status = *c.update(p)
	} else {
		status = *c.create(p)
	}

	return status
}
