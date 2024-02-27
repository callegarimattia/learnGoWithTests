package banking

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {
	t.Parallel()
	t.Run("deposit", func(t *testing.T) {
		t.Parallel()

		wallet := Wallet{0}
		wallet.Deposit(Bitcoin(10))
		assert.Equal(t, wallet.balance, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		t.Parallel()

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		require.NoError(t, err)
		assert.Equal(t, wallet.balance, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		t.Parallel()

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		require.ErrorIs(t, err, ErrInsufficientFunds)
		assert.Equal(t, wallet.balance, Bitcoin(20))
	})
}
