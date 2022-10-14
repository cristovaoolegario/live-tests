package usecase

import (
	"live-tests/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientRepositoryMock struct {
	mock.Mock
}

func (c *ClientRepositoryMock) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	mockRepo := new(ClientRepositoryMock)
	mockRepo.On("Save", mock.Anything).Return(nil)

	createClientUseCase := NewClientUseCase(mockRepo)

	input := CreateClientInputDTO{
		Name:  "John Doe",
		Email: "j@j.com",
	}

	output, err := createClientUseCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j.com", output.Email)
	assert.Equal(t, 0, output.Points)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Save", mock.Anything)

}
