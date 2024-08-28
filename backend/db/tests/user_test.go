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

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := utils.GenerateHashPassword(utils.RandomString(8))
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
	return user
}

func TestCreateUser(t *testing.T) {

	user1 := createRandomUser(t)
	user2, err := testQuery.CreateUser(context.Background(), db.CreateUserParams{
		Email:    user1.Email,
		Password: user1.Password,
	})
	assert.Error(t, err)
	assert.Empty(t, user2)

}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	newPassword, err := utils.GenerateHashPassword(utils.RandomString(8))
	if err != nil {
		log.Fatal("Unable to generate Hash", err)
	}

	arg := db.UpdateUserPasswordParams{
		Password:  newPassword,
		ID:        user.ID,
		UpdatedAt: time.Now(),
	}

	newUser, err := testQuery.UpdateUserPassword(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, user.Email, newUser.Email)
	assert.Equal(t, newUser.Password, arg.Password)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)
}

func TestGetUserByID(t *testing.T) {
	user := createRandomUser(t)

	getUser, err := testQuery.GetUserById(context.Background(), user.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, getUser)

	assert.Equal(t, user.Email, getUser.Email)
	assert.Equal(t, user.Password, getUser.Password)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	getUser, err := testQuery.GetUserByEmail(context.Background(), user.Email)

	assert.NoError(t, err)
	assert.NotEmpty(t, getUser)

	assert.Equal(t, user.Email, getUser.Email)
	assert.Equal(t, user.Password, getUser.Password)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := testQuery.DeleteUser(context.Background(), user.ID)

	assert.NoError(t, err)

	getUser, err := testQuery.GetUserById(context.Background(), user.ID)
	assert.Error(t, err)
	assert.Empty(t, getUser)
}

func TestListUsers(t *testing.T) {
	listLen := 50
	for i := 0; i < listLen; i++ {
		createRandomUser(t)
	}

	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  int32(listLen),
	}

	users, err := testQuery.ListUsers(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Len(t, users, listLen)
}
