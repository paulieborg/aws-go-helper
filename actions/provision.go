package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

var capability string = "CAPABILITY_NAMED_IAM"

type Context struct {
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
func (c *Context) Provision(p ProvisionArgs) (string) {

	if c.exists(p) && c.rollback(p) {
		c.Delete(&p.Stack_name)
		c.create(p)
	} else if c.exists(p) {
		c.update(p)
	} else {
		c.create(p)
	}

	return aws.StringValue(c.Describe(p).Stacks[0].StackStatus)
}
