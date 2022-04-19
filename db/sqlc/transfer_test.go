package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vengolan/simplebank/db/util"
)

func createRandomTransfer(t *testing.T, from, to Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, from.ID, transfer.FromAccountID)
	require.Equal(t, to.ID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func Test_CreateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func Test_GetTransfers(t *testing.T) {

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	createdTransfer := createRandomTransfer(t, fromAccount, toAccount)

	gotTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, gotTransfer)
	require.Equal(t, gotTransfer.FromAccountID, createdTransfer.FromAccountID)
	require.Equal(t, gotTransfer.ToAccountID, createdTransfer.ToAccountID)
	require.Equal(t, gotTransfer.ID, createdTransfer.ID)
	require.Equal(t, createdTransfer.CreatedAt, gotTransfer.CreatedAt)
}

func Test_ListTransfers(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 0,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, len(transfers), 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
