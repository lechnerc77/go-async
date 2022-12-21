package primes

import "fmt"

func Run() {
	candidates := make(chan int, 100)
	done := make(chan bool)

	go func() {
		for {
			candidate, more := <-candidates
			if more {
				if IsPrime(candidate) {
					fmt.Println(candidate)
				}
			} else {
				done <- true
				return
			}
		}

	}()

	for i := 0; i < 10; i++ {
		candidates <- i
	}

	close(candidates)

	<-done
}
