package goroutines

import (
	"fmt"
	"sync"
)

var noSyncwg = sync.WaitGroup{}

var noSyncCounter int = 0

func NoSync() {

	for i := 0; i < 10; i++ {
		noSyncwg.Add(2)
		go noSyncgreetCounterVar()
		go noSyncincrementCounter()

	}
	noSyncwg.Wait()

}

func noSyncgreetCounterVar() {
	fmt.Println("Hello counter ", noSyncCounter, "!!")
	noSyncwg.Done()
}

func noSyncincrementCounter() {
	noSyncCounter++
	noSyncwg.Done()
}
