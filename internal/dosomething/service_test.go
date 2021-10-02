package dosomething

import (
	"testing"

	"github.com/o-aloqaily/go-project-starter/internal/apiclient"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockApiClient := apiclient.NewMockClient(ctrl)
	assert.NotNil(t, NewService(mockApiClient))
}

func TestDoSomething_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockApiClient := apiclient.NewMockClient(ctrl)
	s := NewService(mockApiClient)
	// expect the function CallSomething of the mock client to be called
	// with the same values
	mockApiClient.EXPECT().CallSomething(apiclient.CallSomethingRequest{
		Field1: "blablabla",
	})

	s.DoSomething("blablabla", "blablabla")
}

func TestDoSomething_ApiError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockApiClient := apiclient.NewMockClient(ctrl)
	s := NewService(mockApiClient)
	// expect the function CallSomething of the mock client to be called
	// with the same values
	mockApiClient.EXPECT().CallSomething(apiclient.CallSomethingRequest{
		Field1: "blablabla",
	}).Return(apiclient.CallSomethingResponseErr{
		ErrorCode:    "400",
		ErrorMessage: "error",
	})

	err := s.DoSomething("blablabla", "blablabla")
	assert.NotNil(t, err)
}
