package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"bytes"
	"os"
)

func (p_args *ProvisionArgs) s3upload() (path string, err error) {

	s := session.Must(session.NewSession())
	svc := s3.New(s)

	url := "https://s3-" + os.Getenv("AWS_REGION") + ".amazonaws.com/"
	file_path := "/cloudformation-templates/" + p_args.Stack_name
	path = url + p_args.Bucket + file_path

	params := &s3.PutObjectInput{
		Body:   bytes.NewReader([]byte(p_args.Template)),
		Bucket: &p_args.Bucket,
		Key:    aws.String(file_path),
	}

	_, err = svc.PutObject(params)

	return
}
