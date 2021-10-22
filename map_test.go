package learngogoroutines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, group *sync.WaitGroup, key, value interface{}) {
	defer group.Done()

	group.Add(1)
	data.Store(key, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, group, "index_"+strconv.Itoa(i), i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
