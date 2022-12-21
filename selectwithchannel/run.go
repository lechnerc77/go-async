package selectwithchannel

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetFirstResult(result chan<- int) {

	randomNumber := rand.Intn(10)
	ms := time.Duration(randomNumber * int(time.Second))
	time.Sleep(ms)
	result <- 1
}

func GetSecondResult(result chan<- int) {
	randomNumber := rand.Intn(10)
	ms := time.Duration(randomNumber * int(time.Second))
	time.Sleep(ms)
	result <- 2
}

func Run() {

	firstResult := make(chan int)
	secondResult := make(chan int)

	go GetFirstResult(firstResult)
	go GetSecondResult(secondResult)

	results := []string{}

	for {
		select {
		case result := <-firstResult:
			results = append(results, "First result: "+strconv.Itoa(result))
		case result := <-secondResult:
			results = append(results, "Second result: "+strconv.Itoa(result))
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
		default:
			fmt.Println("No result yet")
			time.Sleep(100 * time.Millisecond)
		}

		if len(results) == 2 {
			break
		}

	}

	for _, result := range results {
		fmt.Println(result)
	}

}
