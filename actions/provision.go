package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

var capability string = "CAPABILITY_NAMED_IAM"

type StackArgs struct {
	Context          aws.Context
	Session          *cf.CloudFormation
	Parameters       []*cf.Parameter
	Stack_name       string
	Template         []byte
	TemplateFileName string
	Bucket           string
	Timeout          int64
}

// Provision a CloudFormation stack
func (s *StackArgs) Provision() (err error) {

	if s.exists() && s.rollback() {
		s.Delete()
		s.create()
	} else if s.exists() {
		s.update()
	} else {
		s.create()
	}

	return
}
