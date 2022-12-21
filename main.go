package main

import (
	"context"
	"fmt"
	"time"
)

func CheckCandidate(ctx context.Context, candidate int, result chan<- bool) {

	isPrime := true

	if candidate < 2 {
		isPrime = false
	}

	for i := 2; i < candidate; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Canceled Checking", candidate)
			return
		default:
			if candidate%i == 0 {
				isPrime = false
			}
		}
	}

	fmt.Println("Done Checking", candidate)
	result <- isPrime

}

func main() {

	firstCandidate := 1000000001
	secondCandidate := 1000000003

	firstResult := make(chan bool)
	secondResult := make(chan bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go CheckCandidate(ctx, firstCandidate, firstResult)
	go CheckCandidate(ctx, secondCandidate, secondResult)

	select {
	case result := <-firstResult:
		fmt.Println(firstCandidate, "is prime: ", result)
	case result := <-secondResult:
		fmt.Println(secondCandidate, "is prime: ", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout, canceling ...")
		cancel()
	}

	time.Sleep(10 * time.Second)

}
