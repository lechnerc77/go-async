package accounting

import "sync"

type Account struct {
	mutex   sync.RWMutex
	balance int
}

func NewAccount() *Account {

	return &Account{
		balance: 0,
	}
}

func (account *Account) Balance() int {
	account.mutex.RLock()
	defer account.mutex.RUnlock()

	return account.balance
}

func (account *Account) Deposit(amount int) {
	account.mutex.Lock()
	defer account.mutex.Unlock()

	account.balance += amount

}

func (account *Account) Withdraw(amount int) {
	account.mutex.Lock()
	defer account.mutex.Unlock()

	account.balance -= amount
}
