package stack

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
	cfiface "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"testing"
)

var (
	create_stack_id = "Create Stack ID is 1234567890"

	TestCreateStack cfapi.Stack = cfapi.Stack{
		StackName:   &stack_name,
		StackStatus: &stack_status,
	}

	create_param_key           = "TestParamKey"
	create_param_value         = "TestParamValue"
	create_bucket_name         = "TestBucket"
	create_stack_timeout int64 = 1
	create_parameters    []*cfapi.Parameter

	name     = &stack_name
	template = []byte{'T', 'e', 's', 't', 'i', 'n', 'g'}
	bucket   = &create_bucket_name
	timeout  = &create_stack_timeout
)

func TestCreate(t *testing.T) {
	//when

	c := Controller(NewMockCreateSVC())

	create_parameters = append(create_parameters, &cfapi.Parameter{
		ParameterKey:   &create_param_key,
		ParameterValue: &create_param_value,
	})

	params := &create_parameters

	config_no_bucket := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		Timeout:    *timeout,
	}

	//then
	output_no_bucket, _ := c.Create(&config_no_bucket)

	if output_no_bucket.StackId != &create_stack_id {
		t.Errorf("expected Create response to be 1234567890, got (%s)", *output_no_bucket.StackId)
	}

	config_with_bucket := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		BucketName: *bucket,
		Timeout:    *timeout,
	}

	//then
	output_with_bucket, _ := c.Create(&config_with_bucket)

	if output_with_bucket.StackId != &create_stack_id {
		t.Errorf("expected Create response to be 1234567890, got (%s)", *output_with_bucket.StackId)
	}

}

func TestCreateWithErr(t *testing.T) {
	//when

	testCreateError := errors.New("bad-create-error")
	c := Controller(NewErrorMockCreateSVC(testCreateError, nil))

	create_parameters = append(create_parameters, &cfapi.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &create_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		BucketName: *bucket,
		Timeout:    *timeout,
	}

	//then

	response, createErr := c.Create(&config)

	if (cfapi.CreateStackOutput{}) != *response {
		t.Errorf("expected Create to return %T, got %T.", cfapi.CreateStackOutput{}, *response)
	}

	if createErr != testCreateError {
		t.Errorf("expected error, got %v.", createErr)
	}

}

func TestCreateUploadWithErr(t *testing.T) {
	//when

	testUploadError := errors.New("bad-upload-error")

	update_parameters = append(update_parameters, &cfapi.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &update_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		BucketName: *bucket,
		Timeout:    *timeout,
	}

	//then

	svc := NewErrorMockCreateSVC(nil, testUploadError)
	c := Controller(svc)

	_, uplErr := c.Create(&config)

	if uplErr != testUploadError {
		t.Errorf("expected %v, got %v.", testUploadError, uplErr)
	}

}

// Helper Methods
type mockCreateCFSVC struct {
	cfiface.CloudFormationAPI
	input    *cfapi.CreateStackInput
	output   *cfapi.CreateStackOutput
	err      error
	isCalled bool
}

type mockCreateS3SVC struct {
	s3iface.S3API
	input    *s3.PutObjectInput
	output   *s3.PutObjectOutput
	err      error
	isCalled bool
}

func (m *mockCreateCFSVC) CreateStackWithContext(ctx aws.Context, input *cfapi.CreateStackInput, opts ...request.Option) (*cfapi.CreateStackOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func (m *mockCreateS3SVC) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockCreateSVC() *Service {

	var stacks []*cfapi.Stack
	stacks = append(stacks, &TestCreateStack)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockCreateCFSVC{
			output: &cfapi.CreateStackOutput{
				StackId: &create_stack_id,
			},
		},
		S3API: &mockCreateS3SVC{
			output: &s3.PutObjectOutput{},
		},
	}
}

func NewErrorMockCreateSVC(cferr error, s3err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockCreateCFSVC{
			output: &cfapi.CreateStackOutput{},
			err:    cferr,
		},
		S3API: &mockCreateS3SVC{
			output: &s3.PutObjectOutput{},
			err:    s3err,
		},
	}

}
