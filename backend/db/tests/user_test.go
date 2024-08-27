package db_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	db "leogsouza.dev/superfin/db/sqlc"
	"leogsouza.dev/superfin/utils"
)

func TestCreateUser(t *testing.T) {
	hashedPassword, err := utils.GenerateHasPassword(utils.RandomString(8))
	if err != nil {
		log.Fatal("Unable to generate hash password", err)
	}
	arg := db.CreateUserParams{
		Email:    utils.RandomEmail(),
		Password: hashedPassword,
	}

	user, err := testQuery.CreateUser(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Email, arg.Email)
	assert.Equal(t, user.Password, hashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

}
