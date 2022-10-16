package goroutines

import (
	"fmt"
	"sync"
)

var mutexWg = sync.WaitGroup{}
var readWriteMutex = sync.RWMutex{}

var mutexCounter int = 0

// it solves the ordering problem of counter but does not fully solve the exact order
// becuase if read happen twice and write happens ones then it'll read the same counter data twice
func SyncWithMutex() {

	for i := 0; i < 10; i++ {
		mutexWg.Add(2)
		go mutexGreetCounterVar()
		go mutexIncrementCounter()

	}
	mutexWg.Wait()

}

func mutexGreetCounterVar() {
	readWriteMutex.RLock()
	fmt.Println("Hello counter ", mutexCounter, "!!")
	readWriteMutex.RUnlock()
	mutexWg.Done()
}

func mutexIncrementCounter() {
	readWriteMutex.Lock()
	mutexCounter++
	readWriteMutex.Unlock()
	mutexWg.Done()
}
