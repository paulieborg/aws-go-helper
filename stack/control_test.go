package stack

import (
	"context"
	cfiface "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	s3iface "github.com/aws/aws-sdk-go/service/s3/s3iface"
	"testing"
)

var (
	stack_name   = "TestStack"
	stack_status = "CREATE_COMPLETE"
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
type mockCFAPI struct {
	cfiface.CloudFormationAPI
}

type mockS3API struct {
	s3iface.S3API
}

func NewMockControlSVC() *Service {

	return &Service{
		Context: context.Background(),
		CFAPI:   &mockCFAPI{},
		S3API:   &mockS3API{},
	}
}
