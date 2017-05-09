package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Update does  ...
func (svc *Service) Update(cfg *Config) (*cfapi.UpdateStackOutput, error) {

	var err error = nil

	u := &cfapi.UpdateStackInput{
		StackName: aws.String(cfg.StackName),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters: cfg.Parameters,
	}

	if cfg.BucketName == "" {
		u.TemplateBody = aws.String(string(cfg.Template))
	} else {
		b := CFBucket{
			StackName:  cfg.StackName,
			Template:   cfg.Template,
			BucketName: cfg.BucketName,
		}

		u.TemplateURL, err = svc.Upload(b)

		if err != nil {
			return nil, err
		}
	}

	return svc.CFAPI.UpdateStackWithContext(svc.Context, u)

}
