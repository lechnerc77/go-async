package waitgroup

import (
	"strconv"
	"sync"

	"github.com/lechnerc77/goasyncdemo/goroutines"
)

func Run() {

	var waitGroup sync.WaitGroup

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			goroutines.Count("Count run: " + strconv.Itoa(i))
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()
}
