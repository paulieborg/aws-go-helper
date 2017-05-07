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
	stack_name   = "TestStack"
	stack_status = "CREATE_COMPLETE"

	TestStack cf.Stack = cf.Stack{
		StackName:   &stack_name,
		StackStatus: &stack_status,
	}
)

func TestStackInfo(t *testing.T) {
	//when

	si := StackInfo(NewMockInfoSVC())

	//then
	_, ok := si.(*Service)
	if !ok {
		t.Errorf("expected Stack Waiter to be (*stack.Service), got (%T)", si)
	}

	exists, _ := si.Exists(&StackName)

	if !exists {
		t.Errorf("expected Exists to return true got %s", exists)
	}

	rollback, _ := si.Rollback(&StackName)
	if rollback {
		t.Fatal("expected Rollback to return false got %s.", rollback)
	}

	response, _ := si.Describe(&StackName)

	if response.Stacks[0].StackName != &stack_name {
		t.Fatal("expected StackName to be be (%v), got (%v).", stack_name, *response.Stacks[0].StackName)
	}

	if response.Stacks[0].StackStatus != &stack_status {
		t.Fatal("expected StackName to be be (%v), got (%v).", stack_status, *response.Stacks[0].StackStatus)
	}

}

func TestStackInfoWithErr(t *testing.T) {
	//when

	testError := errors.New("bad-info-error")
	si := StackInfo(NewErrorMockInfoSVC(testError))

	//then

	exists, existsErr := si.Exists(&StackName)

	if exists {
		t.Errorf("expected Exists to return false, got %s.", exists)
	}

	if existsErr != testError {
		t.Errorf("expected error, got %v.", existsErr)
	}

	rollback, rollbackErr := si.Rollback(&StackName)

	if rollback {
		t.Errorf("expected Rollback to return false, got %s.", rollback)
	}

	if rollbackErr != testError {
		t.Errorf("expected error, got %v.", rollbackErr)
	}

	describe, describeErr := si.Describe(&StackName)

	//todo: Check this test as the sense is wrong
	if (&cf.DescribeStacksOutput{}) == describe {
		t.Errorf("expected Describe to return %s, got %s.", &cf.DescribeStacksOutput{}, describe)
	}

	if describeErr != testError {
		t.Errorf("expected error, got %v.", describeErr)
	}

}

// Helper Methods
type mockInfoSVC struct {
	cfapi.CloudFormationAPI
	input    *cf.DescribeStacksInput
	output   *cf.DescribeStacksOutput
	err      error
	isCalled bool
}

func (m *mockInfoSVC) DescribeStacksWithContext(ctx aws.Context, input *cf.DescribeStacksInput, opts ...request.Option) (*cf.DescribeStacksOutput, error) {
	m.isCalled = true
	m.input = input
	return m.output, m.err
}

func NewMockInfoSVC() *Service {

	var stacks []*cf.Stack
	stacks = append(stacks, &TestStack)

	return &Service{
		Context: context.Background(),
		CFAPI: &mockInfoSVC{
			output: &cf.DescribeStacksOutput{
				Stacks: stacks,
			},
		},
	}
}

func NewErrorMockInfoSVC(err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI: &mockInfoSVC{
			output: &cf.DescribeStacksOutput{},
			err:    err,
		},
	}

}