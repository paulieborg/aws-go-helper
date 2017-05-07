package stack

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"testing"
	"context"
	//"fmt"
	"errors"
)

var (
	update_stack_name   = "TestStack"
	update_stack_status = "CREATE_COMPLETE"
	update_stack_id     = "Update Stack ID is 1234567890"

	TestUpdateStack cf.Stack = cf.Stack{
		StackName:   &update_stack_name,
		StackStatus: &update_stack_status,
	}

	update_param_key   = "TestParamKey"
	update_param_value = "TestParamValue"
	//bucket_name         = "TestBucket"
	update_stack_timeout int64 = 1
	update_parameters    []*cf.Parameter
)

func TestUpdate(t *testing.T) {
	//when

	c := Controller(NewMockUpdateSVC())

	update_parameters = append(update_parameters, &cf.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &update_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		//BucketName: *bucket,
		Timeout: *timeout,
	}

	//then
	update_output, _ := c.Update(&config)

	if update_output.StackId != &update_stack_id {
		t.Errorf("expected Update response to be 1234567890, got (%s)", *update_output.StackId)
	}

}

func TestUpdateWithErr(t *testing.T) {
	//when

	testError := errors.New("bad-info-error")
	c := Controller(NewErrorMockUpdateSVC(testError))

	update_parameters = append(update_parameters, &cf.Parameter{
		ParameterKey:   &update_param_key,
		ParameterValue: &update_param_value,
	})

	params := &update_parameters

	config := Config{
		StackName:  *name,
		Parameters: *params,
		Template:   template,
		//BucketName: *bucket,
		Timeout: *timeout,
	}

	//then

	response, updateErr := c.Update(&config)

	if (cf.UpdateStackOutput{}) != *response {
		t.Errorf("expected Update to return %s, got %s.", cf.UpdateStackOutput{}, *response)
	}

	if updateErr != testError {
		t.Errorf("expected error, got %v.", updateErr)
	}

}

// Helper Methods
type mockUpdateSVC struct {
	cfapi.CloudFormationAPI
	input    *cf.UpdateStackInput
	output   *cf.UpdateStackOutput
	err      error
	isCalled bool
}

func (m *mockUpdateSVC) UpdateStackWithContext(ctx aws.Context, input *cf.UpdateStackInput, opts ...request.Option) (*cf.UpdateStackOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockUpdateSVC() *Service {

	var stacks []*cf.Stack
	stacks = append(stacks, &TestUpdateStack)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockUpdateSVC{
			output: &cf.UpdateStackOutput{
				StackId: &update_stack_id,
			},
		},
	}
}

func NewErrorMockUpdateSVC(err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockUpdateSVC{
			output: &cf.UpdateStackOutput{},
			err:    err,
		},
	}

}
