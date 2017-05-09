package stack

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"testing"
)

var (
	delete_name = &stack_name
)

func TestDelete(t *testing.T) {
	//when

	c := Controller(NewMockDeleteSVC())

	//then
	response, _ := c.Delete(delete_name)

	if (cf.DeleteStackOutput{}) != *response {
		t.Errorf("expected Delete to return %T, got %T.", cf.DeleteStackOutput{}, *response)
	}

}

func TestDeleteWithErr(t *testing.T) {
	//when

	testError := errors.New("bad-delete-error")
	c := Controller(NewErrorMockDeleteSVC(testError))

	//then

	response, deleteErr := c.Delete(delete_name)

	if (cf.DeleteStackOutput{}) != *response {
		t.Errorf("expected Delete to return %T, got %T.", cf.DeleteStackOutput{}, *response)
	}

	if deleteErr != testError {
		t.Errorf("expected error, got %v.", deleteErr)
	}

}

// Helper Methods
type mockDeleteSVC struct {
	cfapi.CloudFormationAPI
	input    *cf.DeleteStackInput
	output   *cf.DeleteStackOutput
	err      error
	isCalled bool
}

func (m *mockDeleteSVC) DeleteStackWithContext(ctx aws.Context, input *cf.DeleteStackInput, opts ...request.Option) (*cf.DeleteStackOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockDeleteSVC() *Service {

	var stacks []*cf.Stack
	stacks = append(stacks, &TestCreateStack)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockDeleteSVC{
			output: &cf.DeleteStackOutput{},
		},
	}
}

func NewErrorMockDeleteSVC(err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockDeleteSVC{
			output: &cf.DeleteStackOutput{},
			err:    err,
		},
	}

}
