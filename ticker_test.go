package learngogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	tickstop := make(chan bool)

	go func() {
		tickstop <- false
		time.Sleep(5 * time.Second)
		ticker.Stop()
		tickstop <- true
	}()

	for {
		select {
		case data := <-ticker.C:
			fmt.Println(data)
		case data := <-tickstop:
			if data {
				fmt.Println("Done")
				t.FailNow()
			}
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for data := range channel {
		fmt.Println(data)
	}
}
