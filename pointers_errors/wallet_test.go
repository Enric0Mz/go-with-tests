package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {
	t.Run("should add 10 bitcoins to balance when .Deposit(10)", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expect := Bitcoin(10)

		actual := wallet.Balance()

		require.Equal(t, expect, actual)
	})
	t.Run("should withdraw 10 bitcoins of balance when .Withdraw()", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		require.NoError(t, err)
		expect := Bitcoin(10)
		actual := wallet.Balance()

		require.Equal(t, expect, actual)
	})

	t.Run("should return error when withdrawing more bitcoins then total balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(200))
		require.Error(t, err)
		require.Equal(t, withdrawErr, err)

	})

	t.Run("should return 10 BTC when printing a bitcoin of 10", func(t *testing.T) {
		btc := Bitcoin(10)

		expect := "10 BTC"
		actual := btc.String()

		require.Equal(t, expect, actual)
	})
}
