package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOne() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOne)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
