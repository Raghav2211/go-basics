package goroutines

import (
	"fmt"
	"sync"
)

var channelWaitGroup sync.WaitGroup = sync.WaitGroup{}

func ChannelBasic() {
	channelWaitGroup.Add(2)
	channel1 := make(chan int)

	// read & write from both goroutines
	go func() {
		i := <-channel1
		fmt.Println("[ReaderWriter] Read", i, "from channel")
		fmt.Println("[ReaderWriter] Write", 1, "from channel")
		channel1 <- 1
		channelWaitGroup.Done()
	}()

	go func() {
		fmt.Println("[ReaderWriter] Write", 0, "into channel")
		channel1 <- 0
		i := <-channel1
		fmt.Println("[ReaderWriter]  Read", i, "from channel")
		channelWaitGroup.Done()
	}()

	// separate reader writer
	channel2 := make(chan int)
	channelWaitGroup.Add(2)
	go channelReader(channel2)
	go channelWriter(channel2, 2)
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
