package learngogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "kosong"
		},
	}

	pool.Put("oman")
	pool.Put("pradipta")
	pool.Put("dewantara")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)

			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("DONE")
}
