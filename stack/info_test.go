package stack

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	cfapi "github.com/aws/aws-sdk-go/service/cloudformation"
	cfiface "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"testing"
)

var (
	Test cfapi.Stack = cfapi.Stack{
		StackName:   &stack_name,
		StackStatus: &stack_status,
	}
)

func TestInfo(t *testing.T) {
	//when

	i := Info(NewMockInfoSVC())

	//then
	_, ok := i.(*Service)
	if !ok {
		t.Errorf("expected Stack Waiter to be (*stack.Service), got (%T)", i)
	}

	exists, _ := i.Exists(&stack_name)

	if !exists {
		t.Errorf("expected Exists to return true got %v", exists)
	}

	rollback, _ := i.Rollback(&stack_name)
	if rollback {
		t.Errorf("expected Rollback to return false got %v.", rollback)
	}

	response, _ := i.Describe(&stack_name)

	if response.Stacks[0].StackName != &stack_name {
		t.Errorf("expected stack_name to be be (%s), got (%s).", stack_name, *response.Stacks[0].StackName)
	}

	if response.Stacks[0].StackStatus != &stack_status {
		t.Errorf("expected stack_name to be be (%s), got (%s).", stack_status, *response.Stacks[0].StackStatus)
	}

}

func TestInfoWithErr(t *testing.T) {
	//when

	testError := errors.New("bad-info-error")
	i := Info(NewErrorMockInfoSVC(testError))

	//then

	exists, existsErr := i.Exists(&stack_name)

	if exists {
		t.Errorf("expected Exists to return false, got %v.", exists)
	}

	if existsErr != testError {
		t.Errorf("expected error, got %v.", existsErr)
	}

	rollback, rollbackErr := i.Rollback(&stack_name)

	if rollback {
		t.Errorf("expected Rollback to return false, got %v.", rollback)
	}

	if rollbackErr != testError {
		t.Errorf("expected error, got %v.", rollbackErr)
	}

	describe, describeErr := i.Describe(&stack_name)

	//todo: Check this test as the sense is wrong
	if (&cfapi.DescribeStacksOutput{}) == describe {
		t.Errorf("expected Describe to return %s, got %s.", &cfapi.DescribeStacksOutput{}, describe)
	}

	if describeErr != testError {
		t.Errorf("expected error, got %v.", describeErr)
	}

}

// Helper Methods
type mockInfoSVC struct {
	cfiface.CloudFormationAPI
	input    *cfapi.DescribeStacksInput
	output   *cfapi.DescribeStacksOutput
	err      error
	isCalled bool
}

func (m *mockInfoSVC) DescribeStacksWithContext(ctx aws.Context, input *cfapi.DescribeStacksInput, opts ...request.Option) (*cfapi.DescribeStacksOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockInfoSVC() *Service {

	var stacks []*cfapi.Stack
	stacks = append(stacks, &Test)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockInfoSVC{
			output: &cfapi.DescribeStacksOutput{
				Stacks: stacks,
			},
		},
	}
}

func NewErrorMockInfoSVC(err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockInfoSVC{
			output: &cfapi.DescribeStacksOutput{},
			err:    err,
		},
	}

}
