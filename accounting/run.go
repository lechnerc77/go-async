package accounting

import (
	"fmt"
	"sync"
)

func Run() {
	var waitGroup sync.WaitGroup
	account := NewAccount()

	for i := 0; i < 10000; i++ {
		waitGroup.Add(1)
		go func() {
			account.Deposit(100)
			//account.Withdraw(100)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Println(account.Balance())
}
