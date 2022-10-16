package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func GoRoutine() {

	fmt.Printf("No of OS threads %d", runtime.GOMAXPROCS(-1))
	println()
	greetGoRountine("main thread")
	println()
	//threads that are scheduled by a virtual machine (VM) instead of natively by the underlying operating system
	// it create different stack
	wg.Add(1)
	go greetWithCustomGoRountine(greetGoRountine, "Green thread")

	wg.Add(1)
	// using annonymous function
	go func() {
		fmt.Println("Hello user using annonymous go routin")
		println()
		wg.Done()
	}()

	wg.Add(1)
	// using variable
	var user string = "Raghav"
	go func() {
		// user variable create dependency between main function & goroutine variable
		// which is bad idea because if we change the variable and call again it gives unpredictable output
		// demonstrate via below example
		fmt.Println("Hello user ", user, "!! using anonymous go routin")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello user ", user, "!! call again using anonymous go routin")
		wg.Done()
		println()
	}()
	user = "Another user"

	// good way to pass variable as we do with functions
	wg.Add(1)
	user = "Raghav"
	go func(user string) {
		fmt.Println("Hello user ", user, "!! call again using anonymous go routin with passing variable")
		wg.Done()
		println()
	}(user)
	user = "Another user"
	wg.Wait()

}

func greetGoRountine(threadName string) {
	fmt.Println("Hello user using ,", threadName, "!!")
}

func greetWithCustomGoRountine(greetGoRoutine func(name string), name string) {
	greetGoRoutine(name)
	println()
	wg.Done()
}
