package stack

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CFBucket is an S3 bucket which contains a CloudFormation template.
type CFBucket struct {
	StackName  string
	Template   []byte
	BucketName string
	URL        string
}

// Upload sends a CFBucket to S3 and returns its URL.
func (svc *Service) Upload(b CFBucket) (*string, error) {
	region := os.Getenv("AWS_REGION") // TODO brittle
	u := fmt.Sprintf("https://s3-%s.amazonaws.com/%s/cloudformation-templates/%s", region, b.BucketName, b.StackName)

	b.URL = u
	err := svc.uploadTemplate(b)

	if err != nil {
		return nil, err
	}

	return &u, err
}

func (svc *Service) uploadTemplate(b CFBucket) error {
	p := &s3.PutObjectInput{
		Body:   bytes.NewReader(b.Template),
		Bucket: &b.BucketName,
		Key:    aws.String(b.URL),
	}

	_, err := svc.S3API.PutObject(p)

	return err
}
