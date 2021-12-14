package account

import (
	"math"
	"sync"
)

// Mutex references
// https://go.dev/tour/concurrency/9
// https://gobyexample.com/mutexes

type Account struct {
	mu sync.Mutex // Declare mutex
	balance int64
	open    bool
}

func Open(amount int64) *Account {
	// If Open is given a negative initial deposit, it must return nil.
	if amount < 0 {
		return nil
	}
	var a Account
	
	// Mutex lock
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance = amount
	a.open = true
	return &a
}

func (a *Account) Balance() (int64, bool) {
	// Mutex lock
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance, a.open
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	// Mutex lock
	a.mu.Lock()
	defer a.mu.Unlock()
	
	// Account is closed
	if !a.open {
		return a.balance, false
	}
	// Check for negative withdrawals
	if amount < 0 && a.balance < int64(math.Abs(float64(amount))) {
		// Trying to withdraw more than balance
		return a.balance, false
	}
	// Deposit/Withdraw
	a.balance += amount
	return a.balance, true
}

func (a *Account) CloseAndZero() {
	a.open = false
	a.balance = 0
}

func (a *Account) Close() (int64, bool) {
	// Mutex lock
	a.mu.Lock()
	defer a.mu.Unlock()
	
	// CLose out after this function returns
	defer a.CloseAndZero()

	return a.balance, a.open
}
