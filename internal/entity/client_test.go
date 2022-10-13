package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithValidData(t *testing.T) {
	client, err := NewClient("John Doe", "email@test.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, 0, client.Points)
}

func TestNewClientWithInvalidData(t *testing.T) {
	client, err := NewClient("", "j@j.com")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client name is required")

	client, err = NewClient("John doe", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client email is required")
}
