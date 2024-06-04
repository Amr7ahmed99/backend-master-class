package tests

import (
	request_params "backend-master-class/apis/requests"
	"backend-master-class/util"
	"context"
	"reflect"
	"testing"
	"time"

	db "backend-master-class/db/sqlc"

	"github.com/stretchr/testify/require"
)

func TestCreateUp(t *testing.T) {
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T) db.User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := db.CreateUserParams{
		Username:       util.RandomOwner(),
		FullName:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	createdUser := CreateRandomUser(t)
	fetchedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)
	// require.Equal(t, createdAccount.ID, fetchedAccount.ID)
	// require.Equal(t, createdAccount.Owner, fetchedAccount.Owner)
	// require.Equal(t, createdAccount.Balance, fetchedAccount.Balance)
	// require.Equal(t, createdAccount.Currency, fetchedAccount.Currency)
	require.WithinDuration(t, createdUser.CreatedAt, fetchedUser.CreatedAt, time.Second)
	require.True(t, reflect.DeepEqual(createdUser, fetchedUser))
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	limit := 5
	req := request_params.ListUsersRequest{
		PageSize: 5,
		PageID:   1,
	}
	fetchedUsers, err := testQueries.ListUser(context.Background(), db.ListUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	require.NoError(t, err)
	require.Len(t, fetchedUsers, limit)

	for _, value := range fetchedUsers {
		require.NotEmpty(t, value)
	}
}
