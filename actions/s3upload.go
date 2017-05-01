package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"bytes"
	"os"
	"log"
)

func (s *StackArgs) s3upload() (string) {

	svc := s3.New(session.Must(session.NewSession()))

	url := "https://s3-" + os.Getenv("AWS_REGION") + ".amazonaws.com/"
	file_path := "/cloudformation-templates/" + s.Stack_name
	path := url + s.Bucket + file_path

	params := &s3.PutObjectInput{
		Body:   bytes.NewReader([]byte(s.Template)),
		Bucket: &s.Bucket,
		Key:    aws.String(file_path),
	}

	_, err := svc.PutObject(params)

	if err != nil {
		log.Fatal(err)
	}

	return path
}
