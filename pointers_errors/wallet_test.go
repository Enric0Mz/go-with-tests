package pointerserrors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	expect := 10

	actual := wallet.Balance()

	require.Equal(t, expect, actual)
}
