package s3

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"testing"
)

var (
	name        = "TestStackName"
	template    = []byte{'T', 'e', 's', 't', 'i', 'n', 'g'}
	bucket_name = "TestBucketName"
	url         = "https://s3-ap-southeast-2.amazonaws.com/TestBucketName/cloudformation-templates/" + name

	bucket = CFBucket{}
)

func TestUpload(t *testing.T) {
	//when

	upload_bucket := CFBucket{
		StackName:  name,
		BucketName: bucket_name,
	}

	response_url := Upload(NewMockUploadSVC(), upload_bucket)

	if response_url != url {
		t.Errorf("Expected %s, got (%v)", url, response_url)
	}

}

func TestUploadTemplate(t *testing.T) {
	//when
	err := uploadTemplate(NewMockUploadSVC(), bucket)

	if err != nil {
		t.Error("Upload returned", err)
	}
}

func TestUploadTemplateErr(t *testing.T) {
	//when
	testError := errors.New("bad-create-error")

	err := uploadTemplate(NewErrorMockCreateSVC(testError), bucket)

	if err != testError {
		t.Errorf("Expected bad-create-error but got (%v)", testError)
	}
}

// Helper Methods
type mockUploadSVC struct {
	s3iface.S3API
	input    *s3.PutObjectInput
	output   *s3.PutObjectOutput
	err      error
	isCalled bool
}

func (m *mockUploadSVC) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockUploadSVC() *Service {

	return &Service{
		Context: context.Background(),
		S3API: &mockUploadSVC{
			output: &s3.PutObjectOutput{},
		},
	}
}

func NewErrorMockCreateSVC(err error) *Service {
	return &Service{
		Context: context.Background(),
		S3API: &mockUploadSVC{
			output: &s3.PutObjectOutput{},
			err:    err,
		},
	}
}
