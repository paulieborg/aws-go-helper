package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cf    "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/paulieborg/aws-go-helper/s3"
)

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
