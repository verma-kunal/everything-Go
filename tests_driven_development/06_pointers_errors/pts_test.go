package main

import (
	"testing"
)

// test for wallet struct
func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		// why "Bitcoin(10)" - it’s a type conversion, which tells Go, "treat the integer value 10 as a Bitcoin type."
		// fmt.Printf("address of balance in test: %p\n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(10)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startBalance)

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
