package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"bytes"
	"os"
	"log"
)

type BucketFactory struct {
	Stack_name string
	Template   []byte
	BucketName string
}

func S3upload(p BucketFactory) (string) {

	svc := s3.New(session.Must(session.NewSession()))

	url := "https://s3-" + os.Getenv("AWS_REGION") + ".amazonaws.com/"
	file_path := "/cloudformation-templates/" + p.Stack_name
	path := url + p.BucketName + file_path

	UploadTemplate(svc, p.Template, &p.BucketName, path)

	return path
}

func UploadTemplate(svc s3iface.S3API, template []byte, bucketName *string, path string) {

	params := &s3.PutObjectInput{
		Body:   bytes.NewReader([]byte(template)),
		Bucket: bucketName,
		Key:    aws.String(path),
	}

	_, err := svc.PutObject(params)

	if err != nil {
		log.Fatal(err)
	}
}
