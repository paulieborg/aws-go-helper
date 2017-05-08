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

	c := Controller(NewMockUpdateSVC())

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
	output_no_bucket, _ := c.Update(&config_no_bucket)

	if output_no_bucket.StackId != &update_stack_id {
		t.Errorf("expected Update response to be 1234567890, got (%s)", *output_no_bucket.StackId)
	}

	config_with_bucket := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		BucketName: *bucket,
		Timeout:    *timeout,
	}

	//then
	output_with_bucket, _ := c.Update(&config_with_bucket)

	if output_with_bucket.StackId != &update_stack_id {
		t.Errorf("expected Update response to be 1234567890, got (%s)", *output_with_bucket.StackId)
	}

}

func TestUpdateWithErr(t *testing.T) {
	//when

	testUpdateError := errors.New("bad-update-error")
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
		Timeout:    *timeout,
	}

	//then

	upd_c := Controller(NewUpdateErrorMockSVC(testUpdateError))

	upd, updErr := upd_c.Update(&config)

	if (cfapi.UpdateStackOutput{}) != *upd {
		t.Errorf("expected Update to return %s, got %s.", cfapi.UpdateStackOutput{}, *upd)
	}

	if updErr != testUpdateError {
		t.Errorf("expected error, got %v.", updErr)
	}

	upl_c := Controller(NewUpdateErrorMockUploadSVC(testUploadError))
	upl, uplErr := upl_c.Update(&config)

	if (cfapi.UpdateStackOutput{}) != *upl {
		t.Errorf("expected Update to return %s, got %s.", cfapi.UpdateStackOutput{}, *upl)
	}


	//todo: This is wrong - it should not be nil
	if uplErr != nil {
		t.Errorf("expected error, got %v.", uplErr)
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

func NewMockUpdateSVC() *Service {

	var stacks []*cfapi.Stack
	stacks = append(stacks, &TestUpdateStack)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockUpdateCFSVC{
			output: &cfapi.UpdateStackOutput{
				StackId: &update_stack_id,
			},
		},
		S3API: &mockUpdateS3SVC{
			output: &s3.PutObjectOutput{},
		},
	}
}

func NewUpdateErrorMockSVC(cferr error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockUpdateCFSVC{
			output: &cfapi.UpdateStackOutput{},
			err:    cferr,
		},
		S3API: &mockUpdateS3SVC{
			output: &s3.PutObjectOutput{},
		},
	}

}

func NewUpdateErrorMockUploadSVC(s3err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockUpdateCFSVC{
			output: &cfapi.UpdateStackOutput{},
		},
		S3API: &mockUpdateS3SVC{
			output: &s3.PutObjectOutput{},
			err:    s3err,
		},
	}

}
