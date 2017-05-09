package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Create does ...
func (svc *Service) Create(cfg *Config) (*cfapi.CreateStackOutput, error) {

	var err error = nil

	c := &cfapi.CreateStackInput{
		StackName: aws.String(cfg.StackName),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       cfg.Parameters,
		TimeoutInMinutes: aws.Int64(cfg.Timeout),
	}

	if cfg.BucketName == "" {
		c.TemplateBody = aws.String(string(cfg.Template))
	} else {
		b := CFBucket{
			StackName:  cfg.StackName,
			Template:   cfg.Template,
			BucketName: cfg.BucketName,
		}
		c.TemplateURL, err = svc.Upload(b)

		if err != nil {
			return nil, err
		}
	}

	return svc.CFAPI.CreateStackWithContext(svc.Context, c)

}
