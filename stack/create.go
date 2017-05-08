package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/paulieborg/aws-go-helper/s3"
)

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
