package s3

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	s3api "github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Service struct {
	Context aws.Context
	S3API   s3api.S3API
}

// CFBucket is an S3 bucket which contains a CloudFormation template.
type CFBucket struct {
	StackName  string
	Template   []byte
	BucketName string
	URL        string
}

type UploadTemplate interface {
	Upload(svc *Service, b CFBucket) (u string)
}

// Upload sends a CFBucket to S3 and returns its URL.
func Upload(svc *Service, b CFBucket) (u string) {
	region := os.Getenv("AWS_REGION") // TODO brittle
	u = fmt.Sprintf("https://s3-%s.amazonaws.com/%s/cloudformation-templates/%s", region, b.BucketName, b.StackName)

	b.URL = u
	err := uploadTemplate(svc, b)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func uploadTemplate(svc *Service, b CFBucket) (err error) {
	p := &s3.PutObjectInput{
		Body:   bytes.NewReader(b.Template),
		Bucket: &b.BucketName,
		Key:    aws.String(b.URL),
	}

	_, err = svc.S3API.PutObject(p)

	return
}
