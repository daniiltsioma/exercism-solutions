package account

import "sync"

// Define the Account type here.
type Account struct {
	mu sync.Mutex

	balance int64
	open bool
}

// Open opens a new account with a given balance.
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance: amount,
		open: true,
	}
}

// Balance checks the balance of an open account.
func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	return a.balance, true
}

// Deposit allows deposits to and withdrawals from an open account.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}
	// do not allow withdrawal overdrafts.
	if a.balance + amount < 0 {
		return a.balance, false
	}
	
	a.balance += amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	endingBalance := a.balance
	a.balance = 0
	a.open = false

	return endingBalance, true
}
