package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Zé", "ze@ze.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "Zé")
	assert.Equal(t, user.Email, "ze@ze.com")
	assert.NotEqual(t, user.Password, "123456")
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Zé", "ze@ze.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}
