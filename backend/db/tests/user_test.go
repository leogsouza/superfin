package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	db "leogsouza.dev/superfin/db/sqlc"
)

func TestCreateUser(t *testing.T) {
	arg := db.CreateUserParams{
		Email:    "user@test.com",
		Password: "pass_secret",
	}

	user, err := testQuery.CreateUser(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Email, arg.Email)
	assert.Equal(t, user.Password, arg.Password)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

}
