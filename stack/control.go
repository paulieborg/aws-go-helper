package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/paulieborg/aws-go-helper/s3"
)

var capability string = "CAPABILITY_NAMED_IAM"

type StackService struct {
	Context aws.Context
	Service cfapi.CloudFormationAPI
}

type StackConfig struct {
	Stack_name string
	Parameters []*cf.Parameter
	Template   []byte
	BucketName string
	Timeout    int64
}

type StackControlProvider interface {
	Create(*StackConfig) (*cf.CreateStackOutput, error)
	Update(*StackConfig) (*cf.UpdateStackOutput, error)
	Delete(*string) (*cf.DeleteStackOutput, error)
}

func StackController(ss *StackService) StackControlProvider {

	service := StackService{
		ss.Context,
		ss.Service,
	}

	return &service
}

func (s *StackService) Create(p *StackConfig) (*cf.CreateStackOutput, error) {

	stackInput := &cf.CreateStackInput{
		StackName: aws.String(p.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       p.Parameters,
		TimeoutInMinutes: aws.Int64(p.Timeout),
	}

	if p.BucketName == "" {
		stackInput.TemplateBody = aws.String(string(p.Template))
	} else {

		b := s3.BucketFactory{
			Stack_name: p.Stack_name,
			Template:   p.Template,
			BucketName: p.BucketName,
		}

		stackInput.TemplateURL = aws.String(s3.S3upload(b))
	}

	return s.Service.CreateStackWithContext(s.Context, stackInput)

}

func (s *StackService) Update(p *StackConfig) (*cf.UpdateStackOutput, error) {

	stack := &cf.UpdateStackInput{
		StackName: aws.String(p.Stack_name),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: p.Parameters,
	}

	if p.BucketName == "" {
		stack.TemplateBody = aws.String(string(p.Template))
	} else {

		b := s3.BucketFactory{
			Stack_name: p.Stack_name,
			Template:   p.Template,
			BucketName: p.BucketName,
		}

		stack.TemplateURL = aws.String(s3.S3upload(b))
	}

	return s.Service.UpdateStackWithContext(s.Context, stack)

}

func (s *StackService) Delete(stack_name *string) (*cf.DeleteStackOutput, error) {

	input := cf.DeleteStackInput{StackName: stack_name}
	return s.Service.DeleteStackWithContext(s.Context, &input)

}
