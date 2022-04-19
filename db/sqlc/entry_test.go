package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vengolan/simplebank/db/util"
)

func createRandomEntry(t *testing.T, ac Account) Entry {
	arg := CreateEntryParams{
		AccountID: ac.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	return entry
}

func Test_CreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func Test_GetEntry(t *testing.T) {

	account := createRandomAccount(t)
	createdEntry := createRandomEntry(t, account)

	gotEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotEntry)

	require.Equal(t, account.ID, gotEntry.AccountID)
	require.Equal(t, createdEntry.Amount, gotEntry.Amount)
	require.Equal(t, createdEntry.CreatedAt, gotEntry.CreatedAt)

}

func Test_ListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		createRandomEntry(t, account)
	}

	entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{5, 0})
	require.NoError(t, err)

	require.Equal(t, len(entries), 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
