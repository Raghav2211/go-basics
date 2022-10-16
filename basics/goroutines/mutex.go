package goroutines

import (
	"fmt"
	"sync"
)

var mutexWg = sync.WaitGroup{}
var readWriteMutex = sync.RWMutex{}

var mutexCounter int = 0

func SyncWithMutex() {
	// it solves the ordering problem of counter but does not fully solve the exact order
	// becuase if read happen twice and write happens ones then it'll read the same counter data twice
	for i := 0; i < 10; i++ {
		mutexWg.Add(2)
		go mutexGreetCounterVarUsingLockUnlock(greetCounter)
		go mutexIncrementCounterUsingLockUnlock(incrementCounter)
	}
	mutexWg.Wait()
	println()
	mutexCounter = 0
	// completely destroying concurrency model using lock unlock
	//becuase all the mutexs forcing the data to be synchornized
	for i := 0; i < 10; i++ {
		mutexWg.Add(2)
		readWriteMutex.RLock()
		go mutexGreetCounterVarUsingUnlock(greetCounter)
		readWriteMutex.Lock()
		go mutexIncrementCounterUsingUnlock(incrementCounter)

	}
	mutexWg.Wait()

}
func mutexGreetCounterVarUsingLockUnlock(greetCounter func()) {
	readWriteMutex.RLock()
	greetCounter()
	readWriteMutex.RUnlock()
	mutexWg.Done()
}

func mutexIncrementCounterUsingLockUnlock(incrementCounter func()) {
	readWriteMutex.Lock()
	incrementCounter()
	readWriteMutex.Unlock()
	mutexWg.Done()
}

func mutexGreetCounterVarUsingUnlock(greetCounter func()) {
	greetCounter()
	readWriteMutex.RUnlock()
	mutexWg.Done()
}

func mutexIncrementCounterUsingUnlock(incrementCounter func()) {
	incrementCounter()
	readWriteMutex.Unlock()
	mutexWg.Done()
}

func greetCounter() {
	fmt.Println("Hello counter ", mutexCounter, "!!")
}
func incrementCounter() {
	mutexCounter++
}
