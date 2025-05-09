package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Caio", "caio@rocha.me")
	assert.Equal(t, "Caio", client.Name)
	assert.Equal(t, "caio@rocha.me", client.Email)
	assert.NotEmpty(t, client.ID)
	assert.NotEmpty(t, client.CreatedAt)
	assert.NotEmpty(t, client.UpdatedAt)
	assert.NoError(t, err)
}

func TestCreateNewClientWhenNameIsEmpty(t *testing.T) {
	_, err := NewClient("", "caio@rocha.me")
	assert.Error(t, err)
}

func TestCreateNewClientWhenEmailIsEmpty(t *testing.T) {
	_, err := NewClient("Caio", "")
	assert.Error(t, err)
}
