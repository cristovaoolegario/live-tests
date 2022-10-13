package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Client struct {
	ID     string
	Name   string
	Email  string
	Points int
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:     uuid.New().String(),
		Name:   name,
		Email:  email,
		Points: 0,
	}
	if client.Name == "" {
		return nil, errors.New("client name is required")
	}
	if client.Email == "" {
		return nil, errors.New("client email is required")
	}
	return client, nil
}
