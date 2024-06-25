package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456789")
	assert.Equal(t, nil, err)

	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456789")
	assert.Equal(t, nil, err)

	assert.NotEqual(t, "123456789", user.Password)
	assert.True(t, user.ValidatePassword("123456789"))
	assert.False(t, user.ValidatePassword("12345678"))
}
