package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
	s3api "github.com/aws/aws-sdk-go/service/s3"
	"github.com/paulieborg/aws-go-helper/s3"
)

// Create does ...
func (svc *Service) Create(cfg *Config) (*cfapi.CreateStackOutput, error) {
	i := &cfapi.CreateStackInput{
		StackName: aws.String(cfg.StackName),
		Capabilities: []*string{
			aws.String(capability),
		},
		Parameters:       cfg.Parameters,
		TimeoutInMinutes: aws.Int64(cfg.Timeout),
	}

	if cfg.BucketName == "" {
		i.TemplateBody = aws.String(string(cfg.Template))
	} else {
		b := s3.CFBucket{
			StackName:  cfg.StackName,
			Template:   cfg.Template,
			BucketName: cfg.BucketName,
		}

		s := s3.Service{
			Context: svc.Context,
			S3API:   s3api.New(session.Must(session.NewSession())),
		}

		i.TemplateURL = aws.String(s3.Upload(&s, b))
	}

	return svc.CFAPI.CreateStackWithContext(svc.Context, i)

}
