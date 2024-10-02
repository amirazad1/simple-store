package db

import (
	"context"
	"database/sql"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: faker.Name(),
		Password: faker.Password(),
		FullName: faker.Name(),
		Mobile:   faker.Phonenumber()[0:10],
		PasswordChangedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	_, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	user, err := testQueries.GetUser(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Mobile, user.Mobile)
	require.WithinDuration(t, arg.PasswordChangedAt.Time, user.PasswordChangedAt.Time, time.Minute)

	return user
}

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestQueries_GetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Mobile, user2.Mobile)
	require.WithinDuration(t, user1.PasswordChangedAt.Time, user2.PasswordChangedAt.Time, time.Minute)
}

func TestQueries_UpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		Username: user1.Username,
		FullName: faker.Name(),
		Mobile:   faker.Phonenumber()[0:10],
	}

	result, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	affected, err := result.RowsAffected()
	require.NoError(t, err)
	require.NotEmpty(t, affected)

	user2, err := testQueries.GetUser(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, arg.Mobile, user2.Mobile)
}

func TestQueries_UpdateUserPassword(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserPasswordParams{
		Username: user1.Username,
		Password: faker.Password(),
		PasswordChangedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	result, err := testQueries.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)
	affected, err := result.RowsAffected()
	require.NoError(t, err)
	require.NotEmpty(t, affected)

	user2, err := testQueries.GetUser(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, arg.Password, user2.Password)
	require.WithinDuration(t, arg.PasswordChangedAt.Time, user2.PasswordChangedAt.Time, time.Minute)
}

func TestQueries_DeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestQueries_ListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
