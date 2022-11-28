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

	user := createRandomUser(t)
	args := CreateAccountParams{
		Owner:    user.Username,
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

	var lastAccount Account
	//add few accounts just incase there are less than 5 accounts while this test runs
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	listAccounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, listAccounts)

	for _, account := range listAccounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}

}
