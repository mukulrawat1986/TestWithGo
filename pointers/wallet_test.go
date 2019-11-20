package main

import "testing"

func TestWallet(t *testing.T) {

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
		assertBalance(t, wallet, Bitcoin(20)) // confirm that no withdrawl is done
		assertError(t, err)                   // confirm that an error is returned
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}
