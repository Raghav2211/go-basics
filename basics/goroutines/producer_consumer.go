package goroutines

import (
	"fmt"
)

var jobs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func producer(link chan int, backpressure <-chan int) {
	for {
		backpressureEvent := <-backpressure
		if len(jobs) == 0 {
			close(link) //  close link channel
			break
		}
		eles := jobs[0:backpressureEvent]
		jobs = jobs[backpressureEvent:len(jobs)]
		for ele := range eles {
			fmt.Println("Send job with id -- ", eles[ele])
			link <- eles[ele]
		}
	}
}
func consumer(link <-chan int, backpressure chan int, isProcessDone chan bool) {
	for {
		backpressure <- 1
		jobid, isClose := <-link
		if !isClose {
			close(backpressure) //  close backpressure event channel
			isProcessDone <- true
			break
		} else {
			fmt.Println("Receive job with id -- ", jobid)
		}
	}
}

func ProducerConsumerExampleWithbackPressure() {
	link := make(chan int)
	backpressure := make(chan int)
	done := make(chan bool)
	go producer(link, backpressure)
	go consumer(link, backpressure, done)
	<-done
	fmt.Println("All Done !! :)")
}
