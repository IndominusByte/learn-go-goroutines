package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup, num int) {
	defer group.Done()

	group.Add(1)

	fmt.Println(num)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("DONE")
}
