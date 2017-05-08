package stack

import (
	"context"
	"testing"

	cfapi "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

func TestControl(t *testing.T) {
	//when

	controller := Controller(NewMockControlSVC())

	//then
	_, ok := controller.(*Service)
	if !ok {
		t.Errorf("expected Controller to be (*stack.Service), got (%T)", controller)
	}
}

//// Helper Methods
type mockControlSVC struct {
	cfapi.CloudFormationAPI
}

func NewMockControlSVC() *Service {

	return &Service{
		Context: context.Background(),
		CFAPI:   &mockControlSVC{},
	}
}
