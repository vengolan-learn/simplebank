package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vengolan/simplebank/db/util"
)

func TestCreateAccount(t *testing.T) {

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

}
