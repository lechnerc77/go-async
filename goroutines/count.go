package goroutines

import (
	"fmt"
	"strconv"
	"time"
)

func Count(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Println(prefix + ": " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 500)
	}
}
