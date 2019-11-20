package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficent funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("Wanted an error but didn't get one")
		}
	})
}
