package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var channelWaitGroup sync.WaitGroup = sync.WaitGroup{}

func Channels() {
	channelWaitGroup.Add(2)
	channel := make(chan int)

	// read & write from both goroutines
	go func() {
		i := <-channel
		fmt.Println("[ReaderWriter] Read", i, "from channel")
		fmt.Println("[ReaderWriter] Write", 1, "from channel")
		channel <- 1
		channelWaitGroup.Done()
	}()

	go func() {
		fmt.Println("[ReaderWriter] Write", 0, "into channel")
		channel <- 0
		i := <-channel
		fmt.Println("[ReaderWriter]  Read", i, "from channel")
		channelWaitGroup.Done()
	}()

	channelWaitGroup.Wait()
	println()

	// separate reader writer
	channel = make(chan int)
	channelWaitGroup.Add(2)
	go channelReader(channel)
	go channelWriter(channel, 2)
	channelWaitGroup.Wait()
	println()

	// solve counter problem with channel (looks like producer & consumer)
	channel = make(chan int)
	channelWaitGroup.Add(2)
	go readCounter(channel)
	go writeCounter(channel, 10)

	channelWaitGroup.Wait()

}

func channelReader(channel <-chan int) {
	i := <-channel
	fmt.Println("[Reader] Read", i, "from channel")
	channelWaitGroup.Done()
}

func channelWriter(channel chan<- int, data int) {
	fmt.Println("[Writer] Write", data, "into channel")
	channel <- data
	channelWaitGroup.Done()

}

func readCounter(channel <-chan int) {
	for {
		i, ok := <-channel
		if !ok {
			break
		}
		time.Sleep(1 * time.Second) // just for showcase // don't do this on production
		fmt.Println("Counter : ", i)
	}

	// second way
	// for i := range channel {
	// 	fmt.Println("Counter : ", i)
	// 	time.Sleep(1 * time.Second) // just for showcase // don't do this on production
	// }
	channelWaitGroup.Done()
}

func writeCounter(channel chan<- int, data int) {
	for i := 0; i < data; i++ {
		channel <- i
	}
	close(channel)
	channelWaitGroup.Done()
}
