package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
	cfiface "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

const capability string = "CAPABILITY_NAMED_IAM"

// Service is ...
type Service struct {
	Context aws.Context
	CFAPI   cfiface.CloudFormationAPI
}

// Config represents a stack
type Config struct {
	StackName  string
	Parameters []*cfapi.Parameter
	Template   []byte
	BucketName string
	Timeout    int64
}

// ControlProvider is ...
type ControlProvider interface {
	Create(*Config) (*cfapi.CreateStackOutput, error)
	Update(*Config) (*cfapi.UpdateStackOutput, error)
	Delete(*string) (*cfapi.DeleteStackOutput, error)
}

// Controller is ...
func Controller(svc *Service) ControlProvider {
	return &Service{
		svc.Context,
		svc.CFAPI,
	}
}
