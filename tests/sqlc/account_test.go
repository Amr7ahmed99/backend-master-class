package tests

import (
	"backend-master-class/api/request_params"
	"backend-master-class/util"
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	db "backend-master-class/db/sqlc"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) db.Account {
	user := createRandomUser(t)
	arg := db.CreateAccountParams{
		Owner:      user.Username,
		Balance:    util.RandomMoney(),
		CurrencyID: int32(util.RandomCurrency()),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.CurrencyID, account.CurrencyID)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)
	// require.Equal(t, createdAccount.ID, fetchedAccount.ID)
	// require.Equal(t, createdAccount.Owner, fetchedAccount.Owner)
	// require.Equal(t, createdAccount.Balance, fetchedAccount.Balance)
	// require.Equal(t, createdAccount.Currency, fetchedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, fetchedAccount.CreatedAt, time.Second)
	require.True(t, reflect.DeepEqual(createdAccount, fetchedAccount))
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	arg := db.UpdateAccountParams{
		ID:         createdAccount.ID,
		Balance:    util.RandomMoney(),
		CurrencyID: int32(util.RandomCurrency()),
	}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.False(t, reflect.DeepEqual(createdAccount, updatedAccount))
	require.Equal(t, arg.ID, updatedAccount.ID)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, arg.CurrencyID, updatedAccount.CurrencyID)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchedAccount)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	limit := 5
	req := request_params.ListAccountRequest{
		PageSize: 5,
		PageID:   1,
	}
	fetchedAccounts, err := testQueries.ListAccount(context.Background(), db.ListAccountParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	require.NoError(t, err)
	require.Len(t, fetchedAccounts, limit)

	for _, value := range fetchedAccounts {
		require.NotEmpty(t, value)
	}
}
