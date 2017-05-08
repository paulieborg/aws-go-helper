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
	StackName string = "TestName"
)

func TestWaiter(t *testing.T) {
	//when

	w := Waiter(NewMockWaiterSVC())

	//then
	_, ok := w.(*Service)
	if !ok {
		t.Errorf("expected Stack Waiter to be (*stack.Service), got (%T)", w)
	}
}

func TestWaitWithErr(t *testing.T) {
	//when

	testError := errors.New("bad-wait-error")
	w := Waiter(NewErrorMockWaiterSVC(testError))

	//then

	errCreate := w.WaitCreate(&StackName)
	if errCreate != testError {
		t.Errorf("expected error to be (%v), got (%v).", "bad-wait-error", errCreate)
	}

	errDelete := w.WaitDelete(&StackName)
	if errDelete != testError {
		t.Errorf("expected error to be (%v), got (%v).", "bad-wait-error", errDelete)
	}

	errUpdate := w.WaitUpdate(&StackName)
	if errUpdate != testError {
		t.Errorf("expected error to be (%v), got (%v).", "bad-wait-error", errUpdate)
	}

}

// Helper Methods
type mockWaiterSVC struct {
	cfapi.CloudFormationAPI
	err      error
	isCalled bool
}

func (m *mockWaiterSVC) WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *cf.DescribeStacksInput, opts ...request.WaiterOption) error {
	m.isCalled = true
	return m.err
}

func (m *mockWaiterSVC) WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *cf.DescribeStacksInput, opts ...request.WaiterOption) error {
	m.isCalled = true
	return m.err
}

func (m *mockWaiterSVC) WaitUntilStackUpdateCompleteWithContext(ctx aws.Context, input *cf.DescribeStacksInput, opts ...request.WaiterOption) error {
	m.isCalled = true
	return m.err
}

func NewMockWaiterSVC() *Service {

	return &Service{
		Context: context.Background(),
		CFAPI:   &mockWaiterSVC{},
	}
}

func NewErrorMockWaiterSVC(err error) *Service {

	return &Service{
		Context: context.Background(),
		CFAPI:   &mockWaiterSVC{err: err, },
	}

}
