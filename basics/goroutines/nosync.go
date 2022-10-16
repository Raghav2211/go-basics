package goroutines

import (
	"fmt"
	"sync"
)

var wrongSyncwg = sync.WaitGroup{}

var wrongSyncCounter int = 0

func NoSync() {

	for i := 0; i < 10; i++ {
		wrongSyncwg.Add(2)
		go greetCounterVar()
		go incrementCounter()

	}
	wrongSyncwg.Wait()

}

func greetCounterVar() {
	fmt.Println("Hello counter ", wrongSyncCounter, "!!")
	wrongSyncwg.Done()
}

func incrementCounter() {
	wrongSyncCounter++
	wrongSyncwg.Done()
}
