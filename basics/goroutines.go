package basics

import (
	goroutine "basics/basics/goroutines"
	"errors"
	"log"
	"strings"
)

func GoRoutine(module string) {
	if strings.EqualFold(module, "Basic") {
		goroutine.Basic()
	} else if strings.EqualFold(module, "NoSync") {
		goroutine.NoSync()
	} else if strings.EqualFold(module, "Mutex") {
		goroutine.SyncWithMutex()
	} else if strings.EqualFold(module, "Channels") {
		goroutine.Channels()
	} else if strings.EqualFold(module, "ProducerConsumer") {
		goroutine.ProducerConsumerExampleWithbackPressure()
	} else {
		err := errors.New("routine example not found")
		log.Fatal(err)
	}
}
