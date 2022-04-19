package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vengolan/simplebank/db/util"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	ac, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Owner, ac.Owner)
	require.Equal(t, args.Balance, ac.Balance)
	require.Equal(t, args.Currency, ac.Currency)

	require.NotZero(t, ac.ID)
	require.NotZero(t, ac.CreatedAt)
	return ac
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func Test_GetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, createdAccount.Balance, gotAccount.Balance)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)

	require.WithinDuration(t, createdAccount.CreatedAt, gotAccount.CreatedAt, time.Second)
}

func Test_UpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	args := UpdateAccountParams{ID: createdAccount.ID, Balance: util.RandomMoney()}
	gotAccount, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, args.Balance, gotAccount.Balance)
	require.NotEqual(t, createdAccount.Balance, gotAccount.Balance)
	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, gotAccount.CreatedAt, time.Second)

}

func Test_DeleteAccount(t *testing.T) {

	createdAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.Error(t, err, sql.ErrNoRows)
	require.Empty(t, gotAccount)
}

func Test_ListAccount(t *testing.T) {

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	listAccounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(listAccounts), 5)

	for _, account := range listAccounts {
		require.NotEmpty(t, account)
	}

}
