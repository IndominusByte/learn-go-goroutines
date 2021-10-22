package learngogoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	channel <- "Nyoman Pradipta Dewantara"
	time.Sleep(2 * time.Second)
	fmt.Println("berhasil mengirim data ke channel")
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Nyoman Pradipta Dewantara"
		time.Sleep(2 * time.Second)
		fmt.Println("berhasil mengirim data ke channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestCreateChannelAsParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	channel <- "oman oman oman"
	time.Sleep(2 * time.Second)
	fmt.Println("berhasil mengirim data ke channel")

}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "nyoman"
		channel <- "pradipta"
		channel <- "dewantara"
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	time.Sleep(5 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("DONE")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
