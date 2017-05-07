package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cf    "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

const capability string = "CAPABILITY_NAMED_IAM"

// Service is ...
type Service struct {
	Context aws.Context
	CFAPI cfapi.CloudFormationAPI
}

// Config represents a stack
type Config struct {
	StackName  string
	Parameters []*cf.Parameter
	Template   []byte
	BucketName string
	Timeout    int64
}

// ControlProvider is ...
type ControlProvider interface {
	Create(*Config) (*cf.CreateStackOutput, error)
	Update(*Config) (*cf.UpdateStackOutput, error)
	Delete(*string) (*cf.DeleteStackOutput, error)
}

// Controller is ...
func Controller(svc *Service) ControlProvider {
	return &Service{
		svc.Context,
		svc.CFAPI,
	}
}
