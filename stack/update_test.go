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
	update_stack_id = "Update Stack ID is 1234567890"

	TestUpdateStack cfapi.Stack = cfapi.Stack{
		StackName:   &stack_name,
		StackStatus: &stack_status,
	}

	update_param_key   = "TestParamKey"
	update_param_value = "TestParamValue"
	update_parameters  []*cfapi.Parameter
)

func TestUpdate(t *testing.T) {
	//when

	svc := NewMockUpdateSVC()
	c := Controller(&svc)

	update_parameters = append(update_parameters, &cfapi.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &update_parameters

	config_no_bucket := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		Timeout:    *timeout,
	}

	//then
	output_no_bucket, aerr := c.Update(&config_no_bucket)

	if output_no_bucket.StackId != &update_stack_id {
		t.Errorf("expected Update response to be 1234567890, got (%s)", *output_no_bucket.StackId)
	}

	if aerr != nil {
		t.Errorf("expected error to be nil, got (%v)", aerr)
	}

	config_with_bucket := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		BucketName: *bucket,
		Timeout:    *timeout,
	}

	//then
	output_with_bucket, berr := c.Update(&config_with_bucket)

	if output_with_bucket.StackId != &update_stack_id {
		t.Errorf("expected Update response to be 1234567890, got (%s)", *output_with_bucket.StackId)
	}

	if berr != nil {
		t.Errorf("expected error to be nil, got (%v)", berr)
	}

}

func TestUpdateWithErr(t *testing.T) {
	//when

	testUpdateError := errors.New("bad-update-error")

	update_parameters = append(update_parameters, &cfapi.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &update_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		Timeout:    *timeout,
	}

	//then

	svc := NewMockUpdateErrSVC(testUpdateError, nil)
	c := Controller(&svc)

	update, updateErr := c.Update(&config)

	if *update != (cfapi.UpdateStackOutput{}) {
		t.Errorf("expected Update to return %s, got %s.", cfapi.UpdateStackOutput{}, *update)
	}

	if updateErr != testUpdateError {
		t.Errorf("expected error, got %v.", updateErr)
	}

}

func TestUpdateUploadWithErr(t *testing.T) {
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

	svc := NewMockUpdateErrSVC(nil, testUploadError)
	c := Controller(&svc)

	_, uplErr := c.Update(&config)

	if uplErr != testUploadError {
		t.Errorf("expected %v, got %v.", testUploadError, uplErr)
	}

}

// Helper Methods
type mockUpdateCFSVC struct {
	cfiface.CloudFormationAPI
	input    *cfapi.UpdateStackInput
	output   *cfapi.UpdateStackOutput
	err      error
	isCalled bool
}

type mockUpdateS3SVC struct {
	s3iface.S3API
	input    *s3.PutObjectInput
	output   *s3.PutObjectOutput
	err      error
	isCalled bool
}

func (m *mockUpdateCFSVC) UpdateStackWithContext(ctx aws.Context, input *cfapi.UpdateStackInput, opts ...request.Option) (*cfapi.UpdateStackOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func (m *mockUpdateS3SVC) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockUpdateSVC() Service {

	return Service{
		Context: context.Background(),
		CFAPI: &mockUpdateCFSVC{
			output: &cfapi.UpdateStackOutput{StackId: &update_stack_id},
		},
		S3API: &mockUpdateS3SVC{
			output: &s3.PutObjectOutput{},
		},
	}
}

func NewMockUpdateErrSVC(cferr error, s3err error) Service {

	return Service{
		Context: context.Background(),
		CFAPI: &mockUpdateCFSVC{
			output: &cfapi.UpdateStackOutput{},
			err:    cferr,
		},
		S3API: &mockUpdateS3SVC{
			output: &s3.PutObjectOutput{},
			err:    s3err,
		},
	}
}
