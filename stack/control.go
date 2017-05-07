package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cf    "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"

	"github.com/paulieborg/aws-go-helper/s3"
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

// Create does ...
func (svc *Service) Create(cfg *Config) (*cf.CreateStackOutput, error) {
	si := &cf.CreateStackInput{
		StackName: aws.String(cfg.StackName),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       cfg.Parameters,
		TimeoutInMinutes: aws.Int64(cfg.Timeout),
	}

	if cfg.BucketName == "" {
		si.TemplateBody = aws.String(string(cfg.Template))
	} else {
		b := s3.CFBucket{
			StackName:  cfg.StackName,
			Template:   cfg.Template,
			BucketName: cfg.BucketName,
		}
		si.TemplateURL = aws.String(s3.Upload(b))
	}

	return svc.CFAPI.CreateStackWithContext(svc.Context, si)

}

// Update does  ...
func (svc *Service) Update(cfg *Config) (*cf.UpdateStackOutput, error) {
	si := &cf.UpdateStackInput{
		StackName: aws.String(cfg.StackName),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: cfg.Parameters,
	}

	if cfg.BucketName == "" {
		si.TemplateBody = aws.String(string(cfg.Template))
	} else {
		b := s3.CFBucket{
			StackName:  cfg.StackName,
			Template:   cfg.Template,
			BucketName: cfg.BucketName,
		}
		si.TemplateURL = aws.String(s3.Upload(b))
	}

	return svc.CFAPI.UpdateStackWithContext(svc.Context, si)

}

// Delete does ...
func (svc *Service) Delete(n *string) (*cf.DeleteStackOutput, error) {
	si := cf.DeleteStackInput{StackName: n}
	return svc.CFAPI.DeleteStackWithContext(svc.Context, &si)

}
