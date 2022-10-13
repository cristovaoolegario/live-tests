package entity

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "email@test.com")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if client.ID == "" {
		t.Errorf("unexpected empty ID")
	}
	if client.Name != "John Doe" {
		t.Errorf("unexpected name: %v", client.Name)
	}
	if client.Email != "email@test.com" {
		t.Errorf("unexpected email: %v", client.Email)
	}
	if client.Points != 0 {
		t.Errorf("unexpected points: %v", client.Points)
	}
}
