package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	now := time.Now()

	// empty password
	got, err := NewAccount("test@email.com", "", now)
	assert.Error(t, err)
	assert.Nil(t, got)

	// success
	got1, err := NewAccount("test@email.com", "test", now)
	assert.NoError(t, err)
	assert.NotEmpty(t, got1.ID)
	assert.Equal(t, "test@email.com", got1.Email)
	assert.NotEmpty(t, got1.Password)
	assert.Equal(t, false, got1.IsVerified)
	assert.Equal(t, false, got1.IsAdmin)
	assert.Equal(t, now, got1.CreatedAt)
}
