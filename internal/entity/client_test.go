package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "teste@teste.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "teste@teste.com", client.Email)
}

func TestShouldNotCreateClientWithInvalidArgs(t *testing.T) {
	client, err := NewClient("", "teste@teste.com")
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestShouldUpdateClient(t *testing.T) {
	client, err := NewClient("John Doe", "teste@teste.com")
	client.Update("Johnas Doe", "teste@teste.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Johnas Doe", client.Name)
	assert.Equal(t, "teste@teste.com", client.Email)
}

func TestShouldNotUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "teste@teste.com")
	err := client.Update("", "teste@teste.com")
	assert.Error(t, err, "name is required")
}
