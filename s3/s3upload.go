package s3

import (
	"bytes"
	"fmt"
	"os"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// CFBucket is an S3 bucket which contains a CloudFormation template.
type CFBucket struct {
	StackName  string
	Template   []byte
	BucketName string
	URL        string
}

// Upload sends a CFBucket to S3 and returns its URL.
func Upload(b CFBucket) (u string) {
	svc := s3.New(session.Must(session.NewSession()))
	region := os.Getenv("AWS_REGION") // TODO brittle
	u = fmt.Sprintf("https://s3-%s.amazonaws.com/%s/cloudformation-templates/%s", region, b.BucketName, b.StackName)

	b.URL = u
	err := uploadTemplate(svc, b)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func uploadTemplate(svc s3iface.S3API, b CFBucket) (err error) {
	p := &s3.PutObjectInput{
		Body:   bytes.NewReader(b.Template),
		Bucket: &b.BucketName,
		Key:    aws.String(b.URL),
	}

	_, err = svc.PutObject(p)

	return
}
