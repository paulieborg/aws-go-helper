package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"testing"
	"context"
	"errors"
)

var (
	create_stack_name   = "TestStack"
	create_stack_status = "CREATE_COMPLETE"
	create_stack_id     = "Create Stack ID is 1234567890"

	TestCreateStack cf.Stack = cf.Stack{
		StackName:   &create_stack_name,
		StackStatus: &create_stack_status,
	}

	create_param_key   = "TestParamKey"
	create_param_value = "TestParamValue"
	//bucket_name         = "TestBucket"
	create_stack_timeout int64 = 1
	create_parameters    []*cf.Parameter

	name     = &create_stack_name
	template = []byte{'T', 'e', 's', 't', 'i', 'n', 'g'}
	//bucket := &bucket_name
	timeout = &create_stack_timeout
)

func TestCreate(t *testing.T) {
	//when

	c := Controller(NewMockCreateSVC())

	create_parameters = append(create_parameters, &cf.Parameter{
		ParameterKey:   &create_param_key,
		ParameterValue: &create_param_value,
	})

	params := &create_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		//BucketName: *bucket,
		Timeout: *timeout,
	}

	//then
	create_output, _ := c.Create(&config)

	if create_output.StackId != &create_stack_id {
		t.Errorf("expected Create response to be 1234567890, got (%s)", *create_output.StackId)
	}

}

func TestCreateWithErr(t *testing.T) {
	//when

	testError := errors.New("bad-info-error")
	c := Controller(NewErrorMockCreateSVC(testError))

	create_parameters = append(create_parameters, &cf.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &create_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		//BucketName: *bucket,
		Timeout: *timeout,
	}

	//then

	response, createErr := c.Create(&config)

	if (cf.CreateStackOutput{}) != *response {
		t.Errorf("expected Create to return %T, got %T.", cf.CreateStackOutput{}, *response)
	}

	if createErr != testError {
		t.Errorf("expected error, got %v.", createErr)
	}

}

// Helper Methods
type mockCreateSVC struct {
	cfapi.CloudFormationAPI
	input    *cf.CreateStackInput
	output   *cf.CreateStackOutput
	err      error
	isCalled bool
}

func (m *mockCreateSVC) CreateStackWithContext(ctx aws.Context, input *cf.CreateStackInput, opts ...request.Option) (*cf.CreateStackOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockCreateSVC() *Service {

	var stacks []*cf.Stack
	stacks = append(stacks, &TestCreateStack)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockCreateSVC{
			output: &cf.CreateStackOutput{
				StackId: &create_stack_id,
			},
		},
	}
}

func NewErrorMockCreateSVC(err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockCreateSVC{
			output: &cf.CreateStackOutput{},
			err:    err,
		},
	}

}
